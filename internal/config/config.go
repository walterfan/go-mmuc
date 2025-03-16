// internal/config/config.go
package config

import (
	"os"
	"strconv"
	"fmt"

	"gopkg.in/yaml.v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ServerPort uint32 `yaml:"server_port"`
	FailFast   bool   `yaml:"fail_fast"`

	LogDir   string `yaml:"log_dir"`
	LogLevel string `yaml:"log_level"`

	DbHost     string `yaml:"db_host"`
	DbPort     int    `yaml:"db_port"`
	DbUsername string `yaml:"db_username"`
	DbPassword string `yaml:"db_password"`
	DbName     string `yaml:"db_name"`
}

// String returns a string representation of the Config struct
func (c Config) String() string {
	return fmt.Sprintf(`Config{
		ServerPort: %d, 
		FailFast: %t, 
		LogDir: %s, 
		LogLevel: %s, 
		DbHost: %s,
		DbPort: %d, 
		DbUsername: %s, 
		DbPassword: %s, 
		DbName: %s}`,
		c.ServerPort, c.FailFast, c.LogDir, c.LogLevel,
		c.DbHost, c.DbPort, c.DbUsername, c.DbPassword, c.DbName)
}

func LoadConfig(configPath string) (Config, error) {
	var config Config

	// Load .env file
	if err := godotenv.Load(); err != nil {
		logrus.Warnf("Failed to load .env file: %v", err)
	}


	file, err := os.Open(configPath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return config, err
	}

	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		config.DbHost = dbHost
	}

	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		if port, err := strconv.Atoi(dbPort); err == nil {
			config.DbPort = port
		} else {
			logrus.Warnf("Invalid DB_PORT value in .env: %v", err)
		}
	}

	if dbUsername := os.Getenv("DB_USER"); dbUsername != "" {
		config.DbUsername = dbUsername
	}

	if dbPassword := os.Getenv("DB_PWD"); dbPassword != "" {
		config.DbPassword = dbPassword
	}

	if dbName := os.Getenv("DB_NAME"); dbName != "" {
		config.DbName = dbName
	}

	return config, nil
}
