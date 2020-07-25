package demoparser

// https://github.com/markus-wa/demoinfocs-golang/blob/master/examples/print-events/print_events.go

import (
	"fmt"
	"os"

	demoinfocs "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	events "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

type MyParser struct {
	parser demoinfocs.Parser
	Result string
	Match  Match
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
	Round int
}

func (p *MyParser) Parse(path string) (Match, error) {
	// Register handlers for events we care about
	var f *os.File
	var err error
	var header common.DemoHeader

	if f, err = os.Open(path); err != nil {
		return p.Match, err
	}

	defer f.Close()

	p.parser = demoinfocs.NewParser(f)
	defer p.parser.Close()

	p.parser.RegisterEventHandler(p.handlerKill)
	p.parser.RegisterEventHandler(p.handlerRoundEnd)
	p.parser.RegisterEventHandler(p.handlerRoundStart)
	p.parser.RegisterEventHandler(p.handlerRankUpdate)
	p.parser.RegisterEventHandler(p.handlerPlayerHurt)
	// p.RegisterEventHandler(handlerChatMessage)

	// Parse header
	if header, err = p.parser.ParseHeader(); err != nil {
		return p.Match, err
	}

	fmt.Println("Map:", header.MapName)

	// Parse the demo
	err = p.parser.ParseToEnd()
	return p.Match, err
}

func (p *MyParser) handlerKill(e events.Kill) {

	var hs string
	if e.IsHeadshot {
		hs = " (HS)"
	}
	var wallBang string
	if e.PenetratedObjects > 0 {
		wallBang = " (WB)"
	}
	fmt.Printf("%s <%v%s%s> %s\n", formatPlayer(e.Killer), e.Weapon, hs, wallBang, formatPlayer(e.Victim))
}

func (p *MyParser) handlerPlayerHurt(e events.PlayerHurt) {}

// func handlerChatMessage(e events.ChatMessage) {
// 	fmt.Printf("Chat - %s says: %s\n", formatPlayer(e.Sender), e.Text)
// }

// Handlers
func (p *MyParser) handlerRankUpdate(e events.RankUpdate) {
	fmt.Printf("Rank Update: %d went from rank %d to rank %d, change: %f\n", e.SteamID32, e.RankOld, e.RankNew, e.RankChange)
}

func (p *MyParser) handlerRoundStart(e events.RoundStart) {

	// An new round has started, increase counter and add it to slice of the
	// output
	p.state.Round += 1

	p.Match.Rounds = append(p.Match.Rounds,
		Round{
			Number: p.state.Round,
		})
}

func (p *MyParser) handlerRoundEnd(e events.RoundEnd) {
	gs := p.parser.GameState()
	switch e.Winner {
	case common.TeamTerrorists:
		// Winner's score + 1 because it hasn't actually been updated yet
		fmt.Printf("Round finished: winnerSide=T  ; score=%d:%d\n", gs.TeamTerrorists().Score()+1, gs.TeamCounterTerrorists().Score())
	case common.TeamCounterTerrorists:
		fmt.Printf("Round finished: winnerSide=CT ; score=%d:%d\n", gs.TeamCounterTerrorists().Score()+1, gs.TeamTerrorists().Score())
	default:
		// Probably match medic or something similar
		fmt.Println("Round finished: No winner (tie)")
	}
}

func formatPlayer(p *common.Player) string {
	if p == nil {
		return "?"
	}

	switch p.Team {
	case common.TeamTerrorists:
		return "[T]" + p.Name
	case common.TeamCounterTerrorists:
		return "[CT]" + p.Name
	}

	return p.Name
}
