package model

import (
	"time"
)

type IgcRecord interface {
	RecordString() string
	RecordType() string
	Time() time.Time
	Position() Position
	IsValid() bool
	PreassureAltitude() int
	GpsAltitude() int
}
