package raspberryFly

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"raspberryFly/model"
	"strconv"
	"strings"
	"time"
)

func FindCircles(records []model.IgcRecord) bool {
	if len(records) < 10 {
		return false
	}

	turnGes := 0.0

	for i := 0; i < len(records)-2; i++ {
		rec := records[len(records)-(i+1)]
		lastRec := records[len(records)-(i+2)]
		preLastRec := records[len(records)-(i+3)]
		time := rec.Time().Sub(lastRec.Time()).Seconds()

		course := int(lastRec.Position().CourseTo(rec.Position()))
		lastCourse := int(preLastRec.Position().CourseTo(lastRec.Position()))

		turn := float64(course-lastCourse) / time

		if turn > 100 {
			turn = float64(-(360-course)-lastCourse) / time
		} else if turn < -100 {
			turn = float64((360-lastCourse)+course) / time
		}
		turnGes += math.Abs(turn)
	}

	if turnGes > 9*float64(len(records)) {
		return true
	}
	return false
}

func ReadNRecords(url string, port int64, count int, dateRead bool) []model.IgcRecord {
	var records []model.IgcRecord
	var lastRec model.IgcRecord

	for len(records) < count {
		rec, _ := ReadRecord(url, port, dateRead)
		if lastRec == nil || rec.Time() != lastRec.Time() {
			records = append(records, rec)
			lastRec = rec
		}
		time.Sleep(50 * time.Millisecond)
	}

	return records
}

func ReadRecord(url string, port int64, dateRead bool) (model.IgcRecord, error) {
	response, err := http.Get(url + ":" + strconv.FormatInt(port, 10) + "/record")
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		record := model.IgcRecordImpl{string(contents)}

		var date time.Time

		if dateRead {
			date, _ = ReadDate(url, port)
		} else {
			date = time.Now()
		}
		record.SetDate(date)

		return record, nil
	}
	return nil, nil
}

func ReadDate(url string, port int64) (time.Time, error) {
	response, err := http.Get(url + ":" + strconv.FormatInt(port, 10) + "/date")
	if err != nil {
		return time.Now(), err
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return time.Now(), err
		}

		year, _ := strconv.ParseInt(string(contents[0:4]), 10, 64)
		month, _ := strconv.ParseInt(strings.TrimSpace(string(contents[4:6])), 10, 64)
		day, _ := strconv.ParseInt(strings.TrimSpace(string(contents[6:8])), 10, 64)

		date := time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)

		return date, nil
	}
	return time.Now(), nil
}

func PrintRecord(rec model.IgcRecord) {

	fmt.Print("Time:", rec.Time())
	fmt.Print(" Lat:", rec.Position().Lat())
	fmt.Print(" Lon:", rec.Position().Lon())
	fmt.Print(" Valid:", rec.IsValid())
	fmt.Print(" Preassure-Alt:", rec.PreassureAltitude(), "m")
	fmt.Print(" GPS-Alt:", rec.GpsAltitude(), "m\n")
}

func ConvertIgcFilenameToDate(filename string) time.Time {
	filename = strings.ToUpper(filename)

	day, err := strconv.Atoi(string(filename[2]))
	if err != nil {
		day = int(filename[2]) - 55
	}
	month, err := strconv.Atoi(string(filename[1]))
	if err != nil {
		day = int(filename[1]) - 55
	}

	year, _ := strconv.Atoi(string(filename[0]))
	year = year + (int(float64(time.Now().Year()) - math.Mod(float64(time.Now().Year()), 10.0)))

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
