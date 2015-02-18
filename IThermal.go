package raspberryFly

import "time"

type Thermal interface {
	Records() []IgcRecord
	AverageStrength() float64
	HeightGain() int
	StartHeight() int
	EndHeight() int
	Date() time.Time
	Duration() time.Duration
	WindDir() int64
	WindStrength() int64
	AproxStartPoint() Position
}
