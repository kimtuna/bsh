package service

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"

	"log"

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
func (s *CompanyService) RegisterCompanyInternal(req models.RegisterRequest) error {
	// 1. 서버 접근 테스트
	if err := testServerAccess(req.IP, req.Port, req.ServerName, req.Password); err != nil {
		return fmt.Errorf("서버 접근 실패: %v", err)
	}

	// 2. 이더리움 주소 유효성 검사
	if !common.IsHexAddress(req.CompanyWallet) {
		return fmt.Errorf("유효하지 않은 이더리움 주소")
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
		return fmt.Errorf("회사 등록 트랜잭션 실패: %v", err)
	}

	// 4. 트랜잭션 영수증 확인
	receipt, err := s.contractClient.WaitForTransaction(tx)
	if err != nil {
		return fmt.Errorf("트랜잭션 확인 실패: %v", err)
	}

	// 5. 트랜잭션이 성공했는지 확인
	if receipt.Status != 1 {
		return fmt.Errorf("트랜잭션이 실패했습니다")
	}

	// 6. 서버 접근 정보 저장
	err = setup.SaveServerAccess(req.CompanyWallet, req.Email, req.IP, req.ServerName, req.Port)
	if err != nil {
		return fmt.Errorf("서버 접근 정보 저장 실패: %v", err)
	}

	return nil
}

// RegisterCompany Gin 핸들러
func (s *CompanyService) RegisterCompany(c *gin.Context) {
	fmt.Printf("[DEBUG] RegisterCompany 요청 시작\n")

	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("[DEBUG] JSON 바인딩 오류: %v\n", err)
		if body, err := c.GetRawData(); err == nil {
			fmt.Printf("[DEBUG] 요청 본문: %s\n", string(body))
		}
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "잘못된 요청 형식: " + err.Error(),
		})
		return
	}

	fmt.Printf("[DEBUG] 요청 데이터: %+v\n", req)

	// 서버 연결 테스트 (직접 ping 사용)
	if err := testServerAccess(req.IP, req.Port, req.ServerName, req.Password); err != nil {
		log.Printf("Server connection test failed: %v", err)
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := s.RegisterCompanyInternal(req); err != nil {
		fmt.Printf("[DEBUG] 회사 등록 실패: %v\n", err)
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	fmt.Printf("[DEBUG] 회사 등록 성공: %s\n", req.CompanyWallet)
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "회원가입이 완료되었습니다",
		Data: map[string]interface{}{
			"company_wallet": req.CompanyWallet,
			"company_name":   req.CompanyName,
			"registered_at":  time.Now().Format(time.RFC3339),
		},
	})
}
