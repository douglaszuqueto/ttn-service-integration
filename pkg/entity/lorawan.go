package entity

import "time"

// UplinkMessage data
type UplinkMessage struct {
	AppID          string      `json:"app_id"`
	DevID          string      `json:"dev_id"`
	HardwareSerial string      `json:"hardware_serial"`
	Port           int         `json:"port"`
	Counter        int         `json:"counter"`
	IsRetry        bool        `json:"is_retry"`
	Confirmed      bool        `json:"confirmed"`
	PayloadRaw     string      `json:"payload_raw"`
	PayloadFields  interface{} `json:"payload_fields"`
	Metadata       `json:"metadata"`
	DownlinkURL    string `json:"downlink_url"`
}

// Metadata data
type Metadata struct {
	Time       time.Time `json:"time"`
	Frequency  float64   `json:"frequency"`
	Modulation string    `json:"modulation"`
	DataRate   string    `json:"data_rate"`
	BitRate    int       `json:"bit_rate"`
	CodingRate string    `json:"coding_rate"`
	Gateways   []struct {
		GtwID     string  `json:"gtw_id"`
		Timestamp int     `json:"timestamp"`
		Time      string  `json:"time,omitempty"`
		Channel   int     `json:"channel"`
		Rssi      int     `json:"rssi"`
		Snr       float64 `json:"snr"`
		RfChain   int     `json:"rf_chain"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Altitude  float64 `json:"altitude"`
	} `json:"gateways"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
}
