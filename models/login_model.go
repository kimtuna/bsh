package models

type LoginRequest struct {
	CompanyWallet string `json:"company_wallet" binding:"required"`
}

type LoginResponse struct {
	IP         string `json:"ip"`
	ServerName string `json:"server_name"`
	Port       uint16 `json:"port"`
	SSHCommand string `json:"ssh_command"` // ssh user@ip -p port 형식
}
