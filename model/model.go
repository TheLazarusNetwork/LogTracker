package model

import (
	"time"

	core "github.com/textileio/go-threads/core/db"
)

// Log Struct to define a schema for the Logs Generated
type Log struct {
	ID        core.InstanceID `json:"_id"`
	Timestamp time.Time       `json:"timestamp"` // Creation Time
	Type      string          `json:"type"`      // Log Type - INFO, WARN, ERROR
	Data      string          `json:"data"`      // Log Data
	Hash      string          `json:"hash"`      // SHA256 Hash of Log’s Data
}
