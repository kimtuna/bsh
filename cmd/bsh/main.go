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

func init() {
	rootCmd.AddCommand(registerCmd)
	rootCmd.AddCommand(loginCmd)

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
	return fmt.Sprintf("http://%s:%d%s", config.Server.Host, config.Server.Port, path)
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
	port := getInput("SSH 포트 (기본값: 22): ", false)
	if port == "" {
		port = "22"
	}

	// API 요청 데이터 구성
	jsonData := fmt.Sprintf(`{
		"company_wallet": "%s",
		"company_name": "%s",
		"ceo_name": "%s",
		"email": "%s",
		"subscription_type": %s,
		"ip": "%s",
		"server_name": "%s",
		"port": %s
	}`, companyWallet, companyName, ceoName, email, subscriptionType, ip, serverName, port)

	// API 요청 전송
	resp, err := http.Post(getServerURL("/api/register"), "application/json", strings.NewReader(jsonData))
	if err != nil {
		color.Red("서버 연결 오류:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		color.Red("등록 실패: HTTP %d", resp.StatusCode)
		os.Exit(1)
	}

	color.Green("회원가입이 완료되었습니다!")
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		IP         string `json:"ip"`
		ServerName string `json:"server_name"`
		Port       uint16 `json:"port"`
		SSHCommand string `json:"ssh_command"`
	} `json:"data"`
}

func login() {
	color.Blue("=== BSH 로그인 ===")

	companyWallet := getInput("회사 지갑 주소: ", false)

	// API 요청 데이터 구성
	jsonData := fmt.Sprintf(`{"company_wallet": "%s"}`, companyWallet)

	// API 요청 전송
	resp, err := http.Post(getServerURL("/api/login"), "application/json", strings.NewReader(jsonData))
	if err != nil {
		color.Red("서버 연결 오류:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// 응답 본문 읽기
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		color.Red("응답 읽기 오류:", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		color.Red("로그인 실패: %s", string(body))
		os.Exit(1)
	}

	// JSON 응답 파싱
	var loginResp LoginResponse
	if err := json.Unmarshal(body, &loginResp); err != nil {
		color.Red("응답 파싱 오류:", err)
		os.Exit(1)
	}

	if !loginResp.Success {
		color.Red("로그인 실패: %s", loginResp.Message)
		os.Exit(1)
	}

	color.Green("로그인 성공! 서버에 접속합니다...")
	color.Cyan("서버 정보:")
	color.Cyan("- IP: %s", loginResp.Data.IP)
	color.Cyan("- 서버 이름: %s", loginResp.Data.ServerName)
	color.Cyan("- 포트: %d", loginResp.Data.Port)

	// SSH 명령어 실행
	cmd := exec.Command("bash", "-c", loginResp.Data.SSHCommand)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		color.Red("SSH 접속 실패:", err)
		os.Exit(1)
	}
}
