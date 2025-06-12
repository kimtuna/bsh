package models

type LoginRequest struct {
	CompanyWallet string `json:"company_wallet" binding:"required"`
}

type LoginResponse struct {
	IP               string `json:"ip"`
	ServerName       string `json:"server_name"`
	Port             uint16 `json:"port"`
	SSHCommand       string `json:"ssh_command"` // ssh user@ip -p port 형식
	CompanyName      string `json:"company_name,omitempty"`
	SubscriptionType uint8  `json:"subscription_type,omitempty"`
	SubscriptionEnd  int64  `json:"subscription_end,omitempty"`
	RemainingDays    int    `json:"remaining_days,omitempty"`
	IsExpired        bool   `json:"is_expired,omitempty"`
}
