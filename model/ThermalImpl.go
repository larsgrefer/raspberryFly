package model

import (
	"math"
	"time"
)

type ThermalImpl struct {
	ThermalRecords []IgcRecord
}

func (t ThermalImpl) Records() []IgcRecord {
	return t.ThermalRecords
}

func (t ThermalImpl) AverageStrength() float64 {
	first := t.ThermalRecords[0]
	last := t.ThermalRecords[len(t.ThermalRecords)-1]

	time := last.Time().Sub(first.Time()).Seconds()

	heightGainGps := last.GpsAltitude() - first.GpsAltitude()
	heightGainPreassure := last.PreassureAltitude() - first.PreassureAltitude()

	averageHeightGain := float64(heightGainGps+heightGainPreassure) / 2.0

	return averageHeightGain / time
}

func (t ThermalImpl) HeightGain() int {
	return t.EndHeight() - t.StartHeight()
}

func (t ThermalImpl) StartHeight() int {
	return t.ThermalRecords[0].GpsAltitude()
}

func (t ThermalImpl) EndHeight() int {
	return t.ThermalRecords[len(t.ThermalRecords)-1].GpsAltitude()
}

func (t ThermalImpl) Time() time.Time {
	return t.ThermalRecords[0].Time()
}

func (t ThermalImpl) Duration() time.Duration {
	return t.ThermalRecords[len(t.ThermalRecords)-1].Time().Sub(t.ThermalRecords[0].Time())
}

func (t ThermalImpl) WindDir() int64 {
	startCourse := t.ThermalRecords[0].Position().CourseTo(t.ThermalRecords[1].Position())

	courseSum := 0.0
	lastCourse := -1.0

	var fittingPos Position
	fittingCourse := 0.0

	for i := len(t.ThermalRecords) - 1; i > 1 && courseSum < 360; i-- {
		course := t.ThermalRecords[i-1].Position().CourseTo(t.ThermalRecords[i].Position())
		if lastCourse == -1 {
			fittingPos = t.ThermalRecords[i-1].Position()
			fittingCourse = course
		} else {
			courseSum += math.Abs(course - lastCourse)
			if math.Abs(course-startCourse) < math.Abs(fittingCourse-startCourse) {
				fittingPos = t.ThermalRecords[i-1].Position()
				fittingCourse = course
			}
		}
		lastCourse = course
	}

	return int64(t.ThermalRecords[0].Position().CourseTo(fittingPos))
}

func (t ThermalImpl) WindSpeed() int64 {
	return int64((t.ThermalRecords[0].Position().DistanceTo(t.ThermalRecords[len(t.ThermalRecords)-1].Position())) / t.Duration().Hours())
}

func (t ThermalImpl) Date() time.Time {
	return t.ThermalRecords[0].Time()
}
