package service

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/kimtuna/bsh/blockchain"
	"github.com/kimtuna/bsh/models"
	"github.com/kimtuna/bsh/setup"
)

type CompanyService struct {
	contractClient *blockchain.ContractClient
}

func NewCompanyService(client *blockchain.ContractClient) *CompanyService {
	return &CompanyService{
		contractClient: client,
	}
}

// testServerAccess 서버 접근 테스트
func testServerAccess(ip string, port uint16, username string, password string) error {
	// ping으로 서버 연결 테스트
	cmd := exec.Command("ping", "-c", "4", ip)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("서버가 없습니다 (IP: %s): %v", ip, string(output))
	}
	return nil
}

// RegisterCompanyInternal 회사 등록 처리
func (s *CompanyService) RegisterCompanyInternal(req models.RegisterRequest) (*models.Response, error) {
	fmt.Printf("[DEBUG] RegisterCompanyInternal 요청 시작\n")

	// 1. 이더리움 주소 유효성 검사
	if !common.IsHexAddress(req.CompanyWallet) {
		return &models.Response{
			Success: false,
			Message: "유효하지 않은 이더리움 주소",
		}, nil
	}

	// 2. 서버 접근 테스트
	if err := testServerAccess(req.IP, req.Port, req.ServerName, req.Password); err != nil {
		return &models.Response{
			Success: false,
			Message: "서버 접근 테스트 실패: " + err.Error(),
		}, nil
	}

	// 3. 스마트 컨트랙트를 통한 회사 등록 트랜잭션 실행
	tx, err := s.contractClient.RegisterCompany(
		common.HexToAddress(req.CompanyWallet),
		req.CompanyName,
		req.CeoName,
		req.Email,
		req.SubscriptionType,
	)
	if err != nil {
		fmt.Printf("[DEBUG] 회사 등록 트랜잭션 실패: %v\n", err)
		return &models.Response{
			Success: false,
			Message: "회사 등록 트랜잭션 실패: " + err.Error(),
		}, nil
	}

	// 4. 트랜잭션 영수증 확인
	receipt, err := s.contractClient.WaitForTransaction(tx)
	if err != nil {
		fmt.Printf("[DEBUG] 트랜잭션 확인 실패: %v\n", err)
		return &models.Response{
			Success: false,
			Message: "트랜잭션 확인 실패: " + err.Error(),
		}, nil
	}

	// 5. 트랜잭션이 성공했는지 확인
	if receipt.Status != 1 {
		return &models.Response{
			Success: false,
			Message: "트랜잭션이 실패했습니다",
		}, nil
	}

	// 6. 데이터베이스에 회사 정보 저장
	company := &models.CompanyRegistered{
		CompanyWallet:    req.CompanyWallet,
		CompanyName:      req.CompanyName,
		CeoName:          req.CeoName,
		Email:            req.Email,
		SubscriptionType: req.SubscriptionType,
		IP:               req.IP,
		ServerName:       req.ServerName,
		Port:             req.Port,
		IsActive:         true,
	}

	// 구독 기간 계산
	var subscriptionDuration time.Duration
	switch req.SubscriptionType {
	case 1:
		subscriptionDuration = 30 * 24 * time.Hour // 1개월
	case 2:
		subscriptionDuration = 90 * 24 * time.Hour // 3개월
	case 3:
		subscriptionDuration = 365 * 24 * time.Hour // 1년
	default:
		return &models.Response{
			Success: false,
			Message: "유효하지 않은 구독 유형",
		}, nil
	}

	company.SubscriptionEnd = time.Now().Add(subscriptionDuration).Unix()

	if err := setup.DB.Create(company).Error; err != nil {
		fmt.Printf("[DEBUG] 데이터베이스 저장 실패: %v\n", err)
		return &models.Response{
			Success: false,
			Message: "데이터베이스 저장 실패: " + err.Error(),
		}, nil
	}

	fmt.Printf("[DEBUG] 회사 등록 완료: %s\n", req.CompanyWallet)

	return &models.Response{
		Success: true,
		Message: "회사 등록이 완료되었습니다",
		Data: map[string]interface{}{
			"company_wallet":    company.CompanyWallet,
			"company_name":      company.CompanyName,
			"subscription_type": company.SubscriptionType,
			"subscription_end":  company.SubscriptionEnd,
			"is_active":         true,
			"registered_at":     time.Unix(company.CreatedAt, 0).Format(time.RFC3339),
		},
	}, nil
}

