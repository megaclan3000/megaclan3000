package demoparser

// https://github.com/markus-wa/demoinfocs-golang/blob/master/examples/print-events/print_events.go

import (
	// "github.com/mitchellh/hashstructure"
	"os"
	"time"

	demoinfocs "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	events "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

type MyParser struct {
	parser demoinfocs.Parser
	Result string
	Match  *InfoStruct
	state  parsingState
}

func NewMyParser() MyParser {
	return MyParser{
		state: parsingState{
			Round: 0,
		},
	}
}

// Used while parsing to hold values while going through the ticks
type parsingState struct {
	// Current round
	Round       int
	WarmupKills []events.Kill
}

func (p *MyParser) Parse(path string, m *InfoStruct) error {
	// Register handlers for events we care about
	p.Match = m
	var f *os.File
	var err error

	if f, err = os.Open(path); err != nil {
		return err
	}

	defer f.Close()

	p.parser = demoinfocs.NewParser(f)
	defer p.parser.Close()

	p.parser.RegisterEventHandler(p.handlerKill)
	p.parser.RegisterEventHandler(p.handlerRoundEnd)
	p.parser.RegisterEventHandler(p.handlerRoundStart)
	p.parser.RegisterEventHandler(p.handlerRankUpdate)
	p.parser.RegisterEventHandler(p.handlerPlayerHurt)
	p.parser.RegisterEventHandler(p.handlerBombPlanted)
	p.parser.RegisterEventHandler(p.handlerBombDefused)
	p.parser.RegisterEventHandler(p.handlerBombExplode)
	// p.RegisterEventHandler(handlerChatMessage)

	// Parse header and set general values
	p.setGeneral()

	// Parse the demo returning errors
	return p.parser.ParseToEnd()

}

func (p *MyParser) setGeneral() error {

	var header common.DemoHeader
	var err error

	if header, err = p.parser.ParseHeader(); err != nil {
		return err
	}

	//TODO implement this
	p.Match.General.MapName = header.MapName
	p.Match.General.MapIconURL = header.MapName
	p.Match.General.UploadTime = time.Now()
	p.Match.General.DemoLinkURL = "https:TODO/"

	return nil
}

func (p *MyParser) handlerKill(e events.Kill) {

	// Append kill to current round or to warmupKills
	if p.parser.GameState().IsWarmupPeriod() {
		p.state.WarmupKills = append(p.state.WarmupKills, e)
	} else {

		// p.Match.Rounds[p.state.Round].Kills = append(p.Match.Rounds[p.state.Round].Kills, e)
	}
}

func (p *MyParser) handlerPlayerHurt(e events.PlayerHurt) {

	//TODO detect and save teamdamage, skipping for now
	// Not sure if the hurt handler is triggered during warump as no teamdamage
	// is possible, better check anyway.
	// if !p.parser.GameState().IsWarmupPeriod() {

	// 	// Check if the player has done any damage at all in this round yet
	// 	if damageDone, ok1 := p.Match.Rounds[p.state.Round].TotalDamagesDone[e.Attacker.SteamID64]; ok1 {

	// 		// Check if the player has done damage to this victim in this round yet
	// 		if _, ok := damageDone.Victims[e.Player.SteamID64]; ok {
	// 			p.Match.Rounds[p.state.Round].TotalDamagesDone[e.Attacker.SteamID64].Victims[e.Player.SteamID64].Amount += e.HealthDamage
	// 		} else {
	// 			p.Match.Rounds[p.state.Round].TotalDamagesDone[e.Attacker.SteamID64].Victims[e.Player.SteamID64].Amount = e.HealthDamage
	// 		}
	// 	}
	// }
}

// func handlerChatMessage(e events.ChatMessage) {
// 	fmt.Printf("Chat - %s says: %s\n", formatPlayer(e.Sender), e.Text)
// }

// Handlers
func (p *MyParser) handlerRankUpdate(e events.RankUpdate) {
	//TODO
	// fmt.Printf("Rank Update: %d went from rank %d to rank %d, change: %f\n", e.SteamID32, e.RankOld, e.RankNew, e.RankChange)
}

func (p *MyParser) handlerRoundStart(e events.RoundStart) {

	// An new round has started, increase counter and add it to slice of the
	// output. The counter should be increased here and *not* in the RoundEnd
	// handler, sice there might happen things "between" the rounds, i.e in the
	// time when a round has ended but the new one has not yet started
	p.state.Round += 1

}

func (p *MyParser) handlerBombPlanted(e events.BombPlanted) {

}

func (p *MyParser) handlerBombDefused(e events.BombDefused) {
}

func (p *MyParser) handlerBombExplode(e events.BombExplode) {
}

func (p *MyParser) handlerRoundEnd(e events.RoundEnd) {

	// Set the winning team
	p.Match.Rounds[p.state.Round].TeamWon = e.Winner

}
