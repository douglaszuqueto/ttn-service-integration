package cmd

import (
	"flag"
	"fmt"
	"log"

	"github.com/douglaszuqueto/ttn-service-integration/config"
)

var configPath = flag.String("config", "", "Path do arquivo de configuração")

// Execute execute
func Execute() {
	flag.Parse()

	if len(*configPath) == 0 {
		log.Fatalln("Arquivo de configuração inválido")
	}

	config.LoadConfig(*configPath)

	// Leitura banco de dados
	fmt.Printf("[Cfg Pg] Host %v => database %v\n", config.Cfg.Postgres.Host, config.Cfg.Postgres.Database)

	// Leitura dados TTN
	for _, item := range config.Cfg.Ttn {
		fmt.Printf("[Cfg TTN] Host %v => APP ID %v\n", item.Host, item.User)
	}
}
