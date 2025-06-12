package service

import (
	"fmt"
	"net/http"
	"time"

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

	// 3. 구독 만료 여부 확인
	isExpired, err := setup.CheckSubscriptionExpired(req.CompanyWallet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "구독 상태 확인 실패: " + err.Error(),
		})
		return
	}

	// 4. 구독 정보 조회
	subscriptionInfo, err := setup.GetSubscriptionInfo(req.CompanyWallet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "구독 정보 조회 실패: " + err.Error(),
		})
		return
	}

	// 5. 남은 일수 계산
	currentTime := time.Now().Unix()
	remainingDays := int((subscriptionInfo.SubscriptionEnd - currentTime) / (24 * 60 * 60))

	if isExpired {
		c.JSON(http.StatusPaymentRequired, models.Response{
			Success: false,
			Message: "구독이 만료되었습니다. 구독을 연장해주세요.",
			Data: map[string]interface{}{
				"company_wallet":    subscriptionInfo.CompanyWallet,
				"company_name":      subscriptionInfo.CompanyName,
				"subscription_type": subscriptionInfo.SubscriptionType,
				"subscription_end":  subscriptionInfo.SubscriptionEnd,
				"is_expired":        true,
				"remaining_days":    0,
				"suggestion":        "bsh charge 명령어로 구독을 연장하세요.",
			},
		})
		return
	}

	// 6. SSH 접속 명령어 생성
	sshCommand := fmt.Sprintf("ssh %s@%s -p %d", company.ServerName, company.IP, company.Port)

	// 7. 응답 반환
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "로그인 성공",
		Data: models.LoginResponse{
			IP:         company.IP,
			ServerName: company.ServerName,
			Port:       company.Port,
			SSHCommand: sshCommand,
			// 구독 정보 추가
			CompanyName:      subscriptionInfo.CompanyName,
			SubscriptionType: subscriptionInfo.SubscriptionType,
			SubscriptionEnd:  subscriptionInfo.SubscriptionEnd,
			RemainingDays:    remainingDays,
			IsExpired:        false,
		},
	})
}
