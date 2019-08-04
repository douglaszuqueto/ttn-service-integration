package esp32

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"time"

	cayenne "github.com/douglaszuqueto/ttn-service-integration/pkg/cayenne-lpp"
	"github.com/douglaszuqueto/ttn-service-integration/pkg/mqtt"

	paho "github.com/eclipse/paho.mqtt.golang"
)

// RunService RunService
func RunService() *mqtt.Client {
	onMessage := func(client paho.Client) {
		token := client.Subscribe("device/+/metric", 0, func(client paho.Client, msg paho.Message) {
			onMessage(msg.Topic(), msg.Payload())
		})
		token.Wait()

		if token.Error() != nil {
			log.Println(token.Error())
		}
	}

	client := mqtt.NewClient("esp32", "192.168.0.150", "", "", onMessage)
	client.Run()

	return client
}

func onMessage(topic string, p []byte) {
	decoded, _ := base64.StdEncoding.DecodeString(string(p))

	lpp := cayenne.CayenneLPP{}
	lpp.DecodeBytes(decoded)

	parts := strings.Split(topic, "/")
	deviceID := parts[1]

	fmt.Printf("=> Cayenne LPP utilizando ESP32 via WiFi e MQTT\n\n")
	fmt.Printf("* Device: \tESP32\n")
	fmt.Printf("* DeviceID: \t%v\n", deviceID)

	fmt.Printf("* Métricas:\n\n")

	for k, v := range lpp.TemperatureSensor {
		fmt.Printf("\t => Temperature %v => %v\n", k, v)
	}

	for k, v := range lpp.HumiditySensor {
		fmt.Printf("\t => Humidity %v => %v\n", k, v)
	}

	for k, v := range lpp.UnixTime {
		fmt.Printf("\t => UnixTime %v => %v\n", k, v)
	}

	fmt.Printf("\nRecebido em: \t%v\n", time.Unix(int64(lpp.UnixTime[1]), 0).Format("02/01/2006 15:04:05"))
	fmt.Printf("Processado em: \t%v\n", time.Now().Local().Format("02/01/2006 15:04:05"))
	fmt.Println()
}
