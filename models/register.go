package models

import (
	"time"
)

type RegisterRequest struct {
	CompanyAddress   string `json:"company_address" binding:"required"`
	CompanyName      string `json:"company_name" binding:"required"`
	CeoName          string `json:"ceo_name" binding:"required"`
	SubscriptionType uint8  `json:"subscription_type" binding:"required,min=1,max=3"` // 1: 1개월, 2: 3개월, 3: 1년
	IP               string `json:"ip" binding:"required"`
	ServerName       string `json:"server_name" binding:"required"`
	Port             string `json:"port" binding:"required"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ServerAccess struct {
	CompanyAddress string `gorm:"primaryKey"`
	IP             string
	ServerName     string
	Port           string
	IsActive       bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