// RegisterCompany 회사 등록 API 핸들러
func (s *CompanyService) RegisterCompany(c *gin.Context) {
	fmt.Printf("[DEBUG] RegisterCompany 요청 시작\n")

	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("[DEBUG] JSON 바인딩 오류: %v\n", err)
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "잘못된 요청 형식: " + err.Error(),
		})
		return
	}

	fmt.Printf("[DEBUG] 회사 등록 요청 데이터: %+v\n", req)

	resp, err := s.RegisterCompanyInternal(req)
	if err != nil {
		fmt.Printf("[DEBUG] 내부 처리 오류: %v\n", err)
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "서버 내부 오류: " + err.Error(),
		})
		return
	}

	if resp.Success {
		c.JSON(http.StatusOK, *resp)
	} else {
		c.JSON(http.StatusBadRequest, *resp)
	}
}

// GetSubscriptionStatus 구독 상태 조회
func (s *CompanyService) GetSubscriptionStatus(c *gin.Context) {
	fmt.Printf("[DEBUG] GetSubscriptionStatus 요청 시작\n")

	var req models.SubscriptionStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("[DEBUG] JSON 바인딩 오류: %v\n", err)
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "잘못된 요청 형식: " + err.Error(),
		})
		return
	}

	fmt.Printf("[DEBUG] 구독 상태 조회 요청 데이터: %+v\n", req)

	// 1. 이더리움 주소 유효성 검사
	if !common.IsHexAddress(req.CompanyWallet) {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "유효하지 않은 이더리움 주소",
		})
		return
	}

	// 2. 데이터베이스에서 회사 정보 조회
	company, err := setup.GetServerAccess(req.CompanyWallet)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "등록되지 않은 회사입니다",
		})
		return
	}

	// 3. 구독 상태 계산
	now := time.Now()
	subscriptionEndTime := time.Unix(company.SubscriptionEnd, 0)
	isExpired := now.After(subscriptionEndTime)
	remainingDays := int(subscriptionEndTime.Sub(now).Hours() / 24)
	if remainingDays < 0 {
		remainingDays = 0
	}

	// 4. 응답 반환
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "구독 상태 조회 성공",
		Data: map[string]interface{}{
			"company_wallet":    company.CompanyWallet,
			"company_name":      company.CompanyName,
			"ceo_name":          company.CeoName,
			"email":             company.Email,
			"subscription_type": company.SubscriptionType,
			"subscription_end":  company.SubscriptionEnd,
			"is_active":         !isExpired,
			"is_expired":        isExpired,
			"remaining_days":    remainingDays,
			"checked_at":        now.Format(time.RFC3339),
		},
	})
}

