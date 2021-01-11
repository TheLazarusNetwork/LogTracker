package model

import (
	"time"

	core "github.com/textileio/go-threads/core/db"
)

// Log Struct to define a schema for the Logs Generated
type Log struct {
	ID        core.InstanceID `json:"_id"`
	Timestamp time.Time       `json:"timestamp"`
	Type      string          `json:"type"`
	Data      string          `json:"data"`
}
