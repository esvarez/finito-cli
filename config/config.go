package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	home, _    = os.UserHomeDir()
	finitoDir  = home + "/.config/finito/"
	configFile = finitoDir + "config.json"
)

type Configuration struct {
	App App `json:"app"`
}

type App struct {
	SheetID *string `mapstructure:"sheet_id" yaml:"sheet_id" json:"sheet_id"`
}

func LoadConfiguration() (*Configuration, error) {
	file := viper.New()
	file.SetConfigFile(configFile)

	if err := file.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Configuration{}
	if err := file.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func SaveConfiguration(cfg *Configuration) error {
	file := viper.New()
	file.SetConfigFile(configFile)
	file.SetConfigType("json")
	file.Set("app", cfg.App)

	if err := file.WriteConfig(); err != nil {
		return err
	}

	return nil
}

func init() {
	if _, err := os.Stat(finitoDir); os.IsNotExist(err) {
		err = os.Mkdir(finitoDir, os.ModePerm)
		if err != nil {
			log.Printf("error creating config directory: %v", err)
		}
	}
	file := viper.New()
	file.SetConfigFile(configFile)
	file.AddConfigPath(finitoDir)
	file.SetConfigType("json")
	file.Set("app.sheet_id", nil)
	file.SafeWriteConfig()
}
