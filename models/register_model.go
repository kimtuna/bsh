package models

import (
	"time"
)

type RegisterRequest struct {
	CompanyWallet    string `json:"company_wallet" binding:"required"`
	CompanyName      string `json:"company_name" binding:"required"`
	CeoName          string `json:"ceo_name" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	SubscriptionType uint8  `json:"subscription_type" binding:"required,min=1,max=3"` // 1: 1개월, 2: 3개월, 3: 1년
	IP               string `json:"ip" binding:"required"`
	ServerName       string `json:"server_name" binding:"required"`
	Port             uint16 `json:"port" binding:"required,min=1,max=65535"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type CompanyRegistered struct {
	CompanyWallet string `gorm:"primaryKey;column:company_wallet;type:varchar(191)"` // 이더리움 지갑 주소
	Email         string `gorm:"uniqueIndex:email;type:varchar(50)"`
	IP            string `gorm:"type:varchar(50)"`
	ServerName    string `gorm:"type:varchar(30)"`
	Port          uint16 `gorm:"type:smallint unsigned"`
	IsActive      bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
