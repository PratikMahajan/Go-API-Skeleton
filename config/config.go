package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)


// Config holds application configuration
type Config struct {

	AccessToken string `required:"True"`

	AccessTokenSecret string `required:"True"`

	ApiKey string `required:"True"`

	ApiSecretKey string `required:"True"`

	// MetricsAddr is the metrics server's bind address
	MetricsAddr string `default:":9090" split_words:"true" required:"true"`

	// APIAddr is the API server's bind address
	APIAddr string `default:":5000" split_words:"true" required:"true"`
}

// NewConfig loads configuration values from environment variables
func NewConfig() (*Config, error){

	var config Config

	if err := envconfig.Process("APP", &config); err != nil {
		return nil, fmt.Errorf("error loading values from environment variables: %s",
		err.Error())
	}

	return &config, nil
}