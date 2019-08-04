package mqtt

import (
	"fmt"
	"log"
	"net/url"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Client Client
type Client struct {
	id       string
	host     string
	user     string
	password string

	client      mqtt.Client
	onMessage   func(client mqtt.Client)
	isConnected bool
}

// NewClient NewClient
func NewClient(id string, host, user, password string, onMessage func(client mqtt.Client)) *Client {
	client := &Client{
		id:       id,
		host:     host,
		user:     user,
		password: password,

		onMessage: onMessage,
	}

	return client
}

func (c *Client) connectionLostHandler(client mqtt.Client, err error) {
	log.Println("[MQTT] Conexão perdida:", c.id)
	c.isConnected = false
	c.connectLoop()
}

func (c *Client) onConnectHandler(client mqtt.Client) {
	log.Println("[MQTT] Conectado:", c.id)
	c.isConnected = true

	c.onMessage(client)
}

func (c *Client) connect(clientID string, uri *url.URL) mqtt.Client {
	opts := mqtt.NewClientOptions()

	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))

	opts.SetClientID(clientID)
	opts.SetCleanSession(true)
	opts.SetAutoReconnect(false)

	opts.SetKeepAlive(5 * time.Second)
	opts.SetPingTimeout(5 * time.Second)
	opts.SetConnectTimeout(30 * time.Second)
	opts.SetWriteTimeout(30 * time.Second)
	opts.SetMaxReconnectInterval(30 * time.Second)

	opts.SetConnectionLostHandler(c.connectionLostHandler)
	opts.SetOnConnectHandler(c.onConnectHandler)

	opts.SetUsername(c.user)
	opts.SetPassword(c.password)

	c.client = mqtt.NewClient(opts)

	c.connectLoop()

	return c.client
}

// Run run
func (c *Client) Run() {
	host := fmt.Sprintf("tcp://%s:%s", c.host, "1883")
	uri, err := url.Parse(host)
	if err != nil {
		log.Fatal(err)
	}

	clientID := "ttn-service-integration" + time.Now().String()

	c.client = c.connect(clientID, uri)
}

func (c *Client) connectLoop() {
	for {
		if token := c.client.Connect(); token.Wait() && token.Error() != nil {
			log.Println("[MQTT] tentando reconectar-se:", c.id, token.Error())
		} else {
			break
		}

		time.Sleep(1000 * time.Millisecond)
	}
}

func (c *Client) reconnectLoop() {
	for {
		if c.isConnected {
			break
		}

		//c.CloseConnection()
		c.connectLoop()
	}
}

// CloseConnection closeConnection
func (c *Client) CloseConnection() {
	log.Println("[MQTT] Fechando conexão:", c.id)
	if c.isConnected {
		c.client.Disconnect(250)
	}
	log.Println("[MQTT] Conexão fechada!", c.id)
}
