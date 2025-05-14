package service

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/kimtuna/bsh/models"
	"github.com/kimtuna/bsh/setup"
)

// Login 회사 로그인 처리
func (s *CompanyService) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "잘못된 요청 형식: " + err.Error(),
		})
		return
	}

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
		c.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "등록되지 않은 회사입니다",
		})
		return
	}

	// 3. SSH 접속 명령어 생성
	sshCommand := fmt.Sprintf("ssh root@%s -p %d", company.IP, company.Port)

	// 4. 응답 반환
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "로그인 성공",
		Data: models.LoginResponse{
			IP:         company.IP,
			ServerName: company.ServerName,
			Port:       company.Port,
			SSHCommand: sshCommand,
		},
	})
}
