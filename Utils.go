package raspberryFly

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func ReadNRecords(url string, port int64, count int, dateRead bool) []IgcRecord {
	var records []IgcRecord
	var lastRec IgcRecord

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

func ReadRecord(url string, port int64, dateRead bool) (IgcRecord, error) {
	response, err := http.Get(url + ":" + strconv.FormatInt(port, 10) + "/record")
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		record := IgcRecordImpl{string(contents)}

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

func PrintRecord(rec IgcRecord) {

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
