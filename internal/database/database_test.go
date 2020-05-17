package database

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/google/go-cmp/cmp"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pinpox/megaclan3000/internal/steamclient"
)

var (
	db       *DataStorage
	fixtures *testfixtures.Loader
)

func prepareDB() {

	var err error
	if db, err = NewDataStorage("../../test/database/test.db"); err != nil {
		panic(err)
	}

	if fixtures, err = testfixtures.New(
		testfixtures.Database(db.db),
		testfixtures.Dialect("sqlite"),
		testfixtures.Directory(
			"../../test/database/fixtures",
		),
	); err != nil {
		panic(err)
	}

	if err = fixtures.Load(); err != nil {
		panic(err)
	}
}

func TestDataStorage_GetPlayerInfoBySteamID(t *testing.T) {
	tests := []struct {
		name    string
		steamID string
		want    steamclient.PlayerInfo
		wantErr bool
	}{
		{
			name:    "Retrieve PlayerInfo from fixtures (ID: playerInfo0 )",
			steamID: "playerInfo0",
			want: steamclient.PlayerInfo{
				PlayerSummary: steamclient.PlayerSummary{
					Avatar:                   "1",
					Avatarfull:               "2",
					Avatarmedium:             "3",
					Cityid:                   "4",
					Commentpermission:        "5",
					Communityvisibilitystate: "6",
					Gameextrainfo:            "7",
					Gameid:                   "8",
					Gameserverip:             "9",
					Lastlogoff:               "10",
					Loccityid:                "11",
					Loccountrycode:           "12",
					Locstatecode:             "13",
					Personaname:              "14",
					Personastate:             "15",
					Primaryclanid:            "16",
					Profilestate:             "17",
					Profileurl:               "18",
					Realname:                 "19",
					SteamID:                  "20",
					Timecreated:              "21",
				},
				UserStatsForGame: steamclient.UserStatsForGame{
					Archivements: steamclient.GameArchievements{},
					Extra: steamclient.GameExtras{
						HitRatio:     "22",
						LastMatchADR: "23",
						LastMatchKD:  "24",
						PlayedHours:  "25",
						TotalADR:     "26",
						TotalKD:      "27",
					},
					GameName: "28",
					Stats: steamclient.GameStats{
						// TODO Just testing some of the stats for now, maybe add more
						TotalWinsMapDeLake:       "29",
						TotalWinsMapDeNuke:       "29",
						TotalWinsMapDeSafehouse:  "29",
						TotalWinsMapDeShorttrain: "29",
						TotalWinsMapDeStmarc:     "29",
						TotalWinsMapDeSugarcane:  "29",
						TotalWinsMapDeTrain:      "29",
						TotalWinsMapDeVertigo:    "29",
						TotalWinsPistolround:     "29",
					},
					SteamID: "30",
				},
				RecentlyPlayedGames: steamclient.RecentlyPlayedGames{
					Appid:                  "31",
					ImgIconURL:             "32",
					ImgLogoURL:             "33",
					Name:                   "34",
					Playtime2Weeks:         "35",
					PlaytimeForever:        "36",
					PlaytimeLinuxForever:   "37",
					PlaytimeMacForever:     "38",
					PlaytimeWindowsForever: "39",
					SteamID:                "40",
				},
				PlayerHistory: steamclient.PlayerHistory{

					SteamID: "41",

					Data: []steamclient.PlayerHistoryEntry{
						{
							HitRatio:                   "42",
							LastMatchADR:               "43",
							LastMatchContributionScore: "44",
							LastMatchDamage:            "45",
							LastMatchDeaths:            "46",
							LastMatchKD:                "47",
							LastMatchKills:             "48",
							LastMatchRounds:            "49",
							Playtime2Weeks:             "50",
							Time:                       "51",
							TotalADR:                   "52",
							TotalKD:                    "53",
							TotalKills:                 "54",
							TotalKillsHeadshot:         "55",
							TotalShotsFired:            "56",
							TotalShotsHit:              "57",
						},
						{
							HitRatio:                   "58",
							LastMatchADR:               "59",
							LastMatchContributionScore: "60",
							LastMatchDamage:            "61",
							LastMatchDeaths:            "62",
							LastMatchKD:                "63",
							LastMatchKills:             "64",
							LastMatchRounds:            "65",
							Playtime2Weeks:             "66",
							Time:                       "67",
							TotalADR:                   "68",
							TotalKD:                    "69",
							TotalKills:                 "70",
							TotalKillsHeadshot:         "71",
							TotalShotsFired:            "72",
							TotalShotsHit:              "73",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prepareDB()
			got, err := db.GetPlayerInfoBySteamID(tt.steamID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.GetPlayerInfoBySteamID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("DataStorage.GetPlayerInfoBySteamID() mismatch (-want +got):\n%s", diff)
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("DataStorage.GetPlayerInfoBySteamID() = %v, want %v", got, tt.want)
			// }
		})
	}
}

// func TestNewDataStorage(t *testing.T) {
// 	type args struct {
// 		path string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *DataStorage
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := NewDataStorage(tt.args.path)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("NewDataStorage() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewDataStorage() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestDataStorage_GetAllPlayers(t *testing.T) {
	type fields struct {
		db         *sql.DB
		statements map[string]*sql.Stmt
	}
	tests := []struct {
		name    string
		fields  fields
		want    []steamclient.PlayerInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DataStorage{
				db:         tt.fields.db,
				statements: tt.fields.statements,
			}
			got, err := ds.GetAllPlayers()
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.GetAllPlayers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataStorage.GetAllPlayers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataStorage_UpdatePlayerInfo(t *testing.T) {
	type fields struct {
		db         *sql.DB
		statements map[string]*sql.Stmt
	}
	type args struct {
		pi steamclient.PlayerInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DataStorage{
				db:         tt.fields.db,
				statements: tt.fields.statements,
			}
			if err := ds.UpdatePlayerInfo(tt.args.pi); (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.UpdatePlayerInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
