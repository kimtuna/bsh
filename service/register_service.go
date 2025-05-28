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
	"golang.org/x/crypto/ssh"
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
func testServerAccess(ip string, port uint16) error {
	// SSH 연결 테스트
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(), // SSH 키 인증
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second, // 5초 타임아웃
	}

	// 서버 연결 시도
	_, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", ip, port), config)
	if err != nil {
		return fmt.Errorf("서버에 접근할 수 없습니다 (IP: %s, Port: %d): %v", ip, port, err)
	}

	return nil
}

// RegisterCompanyInternal 회사 등록 처리
func (s *CompanyService) RegisterCompanyInternal(req models.RegisterRequest) error {
	// 1. 서버 접근 테스트
	if err := testServerAccess(req.IP, req.Port); err != nil {
		return fmt.Errorf("서버 접근 실패: %v", err)
	}

	// 2. 이더리움 주소 유효성 검사
	if !common.IsHexAddress(req.CompanyWallet) {
		return fmt.Errorf("유효하지 않은 이더리움 주소")
	}

	// 3. 스마트 컨트랙트를 통한 회사 등록 트랜잭션 실행
	tx, err := s.contractClient.RegisterCompany(
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
