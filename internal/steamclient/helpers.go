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
	return "n/a"
}

func getWeaponByID(id string) weapon {

	weaponNames := map[string]string{
		"1":  "deagle",
		"2":  "elite",
		"3":  "fiveseven",
		"4":  "glock",
		"5":  "p228",
		"6":  "usp",
		"7":  "ak47",
		"8":  "aug",
		"9":  "awp",
		"10": "famas",
		"11": "g3sg1",
		"12": "galil",
		"13": "galilar",
		"14": "m249",
		"15": "m3",
		"16": "m4a1",
		"17": "mac10",
		"18": "mp5navy",
		"19": "p90",
		"20": "scout",
		"21": "sg550",
		"22": "sg552",
		"23": "tmp",
		"24": "ump45",
		"25": "xm1014",
		"26": "bizon",
		"27": "mag7",
		"28": "negev",
		"29": "sawedoff",
		"30": "tec9",
		"31": "taser",
		"32": "hkp2000",
		"33": "mp7",
		"34": "mp9",
		"35": "nova",
		"36": "p250",
		"37": "scar17",
		"38": "scar20",
		"39": "sg556",
		"40": "ssg08",
		"41": "knifegg",
		"42": "knife",
		"43": "flashbang",
		"44": "hegrenade",
		"45": "smokegrenade",
		"46": "molotov",
		"47": "decoy",
		"48": "incgrenade",
		"49": "c4",
	}
	return weapon{
		ID:       id,
		Name:     weaponNames[id],
		IconPath: "/public/img/weapons/" + weaponNames[id] + ".jpg",
	}
}

type weapon struct {
	ID       string
	Name     string
	IconPath string
}
