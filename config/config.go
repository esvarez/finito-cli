package config

import (
	"os"

	"github.com/spf13/viper"
)

var (
	home, _    = os.UserHomeDir()
	finitoDir  = home + "/.finito/"
	configFile = finitoDir + "config.yaml"
)

type Configuration struct {
	App
}

type App struct {
	SheetID *string `mapstructure:"sheet_id" yaml:"sheet_id"`
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
	file.SetConfigType("yaml")
	file.Set("app", cfg.App)

	if err := file.WriteConfig(); err != nil {
		return err
	}

	return nil
}
func init() {
	file := viper.New()
	file.SetConfigFile(configFile)
	file.AddConfigPath(finitoDir)
	file.SetConfigType("yaml")
	file.Set("app.sheet_id", "")
	file.SafeWriteConfig()
}
