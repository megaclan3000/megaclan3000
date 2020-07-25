package demoparser

import (
	"encoding/csv"
	"os"
	"strconv"

	dem "github.com/markus-wa/demoinfocs-golang"
	"github.com/markus-wa/demoinfocs-golang/common"
)

type Output struct {
	Frame   int
	Events  interface{}
	Players [][]string
}

func main() {
	f, err := os.Open("003425646141260169538_0105970676.dem")
	defer f.Close()
	checkError(err)

	p := dem.NewParser(f)

	var data []Output

	// parse frame by frame
	for ok := true; ok; ok, err = p.ParseNextFrame() {
		checkError(err)

		gs := p.GameState()
		frame := p.CurrentFrame()

		var players [][]string

		for _, player := range gs.Participants().Playing() {
			players = append(players, extractPlayerData(frame, player))
		}

		o := Output{
			Frame:   frame,
			Players: players,
		}

		data = append(data, o)
	}

	err = csvExport(data)
	checkError(err)
}

func extractPlayerData(frame int, player *common.Player) []string {
	return []string{
		strconv.Itoa(frame),
		player.Name,
		strconv.FormatInt(player.SteamID, 10),
		strconv.FormatFloat(player.Position.X, 'G', -1, 64),
		strconv.FormatFloat(player.Position.Y, 'G', -1, 64),
		strconv.FormatFloat(player.Position.Z, 'G', -1, 64),

		strconv.FormatFloat(player.LastAlivePosition.X, 'G', -1, 64),
		strconv.FormatFloat(player.LastAlivePosition.Y, 'G', -1, 64),
		strconv.FormatFloat(player.LastAlivePosition.Z, 'G', -1, 64),

		strconv.FormatFloat(player.Velocity.X, 'G', -1, 64),
		strconv.FormatFloat(player.Velocity.Y, 'G', -1, 64),
		strconv.FormatFloat(player.Velocity.Z, 'G', -1, 64),

		strconv.FormatFloat(float64(player.ViewDirectionX), 'G', -1, 64),
		strconv.FormatFloat(float64(player.ViewDirectionY), 'G', -1, 64),

		strconv.Itoa(player.Hp),
		strconv.Itoa(player.Armor),
		strconv.Itoa(player.Money),
		strconv.Itoa(player.CurrentEquipmentValue),
		strconv.Itoa(player.FreezetimeEndEquipmentValue),
		strconv.Itoa(player.RoundStartEquipmentValue),
		strconv.FormatBool(player.IsDucking),
		strconv.FormatBool(player.HasDefuseKit),
		strconv.FormatBool(player.HasHelmet),
		strconv.Itoa(player.AdditionalPlayerInformation.Kills),
		strconv.Itoa(player.AdditionalPlayerInformation.Deaths),
		strconv.Itoa(player.AdditionalPlayerInformation.Assists),
		strconv.Itoa(player.AdditionalPlayerInformation.Score),
		strconv.Itoa(player.AdditionalPlayerInformation.MVPs),
		strconv.Itoa(player.AdditionalPlayerInformation.TotalCashSpent),
		strconv.Itoa(player.AdditionalPlayerInformation.CashSpentThisRound),
	}
}

func csvExport(data []Output) error {
	file, err := os.OpenFile("result.csv", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// header
	header := []string{
		"Frame", "Name", "SteamID", "Position_X", "Position_Y", "Position_Z", "LastAlivePosition_X", "LastAlivePosition_Y", "LastAlivePosition_Z",
		"Velocity_X", "Velocity_Y", "Velocity_Z", "ViewDirectionX", "ViewDirectionY", "Hp", "Armor", "Money",
		"CurrentEquipmentValue", "FreezetimeEndEquipmentValue", "RoundStartEquipmentValue", "IsDucking", "HasDefuseKit",
		"HasHelmet", "Kills", "Deaths", "Assists", "Score", "MVPs", "TotalCashSpent", "CashSpentThisRound",
	}
	if err := writer.Write(header); err != nil {
		return err // let's return errors if necessary, rather than having a one-size-fits-all error handler
	}

	// data
	for _, frameData := range data {
		for _, player := range frameData.Players {
			err := writer.Write(player)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
