package log

import (
	"github.com/google/uuid"
	"time"
)

// swagger:model
type Log struct {
	ID uuid.UUID `json:"id"`
	L  string    `json:"l"`
	M  string    `json:"m"`
	T  time.Time `json:"t"`
}
