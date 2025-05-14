package service

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type CompanyService struct {
	contractClient *ContractClient
}

func NewCompanyService(client *ContractClient) *CompanyService {
	return &CompanyService{
		contractClient: client,
	}
}

// 회사 등록 요청 구조체
type RegisterRequest struct {
	CompanyAddress string `json:"company_address" binding:"required"`
	CompanyName    string `json:"company_name" binding:"required"`
	IP             string `json:"ip" binding:"required"`
	ServerName     string `json:"server_name" binding:"required"`
	Port           string `json:"port" binding:"required"`
}

// 회사 등록 처리
func (s *CompanyService) RegisterCompany(req RegisterRequest) error {
	// 이더리움 주소 유효성 검사
	if !common.IsHexAddress(req.CompanyAddress) {
		return fmt.Errorf("유효하지 않은 이더리움 주소")
	}

	// 구독 상태 확인
	isSubscribed, err := s.contractClient.CheckSubscription(common.HexToAddress(req.CompanyAddress))
	if err != nil {
		return fmt.Errorf("구독 상태 확인 실패: %v", err)
	}

	if !isSubscribed {
		return fmt.Errorf("구독이 만료되었거나 활성화되지 않은 회사입니다")
	}

	// 서버 접근 정보 저장
	err = SaveServerAccess(req.CompanyAddress, req.IP, req.ServerName, req.Port)
	if err != nil {
		return fmt.Errorf("서버 접근 정보 저장 실패: %v", err)
	}

	return nil
}

// 회사 정보 조회
func (s *CompanyService) GetCompanyInfo(address string) (*ServerAccess, error) {
	return GetServerAccess(address)
}
