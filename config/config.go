package config

// Config config
type Config struct {
	Provider string `json:"provider"`
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
