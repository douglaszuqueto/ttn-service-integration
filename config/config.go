package config

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

// Config config
type Config struct {
	Provider string
	Timezone string
	API      struct {
		Port   int
		Secure struct {
			Active bool
			Token  string
		}
	}
	Ttn []struct {
		Host     string
		User     string
		Password string
		Token    string
	}
	Postgres struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
	}
	Stream struct {
		Active bool
		Broker struct {
			Host     string
			Port     string
			User     string
			Password string
		}
	}
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