// UpdateSubscriptionAfterPayment 결제 완료 후 구독 업데이트
func (s *CompanyService) UpdateSubscriptionAfterPayment(c *gin.Context) {
	fmt.Printf("[DEBUG] UpdateSubscriptionAfterPayment 요청 시작\n")

	var req struct {
		CompanyWallet    string `json:"company_wallet" binding:"required"`
		SubscriptionType uint8  `json:"subscription_type" binding:"required,min=1,max=3"`
		TransactionHash  string `json:"transaction_hash" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("[DEBUG] JSON 바인딩 오류: %v\n", err)
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "잘못된 요청 형식: " + err.Error(),
		})
		return
	}

	fmt.Printf("[DEBUG] 구독 업데이트 요청 데이터: %+v\n", req)

	// 1. 이더리움 주소 유효성 검사
	if !common.IsHexAddress(req.CompanyWallet) {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "유효하지 않은 이더리움 주소",
		})
		return
	}

	// 2. 데이터베이스에서 회사 정보 조회
	company, err := setup.GetServerAccess(req.CompanyWallet)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "등록되지 않은 회사입니다",
		})
		return
	}

	// 3. 구독 기간 계산
	var subscriptionDuration time.Duration
	switch req.SubscriptionType {
	case 1:
		subscriptionDuration = 30 * 24 * time.Hour // 1개월
	case 2:
		subscriptionDuration = 90 * 24 * time.Hour // 3개월
	case 3:
		subscriptionDuration = 365 * 24 * time.Hour // 1년
	default:
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "유효하지 않은 구독 유형",
		})
		return
	}

	// 4. 구독 만료일 업데이트 (기존 만료일에서 연장)
	currentEndTime := time.Unix(company.SubscriptionEnd, 0)
	now := time.Now()

	// 현재 시간이 만료일보다 늦으면 현재 시간부터, 아니면 기존 만료일부터 연장
	var newEndTime time.Time
	if now.After(currentEndTime) {
		newEndTime = now.Add(subscriptionDuration)
	} else {
		newEndTime = currentEndTime.Add(subscriptionDuration)
	}

	// 5. 데이터베이스 업데이트
	err = setup.UpdateSubscriptionInfo(req.CompanyWallet, req.SubscriptionType, newEndTime.Unix())
	if err != nil {
		fmt.Printf("[DEBUG] 데이터베이스 업데이트 실패: %v\n", err)
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "데이터베이스 업데이트 실패: " + err.Error(),
		})
		return
	}

	fmt.Printf("[DEBUG] 구독 업데이트 완료: %s\n", req.CompanyWallet)

	// 6. 응답 반환
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "구독이 성공적으로 연장되었습니다",
		Data: map[string]interface{}{
			"company_wallet":    req.CompanyWallet,
			"company_name":      company.CompanyName,
			"subscription_type": req.SubscriptionType,
			"subscription_end":  newEndTime.Unix(),
			"is_active":         true,
			"transaction_hash":  req.TransactionHash,
			"updated_at":        now.Format(time.RFC3339),
		},
	})
}

// CheckPayment 결제 확인 및 구독 업데이트
func (s *CompanyService) CheckPayment(c *gin.Context) {
	fmt.Printf("[DEBUG] CheckPayment 요청 시작\n")

	var req struct {
		CompanyWallet    string `json:"company_wallet" binding:"required"`
		SubscriptionType int    `json:"subscription_type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("[DEBUG] 요청 바인딩 실패: %v\n", err)
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "잘못된 요청 형식: " + err.Error(),
		})
		return
	}

	fmt.Printf("[DEBUG] 결제 확인 요청 데이터: %+v\n", req)

	// 1. 회사 정보 조회
	var company models.CompanyRegistered
	if err := setup.DB.Where("company_wallet = ?", req.CompanyWallet).First(&company).Error; err != nil {
		fmt.Printf("[DEBUG] 회사 정보 조회 실패: %v\n", err)
		c.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "등록되지 않은 회사입니다.",
		})
		return
	}

	// 2. 블록체인에서 결제 확인
	paymentConfirmed, err := s.checkBlockchainPayment(req.CompanyWallet, req.SubscriptionType)
	if err != nil {
		fmt.Printf("[DEBUG] 블록체인 결제 확인 실패: %v\n", err)
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "블록체인 결제 확인 실패: " + err.Error(),
		})
		return
	}

	if paymentConfirmed {
		// 3. 구독 기간 계산
		var subscriptionDuration time.Duration
		switch req.SubscriptionType {
		case 1:
			subscriptionDuration = 30 * 24 * time.Hour // 1개월
		case 2:
			subscriptionDuration = 90 * 24 * time.Hour // 3개월
		case 3:
			subscriptionDuration = 365 * 24 * time.Hour // 1년
		default:
			c.JSON(http.StatusBadRequest, models.Response{
				Success: false,
				Message: "잘못된 구독 유형입니다.",
			})
			return
		}

		// 4. 구독 기간 업데이트
		newEndTime := time.Now().Add(subscriptionDuration)
		company.SubscriptionEnd = newEndTime.Unix()
		company.SubscriptionType = uint8(req.SubscriptionType)

		if err := setup.DB.Save(&company).Error; err != nil {
			fmt.Printf("[DEBUG] 구독 업데이트 실패: %v\n", err)
			c.JSON(http.StatusInternalServerError, models.Response{
				Success: false,
				Message: "구독 업데이트 실패: " + err.Error(),
			})
			return
		}

		fmt.Printf("[DEBUG] 구독 업데이트 성공: %s, 새로운 만료일: %d\n", company.CompanyWallet, company.SubscriptionEnd)

		c.JSON(http.StatusOK, models.Response{
			Success: true,
			Message: "결제 확인 완료 및 구독 연장 성공",
			Data: map[string]interface{}{
				"company_wallet":    company.CompanyWallet,
				"subscription_type": req.SubscriptionType,
				"subscription_end":  company.SubscriptionEnd,
				"new_end_date":      newEndTime.Format("2006-01-02 15:04:05"),
			},
		})
	} else {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "결제가 확인되지 않았습니다. 잠시 후 다시 시도해주세요.",
		})
	}
}

// checkBlockchainPayment 블록체인에서 결제 확인
func (s *CompanyService) checkBlockchainPayment(fromAddress string, subscriptionType int) (bool, error) {
	fmt.Printf("[DEBUG] 블록체인 결제 확인 시작: %s, 구독유형: %d\n", fromAddress, subscriptionType)

	// 결제 주소
	paymentAddress := "0x578C1E3bE1FD168511618B72A6A7F080eDfa7445"

	// 결제 금액 (wei 단위)
	paymentAmounts := map[int]string{
		1: "1000000000000", // 0.000001 ETH
		2: "2500000000000", // 0.0000025 ETH
		3: "8000000000000", // 0.000008 ETH
	}

	expectedAmount := paymentAmounts[subscriptionType]
	if expectedAmount == "" {
		return false, fmt.Errorf("잘못된 구독 유형: %d", subscriptionType)
	}

	// blockchain 클라이언트를 통해 결제 확인
	return s.contractClient.CheckPayment(fromAddress, paymentAddress, expectedAmount)
}
