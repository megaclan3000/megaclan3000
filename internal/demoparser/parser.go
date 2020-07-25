package demoparser

// https://github.com/markus-wa/demoinfocs-golang/blob/master/examples/print-events/print_events.go

import (
	"fmt"
	"os"

	demoinfocs "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	events "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

var p demoinfocs.Parser

type MyParser struct {
	parser demoinfocs.Parser
	Result string
}

func NewMyParser(path string) (MyParser, error) {

	var p MyParser
	var f *os.File
	var err error
	var header common.DemoHeader

	if f, err = os.Open(path); err != nil {
		return p, err
	}

	defer f.Close()

	p.parser = demoinfocs.NewParser(f)
	defer p.parser.Close()

	// Parse header
	if header, err = p.parser.ParseHeader(); err != nil {
		return p, err
	}

	fmt.Println("Map:", header.MapName)

	// Register handlers for events we care about
	p.parser.RegisterEventHandler(handlerKill)
	p.parser.RegisterEventHandler(handlerRoundEnd)
	p.parser.RegisterEventHandler(handlerRoundStart)
	p.parser.RegisterEventHandler(handlerRankUpdate)
	p.parser.RegisterEventHandler(handlerPlayerHurt)
	// p.RegisterEventHandler(handlerChatMessage)

	return p, nil

}

func (p *MyParser) Parse() error {
	// Parse to end
	err := p.parser.ParseToEnd()
	return err
}

func handlerKill(e events.Kill) {

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

func handlerPlayerHurt(e events.PlayerHurt) {

}

// func handlerChatMessage(e events.ChatMessage) {
// 	fmt.Printf("Chat - %s says: %s\n", formatPlayer(e.Sender), e.Text)
// }

// Handlers
func handlerRankUpdate(e events.RankUpdate) {
	fmt.Printf("Rank Update: %d went from rank %d to rank %d, change: %f\n", e.SteamID32, e.RankOld, e.RankNew, e.RankChange)
}

func handlerRoundStart(e events.RoundStart) {
	//TODO
}

func handlerRoundEnd(e events.RoundEnd) {
	gs := p.GameState()
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
