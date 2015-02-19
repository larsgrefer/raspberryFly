package raspberryFly

import (
	"bufio"
	"os"
)

func AnalyseFile(pathname string) ([]IgcRecord, []Thermal) {
	inFile, _ := os.Open(pathname)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	var records []IgcRecord
	var thermals []Thermal

	var tempForThermals []IgcRecord

	circle := false

	for scanner.Scan() {
		if scanner.Text()[0] == 'B' {
			rec := IgcRecordImpl{scanner.Text()}
			rec.SetDate(ConvertIgcFilenameToDate(pathname))
			records = append(records, rec)
			if len(records) >= 11 {
				circle = findCircles(records[len(records)-11 : len(records)-1])
				if circle {
					tempForThermals = append(tempForThermals, rec)
				} else if len(tempForThermals) > 0 {
					t := ThermalImpl{tempForThermals}
					if t.HeightGain() > 0 && t.Duration().Minutes() >= 1 {
						thermals = append(thermals, t)
					}
					tempForThermals = make([]IgcRecord, 0)
				}
			}
		}
	}
	return records, thermals
}
