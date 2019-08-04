package service

import (
	"encoding/json"
	"log"

	"github.com/douglaszuqueto/ttn-service-integration/pkg/entity"
	"github.com/douglaszuqueto/ttn-service-integration/pkg/storage"
)

// OnMessage OnMessage
func OnMessage(topic string, payloadRaw []byte) {
	var p entity.UplinkMessage

	err := json.Unmarshal(payloadRaw, &p)
	if err != nil {
		log.Println("[onMessage] Unmarshal:", err)
	}

	//

	if len(p.PayloadRaw) == 0 {
		log.Println("PayloadRaw is empty:", p.PayloadRaw)
		return
	}

	payloadFields, ok := p.PayloadFields.(map[string]interface{})
	if !ok {
		log.Println("PayloadFields casting error:", payloadFields)
		return
	}

	m := entity.Metric{
		AppID:   p.AppID,
		DevID:   p.DevID,
		Payload: payloadFields,
		Time:    p.Metadata.Time,
	}

	err = storage.StoreMetric(m)
	if err != nil {
		log.Println("StoreMetric:", err)
	}
}
