package demoparser

// https://github.com/markus-wa/demoinfocs-golang/blob/master/examples/print-events/print_events.go

import (
	// "github.com/mitchellh/hashstructure"
	"os"

	log "github.com/sirupsen/logrus"
	"strconv"
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
	currentTeam common.Team
}

func (p *MyParser) Parse(path string, m *InfoStruct) error {
	log.Warning("parsing file")
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
	p.parser.RegisterEventHandler(p.handlerMatchStart)
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
	log.Warning("set mapname to:", p.Match.General.MapName)

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
	log.Warning("PARSING HEADER")
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
		log.Warning(e.Weapon.String())
		var weapon string
		weapon = e.Weapon.String()
		log.Printf("%T\n", weapon)
		kill := RoundKill{
			KillerWeapon:     weapon,
			KillerSteamID64:  e.Killer.SteamID64,
			KillerTeamString: teamString(e.Killer.Team),
			VictimTeamString: teamString(e.Victim.Team),
			VictimSteamID64:  e.Victim.SteamID64,
		}

		if e.Killer.Team == p.state.currentTeam {
			p.Match.Rounds[p.state.Round-1].TeamClanKills = append(p.Match.Rounds[p.state.Round-1].TeamClanKills, kill)
		} else {
			p.Match.Rounds[p.state.Round-1].TeamEnemyKills = append(p.Match.Rounds[p.state.Round-1].TeamEnemyKills, kill)
		}
	}
}

func teamString(team common.Team) string {

	switch team {
	case common.TeamCounterTerrorists:
		return "CT"
	case common.TeamTerrorists:
		return "T"
	default:
		return ""
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

	//for  := range p.Match.General.PlayerInfos {
	//	// pl.RankIconURL = "/public/img/ranks/" + strconv.Itoa(e.RankOld) + ".png"
	//	pl.RankIconURL = getRankUrlFromSteamID64(e.SteamID64())
	//}

	////TODO set ranks icon URL
	//// for e := range collection {

	//// }
	//// 			RankIconURL: "public/img/ranks/" + ct.Ran,
	////TODO
	log.Printf("Rank Update: %d went from rank %d to rank %d, change: %f\n", e.SteamID32, e.RankOld, e.RankNew, e.RankChange)
	p.Match.General.PlayerInfos[e.Player.SteamID64].RankIconURL = "/public/img/ranks/" + strconv.Itoa(e.RankOld) + ".png"
}

func getAvatarUrlFromSteamID64(steamID uint64) string {

	// TODO implement and use througout the parser
	return "/public/img/avatars/other.jpg"
}

func getRankUrlFromSteamID64(steamID uint64) string {

	// TODO implement and use througout the parser
	return "/public/img/ranks/1.png"
}

func (p *MyParser) handlerMatchStart(e events.MatchStart) {

	p.Match.General.PlayerInfos = make(map[uint64]*ScoreboardTeamMemberInfo)

	var clanStartTeam common.Team
	for _, player := range p.parser.GameState().Participants().Playing() {
		if player.ClanTag() == "megaclan3000" {
			clanStartTeam = player.Team
		}
	}

	p.Match.General.PlayerInfos[0] = &ScoreboardTeamMemberInfo{
		AvatarURL:   getAvatarUrlFromSteamID64(0),
		Name:        "BOT",
		RankIconURL: getRankUrlFromSteamID64(0),
		ClanTag:     "",
	}

	for _, ct := range p.parser.GameState().Participants().Playing() {
		if ct.IsBot {
			continue
		}

		p.Match.General.PlayerInfos[ct.SteamID64] = &ScoreboardTeamMemberInfo{
			AvatarURL:   getAvatarUrlFromSteamID64(ct.SteamID64),
			Name:        ct.Name,
			RankIconURL: getRankUrlFromSteamID64(ct.SteamID64),
			ClanTag:     ct.ClanTag(),
		}

		// info := ScoreboardTeamMemberInfo{
		// 	AvatarURL:   getAvatarUrlFromSteamID64(ct.SteamID64),
		// 	Name:        ct.Name,
		// 	RankIconURL: getRankUrlFromSteamID64(ct.SteamID64),
		// 	ClanTag:     ct.ClanTag(),
		// }

		if ct.ClanTag() == "megaclan3000" {
			//TODO fetch and use correct images
			// avatarURL = "public/img/avatars" + strconv.FormatUint(ct.SteamID64, 10) + ".png"
		}

		line := ScoreboardLine{

			PlayerSteamID64:  ct.SteamID64,
			Kills:            0,
			Deaths:           0,
			Assists:          0,
			KDDiff:           0,
			KD:               0,
			ADR:              0,
			HSPrecent:        0,
			FirstKills:       0,
			FirstDeaths:      0,
			TradeKills:       0,
			TradeDeaths:      0,
			TradeFirstKills:  0,
			TradeFirstDeaths: 0,
			RoundsWonV5:      0,
			RoundsWonV4:      0,
			RoundsWonV3:      0,
			RoundsWonV2:      0,
			RoundsWonV1:      0,
			Rounds5k:         0,
			Rounds4k:         0,
			Rounds3k:         0,
			Rounds2k:         0,
			Rounds1k:         0,
			KAST:             0,
			HLTV:             0,
		}

		if ct.Team == clanStartTeam {
			p.Match.Scoreboard.TeamClan = append(p.Match.Scoreboard.TeamClan, line)
		} else {
			p.Match.Scoreboard.TeamEnemy = append(p.Match.Scoreboard.TeamEnemy, line)
		}

	}

}

func (p *MyParser) handlerRoundStart(e events.RoundStart) {

	// An new round has started, increase counter and add it to slice of the
	// output. The counter should be increased here and *not* in the RoundEnd
	// handler, sice there might happen things "between" the rounds, i.e in the
	// time when a round has ended but the new one has not yet started
	p.state.Round += 1

	for _, ct := range p.parser.GameState().TeamCounterTerrorists().Members() {
		if ct.ClanTag() == "megaclan3000" {
			p.state.currentTeam = common.TeamCounterTerrorists
			break
		}
	}

	for _, t := range p.parser.GameState().TeamTerrorists().Members() {
		if t.ClanTag() == "megaclan3000" {
			p.state.currentTeam = common.TeamTerrorists
			break
		}
	}

	round := ScoreboardRound{}
	p.Match.Rounds = append(p.Match.Rounds, round)

}

func (p *MyParser) handlerBombPlanted(e events.BombPlanted) {

}

func (p *MyParser) handlerBombDefused(e events.BombDefused) {
}

func (p *MyParser) handlerBombExplode(e events.BombExplode) {
}

func (p *MyParser) handlerRoundEnd(e events.RoundEnd) {

	// Set the winning team
	p.Match.Rounds[p.state.Round-1].TeamWon = e.Winner

	if e.Winner == p.state.currentTeam {
		p.Match.Rounds[p.state.Round-1].ClanWonRound = true
	}

	if p.state.currentTeam == common.TeamCounterTerrorists {
		p.Match.Rounds[p.state.Round-1].ScoreClan = p.parser.GameState().TeamCounterTerrorists().Score()
		p.Match.Rounds[p.state.Round-1].ScoreEnemy = p.parser.GameState().TeamTerrorists().Score()
	} else {
		p.Match.Rounds[p.state.Round-1].ScoreEnemy = p.parser.GameState().TeamCounterTerrorists().Score()
		p.Match.Rounds[p.state.Round-1].ScoreClan = p.parser.GameState().TeamTerrorists().Score()
	}
}
