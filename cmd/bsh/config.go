package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Server struct {
		Host     string `json:"host"`
		BasePath string `json:"base_path"`
	} `json:"server"`
}

func loadConfig() (*Config, error) {
	// 홈 디렉토리 찾기
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	// 설정 파일 경로
	configPath := filepath.Join(homeDir, ".bsh", "config.json")

	// 설정 파일 읽기
	data, err := os.ReadFile(configPath)
	if err != nil {
		// 설정 파일이 없으면 기본값 사용
		config := &Config{}
		config.Server.Host = "bsh.kimtuna.com"
		config.Server.BasePath = "/api"
		return config, nil
	}

	// JSON 파싱
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
