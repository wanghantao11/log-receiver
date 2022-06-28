package log

import (
	"time"

	"github.com/google/uuid"
)

// Log DB schema
type Log struct {
	ID uuid.UUID
	L  string
	M  string
	T  time.Time
}
