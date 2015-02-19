package raspberryFly

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

func ReadRecord(url string, port int64) (IgcRecord, error) {
	response, err := http.Get(url + ":" + strconv.FormatInt(port, 10) + "/record")
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return IgcRecordImpl{string(contents)}, nil
	}
	return nil, nil
}
