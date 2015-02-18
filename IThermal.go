package raspberryFly

import "time"

type Thermal interface {
	Records() []IgcRecord
	AverageStrength() float64
	HeightGain() int64
	StartHeight() int64
	EndHeight() int64
	Date() time.Time
	Duration() time.Duration
	WindDir() int64
	WindStrength() int64
	AproxStartPoint() Position
}
