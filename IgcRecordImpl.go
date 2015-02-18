package raspberryFly

import (
	"strconv"
	"time"
)

type IgcRecordImpl struct {
	Record string
}

func (i IgcRecordImpl) RecordString() string {
	return i.Record
}

func (i IgcRecordImpl) RecordType() string {
	return i.Record[0:1]
}

func (i IgcRecordImpl) Time() time.Time {
	hour, _ := strconv.ParseInt(i.Record[1:3], 10, 64)
	minute, _ := strconv.ParseInt(i.Record[3:5], 10, 64)
	second, _ := strconv.ParseInt(i.Record[5:7], 10, 64)

	return time.Date(2015, time.February, 12, int(hour), int(minute), int(second), 0, time.UTC)
}

func (i IgcRecordImpl) Position() Position {
	lat := CoordinateImpl{i.Record[7:15]}
	lon := CoordinateImpl{i.Record[15:24]}

	return PositionImpl{lat, lon}
}

func (i IgcRecordImpl) IsValid() bool {
	return i.Record[24:25] == "A"
}

func (i IgcRecordImpl) PreassureAlt() int {
	offset := i.calcOffset()
	alt, _ := strconv.ParseInt(i.Record[24+offset:29+offset], 10, 64)

	return int(alt)
}

func (i IgcRecordImpl) GpsAlt() int {
	offset := i.calcOffset()
	alt, _ := strconv.ParseInt(i.Record[29+offset:34+offset], 10, 64)

	return int(alt)
}

func (i IgcRecordImpl) calcOffset() int {
	offset := 0
	if i.IsValid() {
		offset = 1
	}
	return offset
}
