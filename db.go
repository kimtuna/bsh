package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ServerAccess 모델 - 회사별 서버 접근 정보
type ServerAccess struct {
	gorm.Model
	CompanyAddress string `gorm:"uniqueIndex;not null"` // 이더리움 주소
	IP             string `gorm:"not null"`
	ServerName     string `gorm:"not null"`
	Port           string `gorm:"not null"`
	IsActive       bool   `gorm:"default:true"`
}

func ConnectDatabase() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	// MySQL 관련 환경 변수 설정
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbName := os.Getenv("DB_NAME")

	// DSN(Data Source Name) 생성
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DbUser, DbPassword, DbHost, DbPort, DbName)

	// 데이터베이스 연결
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("cannot connect to database: %v", err)
	}

	// 테이블 자동 마이그레이션
	if err := DB.AutoMigrate(&ServerAccess{}); err != nil {
		return fmt.Errorf("auto migrate failed: %v", err)
	}

	log.Println("데이터베이스 연결 및 마이그레이션 완료")
	return nil
}

// 서버 접근 정보 저장
func SaveServerAccess(companyAddr, ip, serverName, port string) error {
	access := ServerAccess{
		CompanyAddress: companyAddr,
		IP:             ip,
		ServerName:     serverName,
		Port:           port,
		IsActive:       true,
	}

	result := DB.Create(&access)
	if result.Error != nil {
		return fmt.Errorf("서버 접근 정보 저장 실패: %v", result.Error)
	}

	return nil
}

// 회사 주소로 서버 접근 정보 조회
func GetServerAccess(companyAddr string) (*ServerAccess, error) {
	var access ServerAccess
	result := DB.Where("company_address = ? AND is_active = ?", companyAddr, true).First(&access)
	if result.Error != nil {
		return nil, fmt.Errorf("서버 접근 정보 조회 실패: %v", result.Error)
	}

	return &access, nil
}

// 서버 접근 정보 비활성화
func DeactivateServerAccess(companyAddr string) error {
	result := DB.Model(&ServerAccess{}).
		Where("company_address = ?", companyAddr).
		Update("is_active", false)

	if result.Error != nil {
		return fmt.Errorf("서버 접근 정보 비활성화 실패: %v", result.Error)
	}

	return nil
}
