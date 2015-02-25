package analysis

import (
	"bufio"
	"fmt"
	"os"
	"raspberryFly"
	"raspberryFly/model"
)

func AnalyseFile(pathname string) ([]model.IgcRecord, []model.Thermal) {
	inFile, _ := os.Open(pathname)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	var records []model.IgcRecord
	var thermals []model.Thermal

	var tempForThermals []model.IgcRecord

	circle := false

	for scanner.Scan() {
		if scanner.Text()[0] == 'B' {
			rec := model.IgcRecordImpl{scanner.Text()}
			rec.SetDate(raspberryFly.ConvertIgcFilenameToDate(pathname))
			records = append(records, rec)
			if len(records) >= 13 {
				circle = raspberryFly.FindCircles(records[len(records)-13 : len(records)-1])
				if circle {
					tempForThermals = append(tempForThermals, rec)
				} else if len(tempForThermals) > 0 {
					t := model.ThermalImpl{circumsizeRecords(tempForThermals)}
					if t.HeightGain() > 0 && t.Duration().Minutes() >= 1 {
						thermals = append(thermals, t)
					}
					tempForThermals = make([]model.IgcRecord, 0)
				}
			}
		}
	}
	return records, thermals
}

func circumsizeRecords(records []model.IgcRecord) []model.IgcRecord {

	cutPoint := len(records)
	turnDegree := 0.0

	for i := len(records) - 1; i > 1 && turnDegree < 9; i-- {
		rec := records[i]
		lastRec := records[i-1]
		preLastRec := records[i-2]
		time := rec.Time().Sub(lastRec.Time()).Seconds()

		course := int(lastRec.Position().CourseTo(rec.Position()))
		lastCourse := int(preLastRec.Position().CourseTo(lastRec.Position()))

		turn := float64(course-lastCourse) / time

		fmt.Println(course, lastCourse, turn)

		if turn > 9 {
			cutPoint = i
			turnDegree = turn
		}
	}

	return records[0:cutPoint]
}
