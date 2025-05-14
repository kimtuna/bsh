package models

// RegisterRequest 회사 등록 요청 구조체
type RegisterRequest struct {
	CompanyAddress string `json:"company_address" binding:"required"` // 회사의 이더리움 주소
	CompanyName    string `json:"company_name" binding:"required"`    // 회사명
	IP             string `json:"ip" binding:"required"`              // 서버 IP
	ServerName     string `json:"server_name" binding:"required"`     // 서버 이름
	Port           string `json:"port" binding:"required"`            // 서버 포트
}

// Response API 응답 구조체
type Response struct {
	Success bool        `json:"success"`        // 요청 성공 여부
	Message string      `json:"message"`        // 응답 메시지
	Data    interface{} `json:"data,omitempty"` // 응답 데이터 (선택적)
}

// ServerAccess 서버 접근 정보 구조체
type ServerAccess struct {
	ID             uint   `gorm:"primaryKey"`           // 기본 키
	CompanyAddress string `gorm:"uniqueIndex;not null"` // 회사의 이더리움 주소
	IP             string `gorm:"not null"`             // 서버 IP
	ServerName     string `gorm:"not null"`             // 서버 이름
	Port           string `gorm:"not null"`             // 서버 포트
	IsActive       bool   `gorm:"default:true"`         // 활성화 상태
	CreatedAt      int64  `gorm:"autoCreateTime"`       // 생성 시간
	UpdatedAt      int64  `gorm:"autoUpdateTime"`       // 수정 시간
	DeletedAt      int64  `gorm:"index"`                // 삭제 시간 (소프트 삭제)
}
