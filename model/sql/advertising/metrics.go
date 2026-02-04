package advertising

import (
	"database/sql/driver"
	"github.com/cngamesdk/go-core/model/sql"
)

type Metrics struct {
	Impressions int64   `json:"impressions"`
	Clicks      int64   `json:"clicks"`
	Conversions int64   `json:"conversions"`
	Spend       float64 `json:"spend"`
	CTR         float64 `json:"ctr"`
	CPC         float64 `json:"cpc"`
	CPA         float64 `json:"cpa"`
	ROI         float64 `json:"roi"`
}

// Scan Scanner
func (args *Metrics) Scan(value interface{}) error {
	return sql.JsonScan(args, value)
}

// Value Valuer
func (args Metrics) Value() (driver.Value, error) {
	return sql.JsonValue(args)
}
