package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func MonitorEvents() {
	// ContractClient 초기화
	client, err := NewContractClient()
	if err != nil {
		log.Fatal("컨트랙트 클라이언트 초기화 실패:", err)
	}

	// 이벤트 채널 생성
	eventChan := make(chan CompanyEvent, 100)

	// 컨텍스트 생성
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 이벤트 구독 시작
	err = client.SubscribeToEvents(ctx, eventChan)
	if err != nil {
		log.Fatal("이벤트 구독 실패:", err)
	}

	// 종료 시그널 처리
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("이벤트 모니터링을 시작합니다... (종료하려면 Ctrl+C를 누르세요)")

	// 이벤트 처리 루프
	for {
		select {
		case <-sigChan:
			fmt.Println("\n모니터링을 종료합니다...")
			return
		case event := <-eventChan:
			handleEvent(event)
		}
	}
}

func handleEvent(event CompanyEvent) {
	timestamp := time.Unix(int64(event.Timestamp), 0).Format("2006-01-02 15:04:05")

	switch event.EventType {
	case "CompanyRegistered":
		fmt.Printf("[%s] 새로운 회사 등록\n", timestamp)
		fmt.Printf("  - 회사 주소: %s\n", event.CompanyAddress.Hex())
		if data, ok := event.Data["name"].(string); ok {
			fmt.Printf("  - 회사명: %s\n", data)
		}
		if data, ok := event.Data["subscriptionEnd"].(*big.Int); ok {
			fmt.Printf("  - 구독 만료일: %s\n", time.Unix(data.Int64(), 0).Format("2006-01-02 15:04:05"))
		}

	case "SubscriptionExtended":
		fmt.Printf("[%s] 구독 연장\n", timestamp)
		fmt.Printf("  - 회사 주소: %s\n", event.CompanyAddress.Hex())
		if data, ok := event.Data["newEndDate"].(*big.Int); ok {
			fmt.Printf("  - 새로운 만료일: %s\n", time.Unix(data.Int64(), 0).Format("2006-01-02 15:04:05"))
		}

	case "ServiceExpired":
		fmt.Printf("[%s] 서비스 만료\n", timestamp)
		fmt.Printf("  - 회사 주소: %s\n", event.CompanyAddress.Hex())
	}
}
