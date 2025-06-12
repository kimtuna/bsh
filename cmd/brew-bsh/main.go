package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var (
	rootCmd = &cobra.Command{
		Use:   "bsh",
		Short: "BSH - Blockchain Server Hosting CLI",
		Long:  `BSH는 블록체인 기반의 서버 호스팅 서비스를 위한 CLI 도구입니다.`,
	}

	config *Config
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "회사 등록",
	Run: func(cmd *cobra.Command, args []string) {
		register()
	},
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "서버 로그인",
	Run: func(cmd *cobra.Command, args []string) {
		login()
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "구독 상태 조회",
	Run: func(cmd *cobra.Command, args []string) {
		status()
	},
}

var chargeCmd = &cobra.Command{
	Use:   "charge",
	Short: "구독 연장 (메타마스크 결제)",
	Run: func(cmd *cobra.Command, args []string) {
		charge()
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(chargeCmd)

	// 설정 파일 로드
	var err error
	config, err = loadConfig()
	if err != nil {
		color.Red("설정 파일 로드 실패:", err)
		os.Exit(1)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getInput(prompt string, isPassword bool) string {
	fmt.Print(prompt)
	if isPassword {
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println("\n입력 오류:", err)
			os.Exit(1)
		}
		fmt.Println()
		return string(bytePassword)
	}

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("입력 오류:", err)
		os.Exit(1)
	}
	return strings.TrimSpace(input)
}

func getServerURL(path string) string {
	return fmt.Sprintf("https://www.%s%s%s", config.Server.Host, config.Server.BasePath, path)
}

func isCommandAvailable(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func register() {
	color.Blue("=== BSH 회사 등록 ===")

	companyWallet := getInput("회사 지갑 주소: ", false)
	companyName := getInput("회사 이름: ", false)
	ceoName := getInput("대표자 이름: ", false)
	email := getInput("이메일: ", false)

	fmt.Println("\n구독 유형을 선택하세요:")
	fmt.Println("1. 1개월 (0.000001 ETH)")
	fmt.Println("2. 3개월 (0.0000025 ETH)")
	fmt.Println("3. 1년 (0.000008 ETH)")
	subscriptionType := getInput("구독 유형 (1-3): ", false)

	ip := getInput("서버 IP: ", false)
	serverName := getInput("서버 이름: ", false)
	port := getInput("서버 포트 (기본값: 22): ", false)
	if port == "" {
		port = "22"
	}
	password := getInput("서버 비밀번호: ", true)

	// API 요청 데이터 구성
	jsonData := fmt.Sprintf(`{
		"company_wallet": "%s",
		"company_name": "%s",
		"ceo_name": "%s",
		"email": "%s",
		"subscription_type": %s,
		"ip": "%s",
		"server_name": "%s",
		"port": %s,
		"password": "%s"
	}`, companyWallet, companyName, ceoName, email, subscriptionType, ip, serverName, port, password)

	// API 요청 전송
	resp, err := http.Post(getServerURL("/register"), "application/json", strings.NewReader(jsonData))
	if err != nil {
		color.Red("서버 연결 오류:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// 응답 처리
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		color.Red("응답 읽기 오류:", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		color.Red("등록 실패: HTTP %d", resp.StatusCode)
		color.Red("응답 내용: %s", string(body))
		os.Exit(1)
	}

	color.Green("회원가입이 완료되었습니다!")
}

func login() {
	color.Blue("=== BSH 로그인 ===")

	companyWallet := getInput("회사 지갑 주소: ", false)

	// API 요청 전송
	resp, err := http.Post(getServerURL("/login"), "application/json", strings.NewReader(fmt.Sprintf(`{"company_wallet": "%s"}`, companyWallet)))
	if err != nil {
		color.Red("서버 연결 오류:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// 응답 처리
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		color.Red("응답 읽기 오류:", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		// 구독 만료 시 특별 처리
		if resp.StatusCode == http.StatusPaymentRequired {
			var errorResp struct {
				Success bool   `json:"success"`
				Message string `json:"message"`
				Data    struct {
					CompanyName      string `json:"company_name"`
					SubscriptionType uint8  `json:"subscription_type"`
					RemainingDays    int    `json:"remaining_days"`
					Suggestion       string `json:"suggestion"`
				} `json:"data"`
			}

			if err := json.Unmarshal(body, &errorResp); err == nil {
				color.Red("⚠️  구독이 만료되었습니다!")
				color.Cyan("📊 구독 정보:")
				color.Cyan("- 회사: %s", errorResp.Data.CompanyName)
				color.Cyan("- 남은 일수: %d일", errorResp.Data.RemainingDays)
				color.Yellow("\n💡 %s", errorResp.Data.Suggestion)
			} else {
				color.Red("로그인 실패: %s", string(body))
			}
		} else {
			color.Red("로그인 실패: %s", string(body))
		}
		os.Exit(1)
	}

	var loginResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    struct {
			IP               string `json:"ip"`
			ServerName       string `json:"server_name"`
			Port             uint16 `json:"port"`
			SSHCommand       string `json:"ssh_command"`
			CompanyName      string `json:"company_name,omitempty"`
			SubscriptionType uint8  `json:"subscription_type,omitempty"`
			SubscriptionEnd  int64  `json:"subscription_end,omitempty"`
			RemainingDays    int    `json:"remaining_days,omitempty"`
			IsExpired        bool   `json:"is_expired,omitempty"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &loginResp); err != nil {
		color.Red("응답 파싱 오류:", err)
		os.Exit(1)
	}

	if !loginResp.Success {
		color.Red("로그인 실패: %s", loginResp.Message)
		os.Exit(1)
	}

	color.Green("로그인 성공!")

	// 구독 정보 표시
	if loginResp.Data.CompanyName != "" {
		color.Cyan("📊 구독 정보:")
		color.Cyan("- 회사: %s", loginResp.Data.CompanyName)

		subscriptionTypeText := ""
		switch loginResp.Data.SubscriptionType {
		case 1:
			subscriptionTypeText = "1개월"
		case 2:
			subscriptionTypeText = "3개월"
		case 3:
			subscriptionTypeText = "1년"
		default:
			subscriptionTypeText = "알 수 없음"
		}
		color.Cyan("- 구독 유형: %s", subscriptionTypeText)
		color.Cyan("- 남은 일수: %d일", loginResp.Data.RemainingDays)

		if loginResp.Data.RemainingDays <= 7 {
			color.Yellow("⚠️  구독이 곧 만료됩니다! (%d일 남음)", loginResp.Data.RemainingDays)
			color.Yellow("💡 구독을 연장하려면: bsh charge")
		}
	}

	color.Cyan("\n🔗 서버 정보:")
	color.Cyan("- IP: %s", loginResp.Data.IP)
	color.Cyan("- 서버 이름: %s", loginResp.Data.ServerName)
	color.Cyan("- 포트: %d", loginResp.Data.Port)

	color.Green("\n🚀 서버에 접속합니다...")

	// SSH 접속
	cmd := exec.Command("bash", "-c", loginResp.Data.SSHCommand)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		color.Red("서버 접속 실패:", err)
		os.Exit(1)
	}
}

func status() {
	color.Blue("=== BSH 구독 상태 조회 ===")

	companyWallet := getInput("회사 지갑 주소: ", false)

	// API 요청 데이터 구성
	jsonData := fmt.Sprintf(`{
		"company_wallet": "%s"
	}`, companyWallet)

	// API 요청 전송
	resp, err := http.Post(getServerURL("/subscription-status"), "application/json", strings.NewReader(jsonData))
	if err != nil {
		color.Red("서버 연결 오류:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// 응답 처리
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		color.Red("응답 읽기 오류:", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		color.Red("구독 상태 조회 실패: HTTP %d", resp.StatusCode)
		color.Red("응답 내용: %s", string(body))
		os.Exit(1)
	}

	var statusResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    struct {
			CompanyWallet    string `json:"company_wallet"`
			CompanyName      string `json:"company_name"`
			CeoName          string `json:"ceo_name"`
			Email            string `json:"email"`
			SubscriptionType uint8  `json:"subscription_type"`
			SubscriptionEnd  int64  `json:"subscription_end"`
			IsActive         bool   `json:"is_active"`
			IsExpired        bool   `json:"is_expired"`
			RemainingDays    int    `json:"remaining_days"`
			CheckedAt        string `json:"checked_at"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &statusResp); err != nil {
		color.Red("응답 파싱 오류:", err)
		os.Exit(1)
	}

	if !statusResp.Success {
		color.Red("구독 상태 조회 실패: %s", statusResp.Message)
		os.Exit(1)
	}

	color.Green("구독 상태 조회 성공!")
	color.Cyan("회사 정보:")
	color.Cyan("- 회사 지갑: %s", statusResp.Data.CompanyWallet)
	color.Cyan("- 회사 이름: %s", statusResp.Data.CompanyName)
	color.Cyan("- 대표자: %s", statusResp.Data.CeoName)
	color.Cyan("- 이메일: %s", statusResp.Data.Email)

	color.Cyan("\n구독 정보:")
	subscriptionTypeText := ""
	switch statusResp.Data.SubscriptionType {
	case 1:
		subscriptionTypeText = "1개월"
	case 2:
		subscriptionTypeText = "3개월"
	case 3:
		subscriptionTypeText = "1년"
	default:
		subscriptionTypeText = "알 수 없음"
	}
	color.Cyan("- 구독 유형: %s", subscriptionTypeText)
	color.Cyan("- 구독 만료일: %d", statusResp.Data.SubscriptionEnd)
	color.Cyan("- 활성 상태: %t", statusResp.Data.IsActive)
	color.Cyan("- 만료 여부: %t", statusResp.Data.IsExpired)
	color.Cyan("- 남은 일수: %d일", statusResp.Data.RemainingDays)
	color.Cyan("- 조회 일시: %s", statusResp.Data.CheckedAt)

	if statusResp.Data.IsExpired {
		color.Red("\n⚠️  구독이 만료되었습니다! 구독을 연장해주세요.")
	} else if statusResp.Data.RemainingDays <= 7 {
		color.Yellow("\n⚠️  구독이 곧 만료됩니다! (%d일 남음)", statusResp.Data.RemainingDays)
	} else {
		color.Green("\n✅ 구독이 정상적으로 활성화되어 있습니다.")
	}
}

func charge() {
	color.Blue("=== BSH 구독 연장 (메타마스크 결제) ===")

	companyWallet := getInput("회사 지갑 주소: ", false)

	// 구독 상태 확인
	color.Cyan("📊 현재 구독 상태 확인 중...")

	statusJsonData := fmt.Sprintf(`{"company_wallet": "%s"}`, companyWallet)
	statusResp, err := http.Post(getServerURL("/subscription-status"), "application/json", strings.NewReader(statusJsonData))
	if err != nil {
		color.Red("서버 연결 오류:", err)
		os.Exit(1)
	}
	defer statusResp.Body.Close()

	statusBody, err := io.ReadAll(statusResp.Body)
	if err != nil {
		color.Red("응답 읽기 오류:", err)
		os.Exit(1)
	}

	var statusData struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    struct {
			CompanyName      string `json:"company_name"`
			SubscriptionType uint8  `json:"subscription_type"`
			RemainingDays    int    `json:"remaining_days"`
			IsExpired        bool   `json:"is_expired"`
		} `json:"data"`
	}

	if err := json.Unmarshal(statusBody, &statusData); err != nil {
		color.Red("응답 파싱 오류:", err)
		os.Exit(1)
	}

	if !statusData.Success {
		color.Red("구독 상태 확인 실패: %s", statusData.Message)
		os.Exit(1)
	}

	color.Green("✅ 구독 상태 확인 완료")
	color.Cyan("- 회사: %s", statusData.Data.CompanyName)
	color.Cyan("- 남은 일수: %d일", statusData.Data.RemainingDays)

	// 구독 유형 선택
	fmt.Println("\n💳 구독 유형을 선택하세요:")
	fmt.Println("1. 1개월 (0.000001 ETH)")
	fmt.Println("2. 3개월 (0.0000025 ETH)")
	fmt.Println("3. 1년 (0.000008 ETH)")
	subscriptionType := getInput("구독 유형 (1-3): ", false)

	// 결제 정보 표시
	priceText := ""
	switch subscriptionType {
	case "1":
		priceText = "0.000001 ETH"
	case "2":
		priceText = "0.0000025 ETH"
	case "3":
		priceText = "0.000008 ETH"
	default:
		color.Red("잘못된 구독 유형입니다.")
		os.Exit(1)
	}

	color.Yellow("\n💰 결제 정보:")
	color.Yellow("- 선택한 구독: %s", subscriptionType)
	color.Yellow("- 결제 금액: %s", priceText)
	color.Yellow("- 회사 지갑: %s", companyWallet)

	// 메타마스크 결제 페이지 URL 생성
	paymentURL := fmt.Sprintf("http://localhost:1111/payment?wallet=%s&type=%s", companyWallet, subscriptionType)

	color.Cyan("\n🌐 메타마스크 결제 페이지를 브라우저에서 열어주세요:")
	color.Cyan("URL: %s", paymentURL)

	// 브라우저에서 결제 페이지 열기
	color.Cyan("\n🚀 브라우저에서 결제 페이지를 여는 중...")

	var openCmd *exec.Cmd
	switch {
	case isCommandAvailable("open"):
		openCmd = exec.Command("open", paymentURL)
	case isCommandAvailable("xdg-open"):
		openCmd = exec.Command("xdg-open", paymentURL)
	case isCommandAvailable("start"):
		openCmd = exec.Command("start", paymentURL)
	default:
		color.Yellow("⚠️  브라우저를 자동으로 열 수 없습니다. 위 URL을 복사하여 브라우저에 붙여넣어주세요.")
	}

	if openCmd != nil {
		if err := openCmd.Start(); err != nil {
			color.Yellow("⚠️  브라우저 열기 실패: %v", err)
			color.Yellow("위 URL을 복사하여 브라우저에 붙여넣어주세요.")
		}
	}

	color.Green("\n✅ 결제 페이지가 열렸습니다!")
	color.Cyan("💡 결제 완료 후 'bsh status' 명령어로 구독 상태를 확인하세요.")
}
