package setup

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kimtuna/bsh/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase 데이터베이스 연결
func ConnectDatabase() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// 테이블 자동 마이그레이션
	err = DB.AutoMigrate(&models.ServerAccess{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}

	log.Println("데이터베이스 연결 성공")
	return nil
}

// SaveServerAccess 서버 접근 정보 저장
func SaveServerAccess(companyAddress, ip, serverName, port string) error {
	serverAccess := &models.ServerAccess{
		CompanyAddress: companyAddress,
		IP:             ip,
		ServerName:     serverName,
		Port:           port,
		IsActive:       true,
	}

	result := DB.Create(serverAccess)
	if result.Error != nil {
		return fmt.Errorf("failed to save server access: %v", result.Error)
	}

	return nil
}

// GetServerAccess 회사 주소로 서버 접근 정보 조회
func GetServerAccess(companyAddress string) (*models.ServerAccess, error) {
	var serverAccess models.ServerAccess
	result := DB.Where("company_address = ? AND is_active = ?", companyAddress, true).First(&serverAccess)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get server access: %v", result.Error)
	}

	return &serverAccess, nil
}

// DeactivateServerAccess 서버 접근 비활성화
func DeactivateServerAccess(companyAddress string) error {
	result := DB.Model(&models.ServerAccess{}).
		Where("company_address = ? AND is_active = ?", companyAddress, true).
		Update("is_active", false)
	if result.Error != nil {
		return fmt.Errorf("failed to deactivate server access: %v", result.Error)
	}

	return nil
}
