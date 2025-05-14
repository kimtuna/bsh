package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

)

// 응답 구조체
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func main() {
	// 데이터베이스 연결
	if err := ConnectDatabase(); err != nil {
		log.Fatal("데이터베이스 연결 실패:", err)
	}

	// 컨트랙트 클라이언트 초기화
	contractClient, err := NewContractClient()
	if err != nil {
		log.Fatal("컨트랙트 클라이언트 초기화 실패:", err)
	}

	// 서비스 초기화
	companyService := service.NewCompanyService(contractClient)

	// Gin 라우터 초기화
	r := gin.Default()

	// CORS 미들웨어
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 회원가입 엔드포인트
	r.POST("/api/register", func(c *gin.Context) {
		var req service.RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "잘못된 요청 형식: " + err.Error(),
			})
			return
		}

		err := companyService.RegisterCompany(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, Response{
			Success: true,
			Message: "회원가입이 완료되었습니다",
			Data: map[string]interface{}{
				"company_address": req.CompanyAddress,
				"company_name":    req.CompanyName,
				"registered_at":   time.Now().Format(time.RFC3339),
			},
		})
	})

	// 서버 시작
	port := ":8080"
	log.Printf("API 서버가 %s 포트에서 시작됩니다...\n", port)
	if err := r.Run(port); err != nil {
		log.Fatal("서버 시작 실패:", err)
	}
}
