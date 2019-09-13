package config

import (
	"path/filepath"

	"github.com/spf13/viper"
)

func NewConfig() *config {
	return &config{

	}
}

type config struct {

}

func (c *config) Init(path string) error {
	fullPath, err := filepath.Abs(path) // config/monitoring.yaml
	if err != nil {
		return err
	}
	viper.SetConfigFile(fullPath)
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}