package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

func main() {
	matchToken := "CSGO-jZwRQ-Rry8M-uxRLr-ctbye-ZjenL"
	authCode := "9B3N-VU3HL-9BDM"
	key := "5021D36D8CD05B18B9A5284B8365F96E"
	steamID := "76561198092006615"

	url := "https://api.steampowered.com/ICSGOPlayers_730/GetNextMatchSharingCode/v1?key=" + key + "&steamid=" + steamID + "&steamidkey=" + authCode + "&knowncode=" + matchToken

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	// matchID, outcomeID, tokenID := shareCodeToID(matchToken)

	//   string location = AppSettings.GetFolderCachePath() + Path.DirectorySeparatorChar + demoName + ".bz2";
	//                     Uri uri = new Uri(url);
	//                     await Task.Factory.StartNew(() => webClient.DownloadFile(uri, location));
	shareCodeToID()

}

func shareCodeToID() (string, string, string) {

	shareCode := "CSGO-LoPQD-oNJmE-dVOkB-CLVWi-xOYQM"

	dico := "ABCDEFGHJKLMNOPQRSTUVWXYZabcdefhijkmnopqrstuvwxyz23456789"

	shareCode = strings.ReplaceAll(shareCode[5:], "-", "")

	bigID := big.NewInt(0)
	lenDico := big.NewInt(int64(len(dico)))

	for _, c := range funk.ReverseString(shareCode) {
		cBig := big.NewInt(int64(funk.IndexOf(dico, string(c))))
		bigID = bigID.Add(bigID.Mul(bigID, lenDico), cBig)
	}

	all := bigID.Bytes()
	if len(all) != 2*8+2 {
		all = append([]byte{0}, all...)
	}

	matchIDBytes := strconv.FormatUint(binary.LittleEndian.Uint64(all[0:8]), 10)
	outcomeIDBytes := strconv.FormatUint(binary.LittleEndian.Uint64(all[8:16]), 10)
	tokenIDByte := strconv.FormatUint(uint64(binary.LittleEndian.Uint16(all[16:18])), 10)

	fmt.Println(len(all))
	fmt.Println(all)
	fmt.Println(matchIDBytes)
	fmt.Println(outcomeIDBytes)
	fmt.Println(tokenIDByte)

	return matchIDBytes, outcomeIDBytes, tokenIDByte
}
