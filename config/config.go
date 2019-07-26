package config

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

// Config config
type Config struct {
	Provider string `json:"provider"`
	Timezone string `json:"timezone"`
	API      struct {
		Port   int `json:"port"`
		Secure struct {
			Active bool   `json:"active"`
			Token  string `json:"token"`
		} `json:"secure"`
	} `json:"api"`
	Ttn []struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		Token    string `json:"token"`
	} `json:"ttn"`
	Postgres struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"postgres"`
	Stream struct {
		Active bool `json:"active"`
		Broker struct {
			Host     string `json:"host"`
			Port     string `json:"port"`
			User     string `json:"user"`
			Password string `json:"password"`
		} `json:"broker"`
	} `json:"stream"`
}

// Cfg Cfg
var Cfg Config

// LoadConfig LoadConfig
func LoadConfig(configPath string) {
	// read a config file
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalln("Arquivo de configuração não encontrado:", err)
	}

	err = json.NewDecoder(file).Decode(&Cfg)
	if err != nil {
		log.Fatalln("Arquivo de configuração não pode ser carregado:", err)
	}

	// provider
	if Cfg.Provider != "mqtt" && Cfg.Provider != "http" {
		log.Fatalln("O serviço de integração possui suporte apenas para o protocolo MQTT ou HTTP!", Cfg.Provider)
	}

	// test timezone
	_, err = time.LoadLocation(Cfg.Timezone)
	if err != nil {
		log.Fatalln("Timezone inválida:", err)
	}
}
