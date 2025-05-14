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

// 회사 등록 요청 구조체
type RegisterRequest struct {
	CompanyAddress string `json:"company_address" binding:"required"`
	CompanyName    string `json:"company_name" binding:"required"`
	IP             string `json:"ip" binding:"required"`
	ServerName     string `json:"server_name" binding:"required"`
	Port           string `json:"port" binding:"required"`
}

// 회사 등록 처리
func (s *CompanyService) RegisterCompanyInternal(req models.RegisterRequest) error {
	// 이더리움 주소 유효성 검사
	if !common.IsHexAddress(req.CompanyAddress) {
		return fmt.Errorf("유효하지 않은 이더리움 주소")
	}

	// 구독 상태 확인
	isSubscribed, err := s.contractClient.CheckSubscription(common.HexToAddress(req.CompanyAddress))
	if err != nil {
		return fmt.Errorf("구독 상태 확인 실패: %v", err)
	}

	if !isSubscribed {
		return fmt.Errorf("구독이 만료되었거나 활성화되지 않은 회사입니다")
	}

	// 서버 접근 정보 저장
	err = setup.SaveServerAccess(req.CompanyAddress, req.IP, req.ServerName, req.Port)
	if err != nil {
		return fmt.Errorf("서버 접근 정보 저장 실패: %v", err)
	}

	return nil
}

// 회사 정보 조회
func (s *CompanyService) GetCompanyInfo(address string) (*models.ServerAccess, error) {
	return setup.GetServerAccess(address)
}

// RegisterCompanyHandler Gin 핸들러
func (s *CompanyService) RegisterCompanyHandler(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "잘못된 요청 형식: " + err.Error(),
		})
		return
	}

	if err := s.RegisterCompanyInternal(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "회원가입이 완료되었습니다",
		"data": map[string]interface{}{
			"company_address": req.CompanyAddress,
			"company_name":    req.CompanyName,
			"registered_at":   time.Now().Format(time.RFC3339),
		},
	})
}
