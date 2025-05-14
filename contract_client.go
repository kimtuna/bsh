package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

type ContractClient struct {
	client       *ethclient.Client
	contractAbi  abi.ABI
	contractAddr common.Address
}

// 이벤트 로그를 위한 구조체
type CompanyEvent struct {
	CompanyAddress common.Address
	EventType      string // "CompanyRegistered", "SubscriptionExtended", "ServiceExpired"
	Timestamp      uint64
	Data           map[string]interface{}
}

func NewContractClient() (*ContractClient, error) {
	// .env 파일 로드
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	infuraURL := os.Getenv("INFURA_URL")
	contractAddrStr := os.Getenv("CONTRACT_ADDRESS")

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

	return &ContractClient{
		client:       client,
		contractAbi:  contractAbi,
		contractAddr: contractAddress,
	}, nil
}

func (c *ContractClient) GetCompanyInfo(companyAddr common.Address) (string, string, *big.Int, bool, error) {
	data, err := c.contractAbi.Pack("getCompanyInfo", companyAddr)
	if err != nil {
		return "", "", nil, false, fmt.Errorf("error packing getCompanyInfo: %v", err)
	}

	callMsg := ethereum.CallMsg{
		To:   &c.contractAddr,
		Data: data,
	}

	result, err := c.client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return "", "", nil, false, fmt.Errorf("error calling contract: %v", err)
	}

	var (
		name            string
		ceoName         string
		subscriptionEnd *big.Int
		isActive        bool
	)

	err = c.contractAbi.UnpackIntoInterface(&[]interface{}{&name, &ceoName, &subscriptionEnd, &isActive}, "getCompanyInfo", result)
	if err != nil {
		return "", "", nil, false, fmt.Errorf("error unpacking result: %v", err)
	}

	return name, ceoName, subscriptionEnd, isActive, nil
}

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
							"subscriptionEnd": data[1],
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
