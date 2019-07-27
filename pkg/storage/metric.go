package storage

import (
	"log"

	"github.com/douglaszuqueto/ttn-service-integration/pkg/entity"
)

// GetMetrics GetMetrics
func GetMetrics() ([]entity.Metric, error) {
	l := []entity.Metric{}

	// locationTimezone, err := time.LoadLocation(config.Cfg.Timezone)
	// if err != nil {
	// 	log.Println(err)
	// }
	// localTime := e.Time.In(locationTimezone)

	// fmt.Println(locationTimezone.String(), localTime.Format("02/01/2006 15:04:05"))

	return l, nil
}

// StoreMetric store a user
func StoreMetric(e entity.Metric) error {
	query := `
		INSERT INTO metric
			(app_id, dev_id, payload, time) 
		VALUES 
			($1, $2, $3, $4)
		RETURNING id`

	id, err := doInsert(query, e.AppID, e.DevID, e.Payload, e.Time)
	if err != nil {
		return HandlePSQLError(err)
	}

	log.Println("[Metric] Nova métrica inserida:", id)
	return nil
}
