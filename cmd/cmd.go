package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/douglaszuqueto/ttn-service-integration/config"
	"github.com/douglaszuqueto/ttn-service-integration/pkg/esp32"
	"github.com/douglaszuqueto/ttn-service-integration/pkg/mqtt"
	"github.com/douglaszuqueto/ttn-service-integration/pkg/service"
	"github.com/douglaszuqueto/ttn-service-integration/pkg/storage"

	paho "github.com/eclipse/paho.mqtt.golang"
)

var configPath = flag.String("config", "", "Path do arquivo de configuração")
var clients = make(map[int]*mqtt.Client, 5)

// Execute execute
func Execute() {
	flag.Parse()

	if len(*configPath) == 0 {
		log.Fatalln("Arquivo de configuração inválido")
	}

	config.LoadConfig(*configPath)

	// Configurações
	fmt.Printf("=> Geral\n\n")
	fmt.Printf("* Provider: \t%v\n", config.Cfg.Provider)
	fmt.Printf("* Timezone: \t%v\n\n", config.Cfg.Timezone)

	// Leitura banco de dados
	fmt.Printf("=> PostgreSQL\n\n")
	fmt.Printf("* Host: \t%v\n* Database: \t%v\n\n", config.Cfg.Postgres.Host, config.Cfg.Postgres.Database)

	// Leitura dados TTN
	fmt.Printf("=> TTN\n\n")
	for i, item := range config.Cfg.Ttn {
		fmt.Printf("* %v Host: \t%v\n* %v App ID: \t%v\n\n", i, item.Host, i, item.User)
	}
	fmt.Println()

	load()

	for {
		select {
		case <-time.After(15 * time.Second):
			// log.Println("[APP] tick...")
		}
	}
}

func load() {
	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("[APP] Encerrando aplicação...")
		storage.CloseConnection()

		for _, c := range clients {
			c.CloseConnection()
		}

		time.Sleep(1 * time.Second)

		log.Println("[APP] Aplicação encerrada!")
		os.Exit(0)
	}()

	storage.Connect()
	go func() {
		c := esp32.RunService()
		clients[len(clients)+1] = c
	}()

	if config.Cfg.Provider == "mqtt" {
		initMQTTProvider()
		return
	}

	// init http provider
	initHTTPProvider()
}

func initMQTTProvider() {
	onMessage := func(client paho.Client) {
		token := client.Subscribe("+/devices/+/up", 0, func(client paho.Client, msg paho.Message) {
			service.OnMessage(msg.Topic(), msg.Payload())
		})
		token.Wait()

		if token.Error() != nil {
			log.Println(token.Error())
		}
	}

	for i, item := range config.Cfg.Ttn {
		client := mqtt.NewClient(
			strconv.Itoa(i),
			item.Host,
			item.User,
			item.Password,
			onMessage,
		)
		go client.Run()

		clients[i] = client
	}
}

func initHTTPProvider() {

}
