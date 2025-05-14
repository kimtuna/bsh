package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type Config struct {
	Server struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"server"`
}

func getDefaultConfig() (Config, error) {
	portStr := os.Getenv("BSH_CLIENT_PORT")
	if portStr == "" {
		return Config{}, fmt.Errorf("BSH_CLIENT_PORT 환경 변수가 설정되지 않았습니다")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return Config{}, fmt.Errorf("BSH_CLIENT_PORT가 유효한 포트 번호가 아닙니다: %v", err)
	}

	return Config{
		Server: struct {
			Host string `json:"host"`
			Port int    `json:"port"`
		}{
			Host: "kimtuna.kr",
			Port: port,
		},
	}, nil
}

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".bsh", "config.json"), nil
}

func loadConfig() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	// 설정 파일이 없으면 기본 설정으로 생성
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
			return nil, err
		}

		config, err := getDefaultConfig()
		if err != nil {
			return nil, err
		}
		if err := saveConfig(&config); err != nil {
			return nil, err
		}
		return &config, nil
	}

	// 설정 파일 읽기
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func saveConfig(config *Config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}
