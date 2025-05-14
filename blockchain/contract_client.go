package blockchain

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/kimtuna/bsh/contracts"
)

// ContractClient 스마트 컨트랙트와 상호작용하는 클라이언트
type ContractClient struct {
	client       *ethclient.Client
	contractAbi  abi.ABI
	contractAddr common.Address
	contract     *contracts.CompanyServerAccess
	auth         *bind.TransactOpts
}

// CompanyEvent 컨트랙트 이벤트 로그를 위한 구조체
type CompanyEvent struct {
	CompanyAddress common.Address
	EventType      string // "CompanyRegistered", "SubscriptionExtended", "ServiceExpired"
	Timestamp      uint64
	Data           map[string]interface{}
}

// NewContractClient 새로운 컨트랙트 클라이언트 생성
func NewContractClient() (*ContractClient, error) {
	// .env 파일 로드
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	infuraURL := os.Getenv("INFURA_URL")
	contractAddrStr := os.Getenv("CONTRACT_ADDRESS")
	privateKeyHex := os.Getenv("PRIVATE_KEY")

	// RPC 연결
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		return nil, fmt.Errorf("error connecting to ethereum client: %v", err)
	}

	// ABI 파일 읽기
	abiBytes, err := os.ReadFile("abi.json")
	if err != nil {
		return nil, fmt.Errorf("error reading abi file: %v", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		return nil, fmt.Errorf("error parsing abi: %v", err)
	}

	contractAddress := common.HexToAddress(contractAddrStr)

	// 컨트랙트 인스턴스 생성
	contract, err := contracts.NewCompanyServerAccess(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("error creating contract instance: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error parsing private key: %v", err)
	}

	chainIDStr := os.Getenv("CHAIN_ID")
	chainID, _ := new(big.Int).SetString(chainIDStr, 10)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("error creating auth: %v", err)
	}

	return &ContractClient{
		client:       client,
		contractAbi:  contractAbi,
		contractAddr: contractAddress,
		contract:     contract,
		auth:         auth,
	}, nil
}

// GetCompanyInfo 회사 정보 조회
func (c *ContractClient) GetCompanyInfo(companyAddr common.Address) (string, string, string, *big.Int, bool, error) {
	data, err := c.contractAbi.Pack("getCompanyInfo", companyAddr)
	if err != nil {
		return "", "", "", nil, false, fmt.Errorf("error packing getCompanyInfo: %v", err)
	}

	callMsg := ethereum.CallMsg{
		To:   &c.contractAddr,
		Data: data,
	}

	result, err := c.client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return "", "", "", nil, false, fmt.Errorf("error calling contract: %v", err)
	}

	var (
		name            string
		ceoName         string
		email           string
		subscriptionEnd *big.Int
		isActive        bool
	)

	err = c.contractAbi.UnpackIntoInterface(&[]interface{}{&name, &ceoName, &email, &subscriptionEnd, &isActive}, "getCompanyInfo", result)
	if err != nil {
		return "", "", "", nil, false, fmt.Errorf("error unpacking result: %v", err)
	}

	return name, ceoName, email, subscriptionEnd, isActive, nil
}

// CheckSubscription 구독 상태 확인
func (c *ContractClient) CheckSubscription(companyAddr common.Address) (bool, error) {
	data, err := c.contractAbi.Pack("checkSubscription", companyAddr)
	if err != nil {
		return false, fmt.Errorf("error packing checkSubscription: %v", err)
	}

	callMsg := ethereum.CallMsg{
		To:   &c.contractAddr,
		Data: data,
	}

	result, err := c.client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return false, fmt.Errorf("error calling contract: %v", err)
	}

	var isSubscribed bool
	err = c.contractAbi.UnpackIntoInterface(&[]interface{}{&isSubscribed}, "checkSubscription", result)
	if err != nil {
		return false, fmt.Errorf("error unpacking result: %v", err)
	}

	return isSubscribed, nil
}

// SubscribeToEvents 컨트랙트 이벤트 구독
func (c *ContractClient) SubscribeToEvents(ctx context.Context, eventChan chan<- CompanyEvent) error {
	// 이벤트 시그니처 해시
	companyRegisteredSig := []byte("CompanyRegistered(address,string,uint256)")
	subscriptionExtendedSig := []byte("SubscriptionExtended(address,uint256)")
	serviceExpiredSig := []byte("ServiceExpired(address)")

	companyRegisteredHash := common.BytesToHash(crypto.Keccak256(companyRegisteredSig))
	subscriptionExtendedHash := common.BytesToHash(crypto.Keccak256(subscriptionExtendedSig))
	serviceExpiredHash := common.BytesToHash(crypto.Keccak256(serviceExpiredSig))

	// 이벤트 필터 쿼리
	query := ethereum.FilterQuery{
		Addresses: []common.Address{c.contractAddr},
		Topics: [][]common.Hash{
			{
				companyRegisteredHash,
				subscriptionExtendedHash,
				serviceExpiredHash,
			},
		},
	}

	// 이벤트 구독
	logs := make(chan types.Log)
	sub, err := c.client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return fmt.Errorf("error subscribing to logs: %v", err)
	}

	go func() {
		defer sub.Unsubscribe()

		for {
			select {
			case err := <-sub.Err():
				fmt.Printf("구독 에러: %v\n", err)
				return
			case log := <-logs:
				event := CompanyEvent{
					CompanyAddress: common.HexToAddress(log.Topics[1].Hex()),
					Timestamp:      log.BlockNumber,
				}

				switch log.Topics[0] {
				case companyRegisteredHash:
					event.EventType = "CompanyRegistered"
					// 이벤트 데이터 파싱
					data, err := c.contractAbi.Unpack("CompanyRegistered", log.Data)
					if err == nil {
						event.Data = map[string]interface{}{
							"name":            data[0],
							"email":           data[1],
							"subscriptionEnd": data[2],
						}
					}
				case subscriptionExtendedHash:
					event.EventType = "SubscriptionExtended"
					data, err := c.contractAbi.Unpack("SubscriptionExtended", log.Data)
					if err == nil {
						event.Data = map[string]interface{}{
							"newEndDate": data[0],
						}
					}
				case serviceExpiredHash:
					event.EventType = "ServiceExpired"
				}

				eventChan <- event
			}
		}
	}()

	return nil
}

// RegisterCompany 회사 등록 트랜잭션 실행
func (c *ContractClient) RegisterCompany(name, ceoName, email string, subscriptionType uint8) (*types.Transaction, error) {
	// 구독 가격 계산
	price, err := c.contract.GetSubscriptionPrice(nil, big.NewInt(int64(subscriptionType)))
	if err != nil {
		return nil, fmt.Errorf("구독 가격 조회 실패: %v", err)
	}

	// 트랜잭션 옵션 설정
	c.auth.Value = price

	// 회사 등록 트랜잭션 실행
	tx, err := c.contract.RegisterCompany(c.auth, name, ceoName, email, big.NewInt(int64(subscriptionType)))
	if err != nil {
		return nil, fmt.Errorf("회사 등록 트랜잭션 실패: %v", err)
	}

	return tx, nil
}

// WaitForTransaction 트랜잭션 영수증 대기
func (c *ContractClient) WaitForTransaction(tx *types.Transaction) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, c.client, tx)
	if err != nil {
		return nil, fmt.Errorf("트랜잭션 대기 실패: %v", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return nil, fmt.Errorf("트랜잭션이 실패했습니다")
	}

	return receipt, nil
}
