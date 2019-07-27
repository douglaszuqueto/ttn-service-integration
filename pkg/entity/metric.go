package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// Metric data
type Metric struct {
	ID        string    `json:"id"`
	AppID     string    `json:"app_id"`
	DevID     string    `json:"dev_id"`
	Payload   Payload   `json:"payload"`
	Time      time.Time `json:"time"`
	CreatedAt time.Time `json:"created_at"`
}

// Payload data
type Payload map[string]interface{}

// Value value
func (a Payload) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan scan
func (a *Payload) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
