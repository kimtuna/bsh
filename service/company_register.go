package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kimtuna/bsh/models"
)

func (s *CompanyService) RegisterCompany(c *gin.Context) {
	// 회사 정보 파싱
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "잘못된 요청 형식: " + err.Error(),
		})
		return
	}

	// 스마트 컨트랙트를 통한 회사 등록 및 블록체인 기록
	if err := s.RegisterCompanyInternal(req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// 3. 블록체인에 기록된 회사 정보 응답
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "회원가입이 완료되었습니다",
		Data: map[string]interface{}{
			"company_address": req.CompanyAddress,              // 블록체인에 기록된 회사 주소
			"company_name":    req.CompanyName,                 // 등록된 회사명
			"registered_at":   time.Now().Format(time.RFC3339), // 블록 생성 시간
		},
	})
}
