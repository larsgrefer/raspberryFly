package model

import (
	"fmt"
	"math"
	"strconv"
)

type PositionImpl struct {
	Latitude  Coordinate
	Longitude Coordinate
}

func (pos PositionImpl) DistanceTo(p Position) float64 {
	lat1 := pos.Lat().CoordinateFloatRad()
	lon1 := pos.Lon().CoordinateFloatRad()

	lat2 := p.Lat().CoordinateFloatRad()
	lon2 := p.Lon().CoordinateFloatRad()

	e := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(lon2-lon1))

	dist := e * 6378.137

	return dist
}

func (pos PositionImpl) CourseTo(p Position) float64 {
	lat1 := pos.Lat().CoordinateFloatRad()
	lon1 := pos.Lon().CoordinateFloatRad()

	lat2 := p.Lat().CoordinateFloatRad()
	lon2 := p.Lon().CoordinateFloatRad()

	e := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(lon2-lon1))

	f := (math.Cos(lat2) - math.Cos(lat1)*math.Cos(e)) / (math.Sin(lat1) * math.Sin(e))

	if math.Sin(p.Lon().CoordinateFloat()-pos.Lon().CoordinateFloat()) < 0.0 {
		f = math.Acos(f)
	} else {
		f = 2.0*math.Pi - math.Acos(f)
	}

	course := f / math.Pi * 180.0

	return course
}

func (p PositionImpl) Lat() Coordinate {
	return p.Latitude
}

func (p PositionImpl) Lon() Coordinate {
	return p.Longitude
}

func (p PositionImpl) String() string {
	return fmt.Sprintf("Position:\nLat: %v\nLon: %v", p.Lat(), p.Lon())
}

type CoordinateImpl struct {
	Coord string
}

func (c CoordinateImpl) CoordinateString() string {
	return c.Coord
}

func (c CoordinateImpl) CoordinateFloat() float64 {
	return float64(c.Degrees()) + float64(c.Minutes())/60.0 + c.SecondsFloat()/3600.0
}

func (c CoordinateImpl) CoordinateFloatRad() float64 {
	return c.CoordinateFloat() / 180.0 * math.Pi
}

func (c CoordinateImpl) Degrees() int {
	offset := c.calcOffset()

	deg, _ := strconv.ParseInt(c.Coord[0:2+offset], 10, 64)

	return int(deg)
}

func (c CoordinateImpl) Minutes() int {
	offset := c.calcOffset()

	deg, _ := strconv.ParseInt(c.Coord[2+offset:4+offset], 10, 64)

	return int(deg)
}

func (c CoordinateImpl) Seconds() int {
	return int(c.SecondsFloat())
}

func (c CoordinateImpl) SecondsFloat() float64 {
	offset := c.calcOffset()

	deg, _ := strconv.ParseFloat(c.Coord[4+offset:7+offset], 64)

	return deg / 100 * 6
}

func (c CoordinateImpl) Hemisphere() string {
	offset := c.calcOffset()

	hem := c.Coord[7+offset : 8+offset]

	return hem
}

func (c CoordinateImpl) calcOffset() int {
	offset := 0

	if len(c.Coord) == 9 {
		offset = 1
	}

	return offset
}

func (c CoordinateImpl) String() string {
	return fmt.Sprintf("%3v° %2v' %2v'' %v", c.Degrees(), c.Minutes(), c.Seconds(), c.Hemisphere())
}
