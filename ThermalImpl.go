package raspberryFly

import (
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
