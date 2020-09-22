package demoparser

// https://github.com/markus-wa/demoinfocs-golang/blob/master/examples/print-events/print_events.go

import (
	// "github.com/mitchellh/hashstructure"

	"os"
	"path/filepath"
	"strings"

	"time"

	"github.com/megaclan3000/megaclan3000/internal/steamclient"

	demoinfocs "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	events "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

// DemoParser holds all methods to parse a demo file into a infostruct
type DemoParser struct {
	steamClient *steamclient.SteamClient
	parser      demoinfocs.Parser
	Match       *InfoStruct
	state       parsingState
}

// NewDemoParser constructor for a new demoparser
func NewDemoParser(client *steamclient.SteamClient) DemoParser {
	return DemoParser{
		steamClient: client,
		state: parsingState{
			Round:        0,
			RoundOngoing: false,
		},
	}
}

// Used while parsing to hold values while going through the ticks
type parsingState struct {
	Round        int // Current round
	RoundOngoing bool
	WarmupKills  []events.Kill
	currentTeam  common.Team
}

// Parse starts the parsing process and fills the infostruct with values
// gathered form the demo file
func (p *DemoParser) Parse(path string, m *InfoStruct) error {

	matchID := strings.Split(filepath.Base(path), "_")[0]
	m.MatchID = matchID
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
	p.parser.RegisterEventHandler(p.handlerScoreUpdated)
	p.parser.RegisterEventHandler(p.handlerWeaponFire)
	// p.RegisterEventHandler(handlerChatMessage)

	// Parse header and set general values
	err = p.setGeneral()

	if err != nil {
		return err
	}

	// Parse the demo returning errors
	err = p.parser.ParseToEnd()
	p.calculate()

	return err

}

func (p *DemoParser) playersBySteamID(steamID uint64) *common.Player {
	for _, v := range p.parser.GameState().Participants().All() {
		if v.SteamID64 == steamID {
			return v
		}
	}
	panic("player not found")
}

func (p *DemoParser) calculate() {

	for k, player := range p.Match.Players.Players {

		// Set Kills, Deaths, Assists, MVPs
		p.Match.Players.Players[k].Kills = p.playersBySteamID(player.Steamid64).Kills()
		p.Match.Players.Players[k].Deaths = p.playersBySteamID(player.Steamid64).Deaths()
		p.Match.Players.Players[k].Assists = p.playersBySteamID(player.Steamid64).Assists()
		p.Match.Players.Players[k].MVPs = p.playersBySteamID(player.Steamid64).MVPs()

		var playeradr int = 0

		for _, v := range allWeapons() {
			playeradr += p.Match.Players.Players[k].WeaponStats.getDamage(v)
		}

		p.Match.Players.Players[k].Adr = playeradr / len(p.Match.Rounds)

		// Calculate player's K/D
		if p.Match.Players.Players[k].Deaths != 0 {
			p.Match.Players.Players[k].Kd = float64(p.Match.Players.Players[k].Kills) / float64(p.Match.Players.Players[k].Deaths)
		}

		for _, round := range p.Match.Rounds {

			// Find player's kills and hs
			roundKills := 0
			for _, kill := range append(round.EnemyKills, round.ClanKills...) {
				if kill.Killer.Steamid64 == player.Steamid64 {
					if kill.IsHeadshot {
						p.Match.Players.Players[k].Headshots++
					}
					roundKills++
					// p.Match.Players.Players[k].WeaponStats.AddKill(kill.KillerWeapon, kill)
					// p.Match.Players.Players[k].WeaponStats.Kills[kill.KillerWeapon]++
				}
			}

			// Calculate player's he percentage
			if p.Match.Players.Players[k].Kills != 0 {
				p.Match.Players.Players[k].Hsprecent = float64(p.Match.Players.Players[k].Headshots) / float64(p.Match.Players.Players[k].Kills) * 100
			}

			// Set player's 3k, 4k, 5k rounds
			if roundKills == 5 {
				p.Match.Players.Players[k].Rounds5K++
			}
			if roundKills == 4 {
				p.Match.Players.Players[k].Rounds4K++
			}
			if roundKills == 3 {
				p.Match.Players.Players[k].Rounds3K++
			}
		}
	}
}

func (p *DemoParser) setGeneral() error {

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
	p.Match.General.ScoreClan = 0
	p.Match.General.ScoreEnemy = 0

	return nil
}

// NewScoreBoardPlayer constructor for a ScoreboardPlayer. Initializes some
// values with defaults
func (p *DemoParser) NewScoreBoardPlayer(player *common.Player) ScoreboardPlayer {

	name := "BOT"

	if !player.IsBot {
		name = player.Name
	}

	return ScoreboardPlayer{
		IsBot:            player.IsBot,
		IsClanMember:     player.Team == p.state.currentTeam,
		Name:             name,
		Rank:             0,
		Clantag:          player.ClanTag(),
		Steamid64:        player.SteamID64,
		AvatarURL:        p.steamClient.GetAvatarURL(player.SteamID64),
		Kills:            0,
		Deaths:           0,
		Assists:          0,
		Kd:               0,
		Adr:              0,
		Hsprecent:        0,
		Firstkills:       0,
		Firstdeaths:      0,
		Tradekills:       0,
		Tradedeaths:      0,
		Tradefirstkills:  0,
		Tradefirstdeaths: 0,
		Roundswonv5:      0,
		Roundswonv4:      0,
		Roundswonv3:      0,
		Rounds5K:         0,
		Rounds4K:         0,
		Rounds3K:         0,
		WeaponStats:      NewWeaponstats(),
		PlayerDamages:    NewPlayerDamages(),
	}
}

func (p *DemoParser) handlerWeaponFire(e events.WeaponFire) {

	if p.parser.GameState().IsWarmupPeriod() {
		return
	}

	p.playerByID(e.Shooter)
	shooter, err := p.Match.Players.PlayerNumByID(e.Shooter.SteamID64)

	if err != nil {
		panic(err)
	}
	p.Match.Players.Players[shooter].WeaponStats.addShot(e)

}

func (p *DemoParser) handlerKill(e events.Kill) {

	if e.Killer == nil || e.Victim == nil {
		return
	}

	// Skip all calculations for kills during warmup
	if p.parser.GameState().IsWarmupPeriod() {
		p.state.WarmupKills = append(p.state.WarmupKills, e)
		return
	}

	// Find killer
	killer := p.playerByID(e.Killer)
	killerNum, err := p.Match.Players.PlayerNumByID(e.Killer.SteamID64)
	if err != nil {
		panic(err)
	}

	p.Match.Players.Players[killerNum].WeaponStats.addKill(e)

	if e.IsHeadshot {
		p.Match.Players.Players[killerNum].WeaponStats.addHeadshot(e)
	}

	// Find victim
	victim := p.playerByID(e.Victim)
	victimNum, err := p.Match.Players.PlayerNumByID(e.Victim.SteamID64)
	if err != nil {
		panic(err)
	}

	kill := RoundKill{
		Time:         p.parser.CurrentTime(),
		IsHeadshot:   e.IsHeadshot,
		KillerWeapon: e.Weapon.Type,
		Killer:       killer,
		Victim:       victim,
	}

	if e.Assister != nil {
		assister := p.playerByID(e.Assister)
		p.Match.Players.addAssist(e.Assister.SteamID64)
		kill.Assister = assister
	}

	// Find fistkills and firstdeaths
	if e.Killer.Team == p.state.currentTeam {

		// Check if it's the first kill of the round
		if len(p.Match.Rounds[p.state.Round-1].ClanKills) == 0 {
			p.Match.Players.Players[killerNum].Firstkills++
			p.Match.Players.Players[victimNum].Firstdeaths++
		}

		for _, v := range p.Match.Rounds[p.state.Round-1].ClanKills {
			if v.Killer.Steamid64 == e.Victim.SteamID64 && ((p.parser.CurrentTime() - v.Time) < (5 * time.Second)) {
				p.Match.Players.Players[killerNum].Tradekills++
				p.Match.Players.Players[victimNum].Tradedeaths++

				if len(p.Match.Rounds[p.state.Round-1].ClanKills) == 0 {
					p.Match.Players.Players[killerNum].Tradefirstkills++
					p.Match.Players.Players[victimNum].Tradefirstdeaths++
				}
			}
		}

		// Append to clankills
		p.Match.Rounds[p.state.Round-1].ClanKills = append(p.Match.Rounds[p.state.Round-1].ClanKills, kill)
	} else {

		// Check if it's the first kill of the round
		if len(p.Match.Rounds[p.state.Round-1].EnemyKills) == 0 {
			p.Match.Players.Players[killerNum].Firstkills++
			p.Match.Players.Players[victimNum].Firstdeaths++
		}

		for _, v := range p.Match.Rounds[p.state.Round-1].EnemyKills {
			if v.Killer.Steamid64 == e.Victim.SteamID64 && ((p.parser.CurrentTime() - v.Time) < (5 * time.Second)) {
				p.Match.Players.Players[killerNum].Tradekills++
				p.Match.Players.Players[victimNum].Tradedeaths++
			}

			if len(p.Match.Rounds[p.state.Round-1].EnemyKills) == 0 {
				p.Match.Players.Players[killerNum].Tradefirstkills++
				p.Match.Players.Players[victimNum].Tradefirstdeaths++
			}
		}

		// Append to enemykills
		p.Match.Rounds[p.state.Round-1].EnemyKills = append(p.Match.Rounds[p.state.Round-1].EnemyKills, kill)
	}

	// Find 1v5, 1v4, 1v3
	if p.matesAlive(e.Killer) == 1 {
		switch p.matesAlive(e.Victim) {
		case 5:
			p.Match.Players.Players[killerNum].Roundswonv5++
		case 4:
			p.Match.Players.Players[killerNum].Roundswonv4++
		case 3:
			p.Match.Players.Players[killerNum].Roundswonv3++
		}
	}
}

func (p DemoParser) matesAlive(player *common.Player) int {
	alive := 0
	for _, v := range p.parser.GameState().Participants().Playing() {
		if v.IsAlive() && v.Team == player.Team {
			alive++
		}
	}
	return alive
}

func (p *DemoParser) handlerPlayerHurt(e events.PlayerHurt) {

	if e.Attacker == nil || e.Player == nil {
		return
	}

	for k, v := range p.Match.Players.Players {
		if v.Steamid64 == e.Attacker.SteamID64 {

			// Add damage stats for weapon
			p.Match.Players.Players[k].WeaponStats.addDamage(e)

			// Add hit stats for weapon
			p.Match.Players.Players[k].WeaponStats.addHit(e)

			// Add damage stats for PvP
			_ = p.playerByID(e.Player)
			victimNum, err := p.Match.Players.PlayerNumByID(e.Player.SteamID64)

			if err != nil {
				panic(err)
			}

			p.Match.Players.Players[k].addDamage(e.HealthDamage, &p.Match.Players.Players[victimNum])
			return
		}
	}
}

// func handlerChatMessage(e events.ChatMessage) {
// 	fmt.Printf("Chat - %s says: %s\n", formatPlayer(e.Sender), e.Text)
// }

// Handlers
func (p *DemoParser) handlerRankUpdate(e events.RankUpdate) {

	for k, v := range p.Match.Players.Players {
		if v.Steamid64 == e.SteamID64() {
			p.Match.Players.Players[k].Rank = e.RankNew
			return
		}
	}
	//TODO handle error
	panic("player not found setting rank")
}

func (p *DemoParser) handlerMatchStart(e events.MatchStart) {

	// Determine start team of clan
	for _, player := range p.parser.GameState().Participants().Playing() {
		if player.ClanTag() == "megaclan3000" {
			p.state.currentTeam = player.Team
			p.Match.MatchValid = true
		}
	}

	// Add all players to the match, that are no bots
	for _, ct := range p.parser.GameState().Participants().Playing() {

		if ct.IsBot {
			continue
		}

		player := p.NewScoreBoardPlayer(ct)

		p.Match.Players.Players = append(p.Match.Players.Players, player)
	}
}

func (p *DemoParser) handlerRoundStart(e events.RoundStart) {

	p.state.RoundOngoing = true

	// An new round has started, increase counter and add it to slice of the
	// output. The counter should be increased here and *not* in the RoundEnd
	// handler, sice there might happen things "between" the rounds, i.e in the
	// time when a round has ended but the new one has not yet started
	p.state.Round++

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

func (p *DemoParser) handlerBombPlanted(e events.BombPlanted) {
}

func (p *DemoParser) handlerBombDefused(e events.BombDefused) {
}

func (p *DemoParser) handlerBombExplode(e events.BombExplode) {
}

func (p *DemoParser) handlerScoreUpdated(e events.ScoreUpdated) {

	scoreCT := p.parser.GameState().TeamCounterTerrorists().Score()
	scoreT := p.parser.GameState().TeamTerrorists().Score()

	if p.state.currentTeam == common.TeamCounterTerrorists {
		p.Match.Rounds[p.state.Round-1].ScoreClan = scoreCT
		p.Match.Rounds[p.state.Round-1].ScoreEnemy = scoreT
		p.Match.General.ScoreClan = scoreCT
		p.Match.General.ScoreEnemy = scoreT
		return
	}

	if p.state.currentTeam == common.TeamTerrorists {
		p.Match.Rounds[p.state.Round-1].ScoreClan = scoreT
		p.Match.Rounds[p.state.Round-1].ScoreEnemy = scoreCT
		p.Match.General.ScoreClan = scoreT
		p.Match.General.ScoreEnemy = scoreCT
		return
	}

	// log.Warning("Scoreparsing did something strange", p.state.currentTeam)
}

func (p *DemoParser) handlerRoundEnd(e events.RoundEnd) {

	if !p.state.RoundOngoing {
		return
	}

	p.state.RoundOngoing = false

	// Set the winning team
	p.Match.Rounds[p.state.Round-1].TeamWon = e.Winner

	if e.Winner == p.state.currentTeam {
		p.Match.Rounds[p.state.Round-1].ClanWonRound = true
		p.Match.General.ScoreClan++
	} else {
		p.Match.General.ScoreEnemy++
	}

}
