package advertising

import (
	"database/sql/driver"
	"github.com/cngamesdk/go-core/model/sql"
)

type Creative struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Images       []string `json:"images"`
	Videos       []string `json:"videos"`
	LandingPage  string   `json:"landing_page"`
	CallToAction string   `json:"call_to_action"`
}

// Scan Scanner
func (args *Creative) Scan(value interface{}) error {
	return sql.JsonScan(args, value)
}

// Value Valuer
func (args Creative) Value() (driver.Value, error) {
	return sql.JsonValue(args)
}
