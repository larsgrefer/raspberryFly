package model

type Coordinate interface {
	CoordinateString() string
	CoordinateFloat() float64
	CoordinateFloatRad() float64
	Degrees() int
	Minutes() int
	Seconds() int
	SecondsFloat() float64
	Hemisphere() string
}

type Position interface {
	Lat() Coordinate
	Lon() Coordinate
	DistanceTo(p Position) float64
	CourseTo(p Position) float64
}
