package service

import (
	"fmt"
	"net/http"
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

// 회사 등록 처리
func (s *CompanyService) RegisterCompanyInternal(req models.RegisterRequest) error {
	// 1. 이더리움 주소 유효성 검사
	if !common.IsHexAddress(req.CompanyWallet) {
		return fmt.Errorf("유효하지 않은 이더리움 주소")
	}

	// 2. 스마트 컨트랙트를 통한 회사 등록 트랜잭션 실행
	tx, err := s.contractClient.RegisterCompany(
		req.CompanyName,
		req.CeoName,
		req.Email,
		req.SubscriptionType,
	)
	if err != nil {
		return fmt.Errorf("회사 등록 트랜잭션 실패: %v", err)
	}

	// 3. 트랜잭션 영수증 확인
	receipt, err := s.contractClient.WaitForTransaction(tx)
	if err != nil {
		return fmt.Errorf("트랜잭션 확인 실패: %v", err)
	}

	// 4. 트랜잭션이 성공했는지 확인
	if receipt.Status != 1 {
		return fmt.Errorf("트랜잭션이 실패했습니다")
	}

	// 5. 서버 접근 정보 저장
	err = setup.SaveServerAccess(req.CompanyWallet, req.IP, req.ServerName, req.Port)
	if err != nil {
		return fmt.Errorf("서버 접근 정보 저장 실패: %v", err)
	}

	return nil
}

// 회사 정보 조회
func (s *CompanyService) GetCompanyInfo(address string) (*models.CompanyRegistered, error) {
	return setup.GetServerAccess(address)
}

// RegisterCompany Gin 핸들러
func (s *CompanyService) RegisterCompany(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "잘못된 요청 형식: " + err.Error(),
		})
		return
	}

	if err := s.RegisterCompanyInternal(req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

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
