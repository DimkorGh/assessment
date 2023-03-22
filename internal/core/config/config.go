package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	viper  *viper.Viper
	Server Server
	Logger Logger
}

type Server struct {
	Address string
	Port    int
}

type Logger struct {
	Level string
}

func NewConfig(viper *viper.Viper) *Config {
	return &Config{
		viper: viper,
	}
}

func (cfg *Config) Load(configName string) *Config {
	cfg.viper.SetConfigName(configName)
	cfg.viper.AddConfigPath(cfg.getConfigFilePath())
	cfg.viper.SetConfigType("yaml")

	if err := cfg.viper.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading config file: %s", err.Error())
	}

	if err := cfg.viper.Unmarshal(cfg); err != nil {
		log.Fatalf("Error while unmarshalling config file into a struct: %s", err.Error())
	}

	return cfg
}

func (cfg *Config) getConfigFilePath() string {
	path, _ := os.Getwd()

	for {
		if _, err := os.Stat(path + "/config/config.yaml"); err != nil {
			lastSlash := strings.LastIndex(path, "/")
			if lastSlash == -1 {
				log.Fatal("Error no config file found")
			}
			path = path[:lastSlash]

			continue
		}

		break
	}

	return path + "/config/"
}

func (cfg *Config) SetServerAddress(address string) {
	cfg.Server.Address = address
}

func (cfg *Config) SetServerPort(port int) {
	cfg.Server.Port = port
}
