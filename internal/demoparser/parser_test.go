package demoparser

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/google/go-cmp/cmp/cmpopts"
	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

func parseDurationNoErr(dur string) time.Duration {
	d, err := time.ParseDuration(dur)
	if err != nil {
		panic(err)
	}
	return d
}

func TestMyParser_Parse(t *testing.T) {

	// FIXME: I'm sure there is a better way of passing all these values, but it
	// should be good enought for the tests. If you care to clean it up, please do.
	demo1Times := map[string]string{
		"1Start": "1m49.046874112s", "1End": "3m38.953121792s",
		"2Start": "3m46.078130176s", "2End": "4m39.187488768s",
		"3Start": "4m46.28123648s", "3End": "6m5.24998656s",
		"4Start": "6m12.249985024s", "4End": "7m35.31250688s",
		"5Start": "7m42.406254592s", "5End": "8m20.406255616s",
		"6Start": "8m27.437514752s", "6End": "10m13.093736448s",
		"7Start": "10m20.203147264s", "7End": "11m43.078137856s",
		"8Start": "11m50.156222464s", "8End": "12m52.656267264s",
		"9Start": "12m59.65623296s", "9End": "14m10.437472256s",
		"10Start": "14m17.546883072s", "10End": "16m35.20315392s",
		"11Start": "16m42.26564096s", "11End": "18m16.859385856s",
		"12Start": "18m23.968731136s", "12End": "20m33.874976768s",
		"13Start": "20m40.875008s", "13End": "21m53.859436544s",
		"14Start": "22m0.859336704s", "14End": "22m55.765659648s",
		"15Start": "23m2.875004928s", "15End": "24m18.781290496s",
		"16Start": "24m33.85942016s", "16End": "25m31.015593984s",
		"17Start": "25m38.07814656s", "17End": "27m46.734358528s",
		"18Start": "27m53.812508672s", "18End": "28m59.249942528s",
		"19Start": "29m6.328092672s", "19End": "29m52.859439104s",
		"20Start": "29m59.859339264s", "20End": "31m4.984428544s",
		"21Start": "31m11.984328704s", "21End": "32m49.234378752s",
		"22Start": "32m56.234409984s", "22End": "33m59.062462464s",
	}

	// First round was the only one with 10 real players, after that, on player
	// left and a bot (no steamid) was added
	demo1AllPlayersFirstRound := []common.Player{
		{SteamID64: 76561198871105662},
		{SteamID64: 76561197978562286},
		{SteamID64: 76561198070048497},
		{SteamID64: 76561198092006615},
		{SteamID64: 76561198103322640},
		{SteamID64: 76561198104947907},
		{SteamID64: 76561198114207134},
		{SteamID64: 76561198261800498},
		{SteamID64: 76561198882848278},
		{SteamID64: 76561199000235131},
	}
	demo1AllPlayers := []common.Player{
		{}, // BOT has no steamID
		{SteamID64: 76561197978562286},
		{SteamID64: 76561198070048497},
		{SteamID64: 76561198092006615},
		{SteamID64: 76561198103322640},
		{SteamID64: 76561198104947907},
		{SteamID64: 76561198114207134},
		{SteamID64: 76561198261800498},
		{SteamID64: 76561198882848278},
		{SteamID64: 76561199000235131},
	}

	tests := []struct {
		name    string
		path    string
		want    Match
		wantErr bool
	}{
		{
			name: "Parse demo1 file",
			path: "testdata/demo1.dem",
			want: Match{
				// TODO find out how to get a proper ID
				ID: 4619025276304667104,

				// TODO find out how to get the time when the match was played
				UploadTime: time.Date(2020, time.July, 1, 3, 4, 5, 6, time.UTC),
				Map:        "de_mirage",

				// Match in the demo took 22 rounds
				// TODO fill in correct values
				WarmupKills: []events.Kill{
					{
						Victim: &common.Player{Name: "Allen"},
						Killer: &common.Player{Name: "randolf"},
					},
					{
						Victim: &common.Player{Name: "Scarlett O'Hara"},
						Killer: &common.Player{Name: "randolf"},
					},
					{
						Victim: &common.Player{Name: "Rhett Butler"},
						Killer: &common.Player{Name: "randolf"},
					},
					{
						Victim: &common.Player{Name: "Scarlett O'Hara"},
						Killer: &common.Player{Name: "randolf"},
					},
					{
						Victim: &common.Player{Name: "randolf"},
						Killer: &common.Player{Name: "Lucifer"},
					},
					{
						Victim: &common.Player{Name: "Rhett Butler"},
						Killer: &common.Player{Name: "der rote Rivale"},
					},
					{
						Victim: &common.Player{Name: "der rote Rivale"},
						Killer: &common.Player{Name: "Lucifer"},
					},
					{
						Victim: &common.Player{Name: "salatkopf"},
						Killer: &common.Player{Name: "Wesley"},
					},
					{
						Victim: &common.Player{Name: "Wesley"},
						Killer: &common.Player{Name: "killer strike 2"},
					},
					{
						Victim: &common.Player{Name: "randolf"},
						Killer: &common.Player{Name: "Lucifer"},
					},
					{
						Victim: &common.Player{Name: "randolf"},
						Killer: &common.Player{Name: "my name is гон(дон)"},
					},
					{
						Victim: &common.Player{Name: "my name is гон(дон)"},
						Killer: &common.Player{Name: "killer strike 2"},
					},
					{
						Victim: &common.Player{Name: "Scarlett O'Hara"},
						Killer: &common.Player{Name: "der rote Rivale"},
					},
					{
						Victim: &common.Player{Name: "afkR0多y-"},
						Killer: &common.Player{Name: "der rote Rivale"},
					},
					{
						Victim: &common.Player{Name: "Rhett Butler"},
						Killer: &common.Player{Name: "der rote Rivale"},
					},
					{
						Victim: &common.Player{Name: "randolf"},
						Killer: &common.Player{Name: "my name is гон(дон)"},
					},
					{
						Victim: &common.Player{Name: "my name is гон(дон)"},
						Killer: &common.Player{Name: "killer strike 2"},
					},
					{
						Victim: &common.Player{Name: "Lucifer"},
						Killer: &common.Player{Name: "der rote Rivale"},
					},
					{
						Victim: &common.Player{Name: "der rote Rivale"},
						Killer: &common.Player{Name: "my name is гон(дон)"},
					},
				},
				Rounds: Rounds{
					1: {
						TimeStart:    parseDurationNoErr(demo1Times["1Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["1End"]),
						TeamWon:      common.TeamTerrorists,
						BombPlanted:  true,
						BombDefused:  false,
						BombExploded: false,
						Players:      demo1AllPlayersFirstRound,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "Rhett Butler"}, Killer: &common.Player{Name: "der rote Rivale"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "randolf"}},
							{Victim: &common.Player{Name: "randolf"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "salatkopf"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "killer strike 2"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
							{Victim: &common.Player{Name: "der rote Rivale"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "Kapt'n Turbot"}, Killer: &common.Player{Name: "Lucifer"}},
						},
					},
					2: {
						TimeStart:    parseDurationNoErr(demo1Times["2Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["2End"]),
						TeamWon:      common.TeamTerrorists,
						BombPlanted:  true,
						BombDefused:  false,
						BombExploded: false,
						Players:      demo1AllPlayers,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "afkR0多y-"}, Killer: &common.Player{Name: "afkR0多y-"}},
							{Victim: &common.Player{Name: "Kapt'n Turbot"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "der rote Rivale"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "salatkopf"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "randolf"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "killer strike 2"}, Killer: &common.Player{Name: "Rhett Butler"}},
						},
					},
					3: {
						TimeStart:    parseDurationNoErr(demo1Times["3Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["3End"]),
						TeamWon:      common.TeamTerrorists,
						BombPlanted:  true,
						BombDefused:  false,
						BombExploded: false,
						Players:      demo1AllPlayers,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "randolf"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "salatkopf"}},
							{Victim: &common.Player{Name: "Lucifer"}, Killer: &common.Player{Name: "salatkopf"}},
							{Victim: &common.Player{Name: "salatkopf"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
							{Victim: &common.Player{Name: "killer strike 2"}, Killer: &common.Player{Name: "Rhett Butler"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
							{Victim: &common.Player{Name: "Kapt'n Turbot"}, Killer: &common.Player{Name: "Rhett Butler"}},
							{Victim: &common.Player{Name: "der rote Rivale"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
						},
					},
					4: {
						TimeStart:    parseDurationNoErr(demo1Times["4Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["4End"]),
						TeamWon:      common.TeamTerrorists,
						BombPlanted:  true,
						BombDefused:  false,
						BombExploded: false,
						Players:      demo1AllPlayers,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "randolf"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "Irwin"}, Killer: &common.Player{Name: "salatkopf"}},
							{Victim: &common.Player{Name: "salatkopf"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "der rote Rivale"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "Kapt'n Turbot"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
							{Victim: &common.Player{Name: "killer strike 2"}, Killer: &common.Player{Name: "Lucifer"}},
						},
					},
					5: {
						TimeStart:    parseDurationNoErr(demo1Times["5Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["5End"]),
						TeamWon:      common.TeamTerrorists,
						BombPlanted:  true,
						BombDefused:  false,
						BombExploded: false,
						Players:      demo1AllPlayers,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "Kapt'n Turbot"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "der rote Rivale"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "randolf"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "salatkopf"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "killer strike 2"}},
							{Victim: &common.Player{Name: "killer strike 2"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
						},
					},
					6: {
						TimeStart:    parseDurationNoErr(demo1Times["6Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["6End"]),
						TeamWon:      common.TeamTerrorists,
						BombPlanted:  true,
						BombDefused:  false,
						BombExploded: false,
						Players:      demo1AllPlayers,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "Rhett Butler"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
							{Victim: &common.Player{Name: "Kapt'n Turbot"}, Killer: &common.Player{Name: "Rhett Butler"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "der rote Rivale"}},
							{Victim: &common.Player{Name: "Rhett Butler"}, Killer: &common.Player{Name: "randolf"}},
							{Victim: &common.Player{Name: "randolf"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "Lucifer"}, Killer: &common.Player{Name: "salatkopf"}},
							{Victim: &common.Player{Name: "salatkopf"}, Killer: &common.Player{Name: "Scarlett O'Hara"}, PenetratedObjects: 1},
							{Victim: &common.Player{Name: "der rote Rivale"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
							{Victim: &common.Player{Name: "killer strike 2"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
						},
					},
					7: {
						TimeStart:    parseDurationNoErr(demo1Times["7Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["7End"]),
						TeamWon:      common.TeamTerrorists,
						BombPlanted:  true,
						BombDefused:  false,
						BombExploded: false,
						Players:      demo1AllPlayers,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "salatkopf"}},
							{Victim: &common.Player{Name: "randolf"}, Killer: &common.Player{Name: "Rhett Butler"}},
							{Victim: &common.Player{Name: "salatkopf"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "Rhett Butler"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
							{Victim: &common.Player{Name: "killer strike 2"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
							{Victim: &common.Player{Name: "der rote Rivale"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
							{Victim: &common.Player{Name: "Lucifer"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
							{Victim: &common.Player{Name: "Scarlett O'Hara"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
							{Victim: &common.Player{Name: "Kapt'n Turbot"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
						},
					},
					8: {
						TimeStart:    parseDurationNoErr(demo1Times["8Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["8End"]),
						TeamWon:      common.TeamTerrorists,
						BombPlanted:  true,
						BombDefused:  false,
						BombExploded: false,
						Players:      demo1AllPlayers,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "randolf"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "salatkopf"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
							{Victim: &common.Player{Name: "killer strike 2"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
							{Victim: &common.Player{Name: "Lucifer"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
							{Victim: &common.Player{Name: "Kapt'n Turbot"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
							{Victim: &common.Player{Name: "Scarlett O'Hara"}, Killer: &common.Player{Name: "der rote Rivale"}},
							{Victim: &common.Player{Name: "der rote Rivale"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
						},
					},
					9: {
						TimeStart:    parseDurationNoErr(demo1Times["9Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["9End"]),
						TeamWon:      common.TeamTerrorists,
						BombPlanted:  true,
						BombDefused:  false,
						BombExploded: false,
						Players:      demo1AllPlayers,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "killer strike 2"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "salatkopf"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
							{Victim: &common.Player{Name: "randolf"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
							{Victim: &common.Player{Name: "Lucifer"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "der rote Rivale"}},
							{Victim: &common.Player{Name: "Rhett Butler"}, Killer: &common.Player{Name: "der rote Rivale"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
							{Victim: &common.Player{Name: "der rote Rivale"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
							{Victim: &common.Player{Name: "Kapt'n Turbot"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
						},
					},
					10: {
						TimeStart:    parseDurationNoErr(demo1Times["10Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["10End"]),
						TeamWon:      common.TeamCounterTerrorists,
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: false,
						Players:      demo1AllPlayers,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "Irwin"}, Killer: &common.Player{Name: "killer strike 2"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "killer strike 2"}},
							{Victim: &common.Player{Name: "killer strike 2"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "Rhett Butler"}, Killer: &common.Player{Name: "salatkopf"}},
							{Victim: &common.Player{Name: "der rote Rivale"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
							{Victim: &common.Player{Name: "Scarlett O'Hara"}, Killer: &common.Player{Name: "randolf"}},
							{Victim: &common.Player{Name: "salatkopf"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "randolf"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "Lucifer"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
						},
					},
					11: {
						TimeStart:    parseDurationNoErr(demo1Times["11Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["11End"]),
						TeamWon:      common.TeamTerrorists,
						BombPlanted:  true,
						BombDefused:  false,
						BombExploded: false,
						Players:      demo1AllPlayers,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "salatkopf"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "killer strike 2"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "der rote Rivale"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "der rote Rivale"}},
							{Victim: &common.Player{Name: "der rote Rivale"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
							{Victim: &common.Player{Name: "Kapt'n Turbot"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
							{Victim: &common.Player{Name: "Lucifer"}, Killer: &common.Player{Name: "randolf"}},
							{Victim: &common.Player{Name: "Scarlett O'Hara"}, Killer: &common.Player{Name: "randolf"}},
							{Victim: &common.Player{Name: "randolf"}, Killer: &common.Player{Name: "Rhett Butler"}},
						},
					},
					12: {
						TimeStart:    parseDurationNoErr(demo1Times["12Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["12End"]),
						TeamWon:      common.TeamCounterTerrorists,
						BombPlanted:  false,
						BombDefused:  false,
						BombExploded: false,
						Players:      demo1AllPlayers,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "randolf"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "killer strike 2"}},
							{Victim: &common.Player{Name: "Rhett Butler"}, Killer: &common.Player{Name: "killer strike 2"}},
							{Victim: &common.Player{Name: "salatkopf"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "Lucifer"}, Killer: &common.Player{Name: "killer strike 2"}},
							{Victim: &common.Player{Name: "der rote Rivale"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "killer strike 2"}},
							{Victim: &common.Player{Name: "killer strike 2"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
						},
					},
					13: {
						TimeStart:    parseDurationNoErr(demo1Times["13Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["13End"]),
						TeamWon:      common.TeamTerrorists,
						BombPlanted:  true,
						BombDefused:  false,
						BombExploded: false,
						Players:      demo1AllPlayers,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "salatkopf"}},
							{Victim: &common.Player{Name: "killer strike 2"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "Scarlett O'Hara"}, Killer: &common.Player{Name: "randolf"}},
							{Victim: &common.Player{Name: "randolf"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "Rhett Butler"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
							{Victim: &common.Player{Name: "salatkopf"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "Kapt'n Turbot"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "der rote Rivale"}},
							{Victim: &common.Player{Name: "der rote Rivale"}, Killer: &common.Player{Name: "Lucifer"}, PenetratedObjects: 1},
						},
					},
					14: {
						TimeStart:    parseDurationNoErr(demo1Times["14Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["14End"]),
						TeamWon:      common.TeamTerrorists,
						BombPlanted:  false,
						BombDefused:  false,
						BombExploded: false,
						Players:      demo1AllPlayers,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "killer strike 2"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "Rhett Butler"}, Killer: &common.Player{Name: "randolf"}},
							{Victim: &common.Player{Name: "randolf"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
							{Victim: &common.Player{Name: "salatkopf"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
							{Victim: &common.Player{Name: "der rote Rivale"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "Rhett Butler"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
							{Victim: &common.Player{Name: "Scarlett O'Hara"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
							{Victim: &common.Player{Name: "Kapt'n Turbot"}, Killer: &common.Player{Name: "Lucifer"}},
						},
					},
					15: {
						TimeStart:    parseDurationNoErr(demo1Times["15Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["15End"]),
						TeamWon:      common.TeamCounterTerrorists,
						BombPlanted:  false,
						BombDefused:  false,
						BombExploded: false,
						Players:      demo1AllPlayers,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "Rhett Butler"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
							{Victim: &common.Player{Name: "Lucifer"}, Killer: &common.Player{Name: "salatkopf"}},
							{Victim: &common.Player{Name: "Scarlett O'Hara"}, Killer: &common.Player{Name: "salatkopf"}},
							{Victim: &common.Player{Name: "killer strike 2"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "randolf"}, Killer: &common.Player{Name: "Rhett Butler"}},
							{Victim: &common.Player{Name: "der rote Rivale"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "salatkopf"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
							{Victim: &common.Player{Name: "Rhett Butler"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
						},
					},
					16: {
						TimeStart:    parseDurationNoErr(demo1Times["16Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["16End"]),
						TeamWon:      common.TeamTerrorists,
						BombPlanted:  true,
						BombDefused:  false,
						BombExploded: false,
						Players:      demo1AllPlayers,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "salatkopf"}, Killer: &common.Player{Name: "Rhett Butler"}},
							{Victim: &common.Player{Name: "Rhett Butler"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
							{Victim: &common.Player{Name: "Scarlett O'Hara"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
							{Victim: &common.Player{Name: "der rote Rivale"}, Killer: &common.Player{Name: "Rhett Butler"}},
							{Victim: &common.Player{Name: "Lucifer"}, Killer: &common.Player{Name: "killer strike 2"}},
							{Victim: &common.Player{Name: "Rhett Butler"}, Killer: &common.Player{Name: "killer strike 2"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
						},
					},
					17: {
						TimeStart:    parseDurationNoErr(demo1Times["17Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["17End"]),
						TeamWon:      common.TeamCounterTerrorists,
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: false,
						Players:      demo1AllPlayers,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "Kapt'n Turbot"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "salatkopf"}},
							{Victim: &common.Player{Name: "salatkopf"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "killer strike 2"}},
							{Victim: &common.Player{Name: "killer strike 2"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "Lucifer"}, Killer: &common.Player{Name: "der rote Rivale"}},
							{Victim: &common.Player{Name: "randolf"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
							{Victim: &common.Player{Name: "Scarlett O'Hara"}, Killer: &common.Player{Name: "der rote Rivale"}},
							{Victim: &common.Player{Name: "der rote Rivale"}, Killer: &common.Player{Name: "Rhett Butler"}},
						},
					},
					18: {
						TimeStart:    parseDurationNoErr(demo1Times["18Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["18End"]),
						TeamWon:      common.TeamCounterTerrorists,
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: false,
						Players:      demo1AllPlayers,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "killer strike 2"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "Kapt'n Turbot"}, Killer: &common.Player{Name: "Rhett Butler"}},
							{Victim: &common.Player{Name: "Rhett Butler"}, Killer: &common.Player{Name: "salatkopf"}},
							{Victim: &common.Player{Name: "der rote Rivale"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "randolf"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "salatkopf"}, Killer: &common.Player{Name: "Lucifer"}},
						},
					},
					19: {
						TimeStart:    parseDurationNoErr(demo1Times["19Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["19End"]),
						TeamWon:      common.TeamTerrorists,
						BombPlanted:  false,
						BombDefused:  false,
						BombExploded: false,
						Players:      demo1AllPlayers,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "Irwin"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
							{Victim: &common.Player{Name: "Kapt'n Turbot"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "der rote Rivale"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
							{Victim: &common.Player{Name: "randolf"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "killer strike 2"}, Killer: &common.Player{Name: "Rhett Butler"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "salatkopf"}},
							{Victim: &common.Player{Name: "Lucifer"}, Killer: &common.Player{Name: "salatkopf"}},
							{Victim: &common.Player{Name: "Rhett Butler"}, Killer: &common.Player{Name: "salatkopf"}},
							{Victim: &common.Player{Name: "Scarlett O'Hara"}, Killer: &common.Player{Name: "salatkopf"}},
						},
					},
					20: {
						TimeStart:    parseDurationNoErr(demo1Times["20Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["20End"]),
						TeamWon:      common.TeamTerrorists,
						BombPlanted:  true,
						BombDefused:  false,
						BombExploded: false,
						Players:      demo1AllPlayers,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "Rhett Butler"}, Killer: &common.Player{Name: "killer strike 2"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "killer strike 2"}},
							{Victim: &common.Player{Name: "killer strike 2"}, Killer: &common.Player{Name: "Scarlett O'Hara"}},
							{Victim: &common.Player{Name: "Rhett Butler"}, Killer: &common.Player{Name: "salatkopf"}},
							{Victim: &common.Player{Name: "Scarlett O'Hara"}, Killer: &common.Player{Name: "der rote Rivale"}},
							{Victim: &common.Player{Name: "Lucifer"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
						},
					},
					21: {
						TimeStart:    parseDurationNoErr(demo1Times["21Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["21End"]),
						TeamWon:      common.TeamCounterTerrorists,
						BombPlanted:  false,
						BombDefused:  false,
						BombExploded: false,
						Players:      demo1AllPlayers,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "salatkopf"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "randolf"}},
							{Victim: &common.Player{Name: "randolf"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "der rote Rivale"}, Killer: &common.Player{Name: "Rhett Butler"}},
							{Victim: &common.Player{Name: "Scarlett O'Hara"}, Killer: &common.Player{Name: "killer strike 2"}},
							{Victim: &common.Player{Name: "Rhett Butler"}, Killer: &common.Player{Name: "Kapt'n Turbot"}},
							{Victim: &common.Player{Name: "killer strike 2"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "Kapt'n Turbot"}, Killer: &common.Player{Name: "Lucifer"}},
						},
					},
					22: {
						TimeStart:    parseDurationNoErr(demo1Times["22Start"]),
						TimeEnd:      parseDurationNoErr(demo1Times["22End"]),
						TeamWon:      common.TeamCounterTerrorists,
						BombPlanted:  false,
						BombDefused:  false,
						BombExploded: false,
						Players:      demo1AllPlayers,
						Kills: []events.Kill{
							{Victim: &common.Player{Name: "killer strike 2"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "Kapt'n Turbot"}, Killer: &common.Player{Name: "Lucifer"}},
							{Victim: &common.Player{Name: "my name is гон(дон)"}, Killer: &common.Player{Name: "randolf"}},
							{Victim: &common.Player{Name: "Lucifer"}, Killer: &common.Player{Name: "salatkopf"}},
							{Victim: &common.Player{Name: "randolf"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "der rote Rivale"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
							{Victim: &common.Player{Name: "salatkopf"}, Killer: &common.Player{Name: "my name is гон(дон)"}},
						},
					},
				},
			},

			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewMyParser()

			got, err := p.Parse(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("MyParser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// When comparing common.Player, just compare the SteamID64 to keep the test data shorter.
			// BOTs have *no* steamID, so we must add a empty common.Player to
			// the list of players in the test data if they are on a team
			opt1 := cmp.Comparer(func(x, y common.Player) bool {
				return x.SteamID64 == y.SteamID64
			})

			// When comoparing kills, just compare by attacker and victims Name
			opt2 := cmp.Comparer(func(x, y events.Kill) bool {
				return x.Killer.Name == y.Killer.Name && x.Victim.Name == y.Victim.Name
			})

			opt3 := cmpopts.IgnoreFields(Match{}, "UploadTime")

			// Sort the slices by SteamID64 so the tests don't fail. The order
			// in which they come from the parser is not reliably the same
			// every run
			sorter := cmpopts.SortSlices(func(x, y common.Player) bool { return x.SteamID64 < y.SteamID64 })
			if diff := cmp.Diff(tt.want, got, sorter, opt1, opt2, opt3); diff != "" {
				t.Errorf("MyParser.Parse() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// func less(x, y common.Player) bool {
// 	return x.SteamID64 < y.SteamID64
// }
