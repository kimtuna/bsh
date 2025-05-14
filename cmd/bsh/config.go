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

func getDefaultConfig() (Config, error) {
	return Config{
		Server: struct {
			Host     string `json:"host"`
			BasePath string `json:"base_path"`
		}{
			Host:     "kimtuna.kr",
			BasePath: "/bsh/api",
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
