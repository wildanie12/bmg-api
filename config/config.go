package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// AppConfig main struct
type AppConfig struct {
	App struct {
		BaseURL string
		Port string
		URL string
	}
	DB struct {
		Name string
		Username string
		Password string
		Host string
	}
	DBTest struct {
		Name string
		Username string
		Password string
		Host string
	}
}

var appConfig *AppConfig

// New return new or old instance from initiated app config
func New() *AppConfig {
	
	if appConfig == nil {
		appConfig = initConfig()
	}
	return appConfig
}

func initConfig() *AppConfig {
	config := AppConfig{}

	err := godotenv.Load()
	if err != nil {
		config.App.BaseURL = "http://localhost"
		config.App.Port = "8000"
		config.DB.Name = "BMG_DB"
		config.DB.Username = "root"
		config.DB.Password = ""
		config.DB.Host = "localhost"
		config.DBTest.Name = os.Getenv("DB_NAME")
		config.DBTest.Username = os.Getenv("DB_USER")
		config.DBTest.Password = os.Getenv("DB_PASS")
		config.DBTest.Host = os.Getenv("DB_HOST")
	} else {
		config.App.BaseURL = os.Getenv("APP_BASE_URL")
		config.App.Port = os.Getenv("APP_PORT")
		config.DB.Name = os.Getenv("DB_NAME")
		config.DB.Username = os.Getenv("DB_USER")
		config.DB.Password = os.Getenv("DB_PASS")
		config.DB.Host = os.Getenv("DB_HOST")
		config.DBTest.Name = os.Getenv("DB_TEST_NAME")
		config.DBTest.Username = os.Getenv("DB_TEST_USER")
		config.DBTest.Password = os.Getenv("DB_TEST_PASS")
		config.DBTest.Host = os.Getenv("DB_TEST_HOST")
	}
	config.App.URL = strings.TrimRight(config.App.BaseURL, "/") + map[bool]string{ true: "", false: config.App.Port }[config.App.Port == "80"]

	return &config
}