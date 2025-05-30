package config

import (
	"fmt"
	"strings"

	"github.com/SPVJ/fs-common-lib/core/client"
	"github.com/spf13/viper"
)

type Config struct {
	Server    Server
	Webclient WebClient
}

type Server struct {
	IsProductionMode bool
	Port             string
}

type WebClient struct {
	DocumentService        DocumentService
	DocumentRequestService DocumentRequestService
}

type DocumentService struct {
	client.HttpClientConfig
	BaseURL string
	URL     URL
}

type DocumentRequestService struct {
	client.HttpClientConfig
	BaseURL string
	URL     DocumentRequestURL
}

type URL struct {
	Base                    string
	ValidateDocument        string
	DeleteDocument          string
	UpdateDocument          string
	ReTranslateDocument     string
	GetDocumentStatus       string
	ValidateDocumentRequest string
	DeleteDocumentRequest   string
	ListDocumentRequests    string
	CreateDocumentRequest   string
	GetCompanies            string
}

type DocumentRequestURL struct {
	Base                    string
	GetDocumentRequest      string
	CreateDocumentRequest   string
	UpdateDocumentRequest   string
	DeleteDocumentRequest   string
	ListDocumentRequests    string
	ValidateDocumentRequest string
}

var AppConfig Config

func LoadConfig(configPath string) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	// Add the custom config path provided as an argument
	if configPath != "" {
		viper.AddConfigPath(configPath)
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		if _, isFileNotFound := err.(viper.ConfigFileNotFoundError); isFileNotFound {
			return fmt.Errorf("config file not found")
		} else {
			return fmt.Errorf("CONFIG:fatal error config file: %s ", err)
		}
	}

	viper.AutomaticEnv()
	err := viper.Unmarshal(&AppConfig)

	return err
}
