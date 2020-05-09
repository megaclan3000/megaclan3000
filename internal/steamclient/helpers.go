package steamclient

import (
	"encoding/json"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"

	"net/http"
	"strconv"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target interface{}) error {
	// log.Println("Downloading:", url)
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}

	if r.StatusCode != 200 {
		log.Warn("Failed to get data from:")
		log.Warn("'" + url + "'")
		log.Warn(r.StatusCode, http.StatusText(r.StatusCode))
		return errors.New("Failed to fetch from URL")
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func divideStringFloats(a, b string) string {

	if aF, err := strconv.ParseFloat(a, 64); err == nil {
		if bF, err := strconv.ParseFloat(b, 64); err == nil {
			return fmt.Sprintf("%.4f", aF/bF)
		}
	}
	return ""
}
