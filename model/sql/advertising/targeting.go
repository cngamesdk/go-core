package advertising

import (
	"database/sql/driver"
	"github.com/cngamesdk/go-core/model/sql"
)

type Targeting struct {
	Location    []string `json:"location"`
	AgeRange    [2]int   `json:"age_range"`
	Gender      string   `json:"gender"`
	Interests   []string `json:"interests"`
	Platforms   []string `json:"platforms"`
	DeviceTypes []string `json:"device_types"`
	NetworkType string   `json:"network_type"`
}

// Scan Scanner
func (args *Targeting) Scan(value interface{}) error {
	return sql.JsonScan(args, value)
}

// Value Valuer
func (args Targeting) Value() (driver.Value, error) {
	return sql.JsonValue(args)
}
