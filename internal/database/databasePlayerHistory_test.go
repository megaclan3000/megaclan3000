package database

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/pinpox/megaclan3000/internal/steamclient"
)

func TestDataStorage_GetPlayerHistory(t *testing.T) {
	tests := []struct {
		name    string
		steamID string
		want    steamclient.PlayerHistory
		wantErr bool
	}{
		{
			name:    "Retrieve PlayerHistory from fixtures (ID: 1)",
			steamID: "1",
			want: steamclient.PlayerHistory{
				SteamID: "1",
				Data: []steamclient.PlayerHistoryEntry{
					{
						Time:                       "100",
						TotalKills:                 "101",
						TotalADR:                   "102",
						TotalShotsHit:              "103",
						TotalShotsFired:            "104",
						TotalKillsHeadshot:         "105",
						TotalKD:                    "106",
						LastMatchContributionScore: "107",
						LastMatchDamage:            "108",
						LastMatchDeaths:            "109",
						LastMatchKills:             "1010",
						LastMatchRounds:            "1011",
						LastMatchKD:                "1012",
						LastMatchADR:               "1013",
						HitRatio:                   "1014",
						Playtime2Weeks:             "1015",
					},
				},
			},
			wantErr: false,
		},

		{
			name:    "Retrieve PlayerHistory from fixtures (ID: 2)",
			steamID: "2",
			want: steamclient.PlayerHistory{
				SteamID: "2",
				Data: []steamclient.PlayerHistoryEntry{
					{
						Time:                       "200",
						TotalKills:                 "201",
						TotalADR:                   "202",
						TotalShotsHit:              "203",
						TotalShotsFired:            "204",
						TotalKillsHeadshot:         "205",
						TotalKD:                    "206",
						LastMatchContributionScore: "207",
						LastMatchDamage:            "208",
						LastMatchDeaths:            "209",
						LastMatchKills:             "2010",
						LastMatchRounds:            "2011",
						LastMatchKD:                "2012",
						LastMatchADR:               "2013",
						HitRatio:                   "2014",
						Playtime2Weeks:             "2015",
					},
				},
			},
			wantErr: false,
		},

		{
			name:    "Retrieve PlayerHistory from fixtures (ID: 3)",
			steamID: "3",
			want: steamclient.PlayerHistory{
				SteamID: "3",
				Data: []steamclient.PlayerHistoryEntry{
					{
						Time:                       "300",
						TotalKills:                 "301",
						TotalADR:                   "302",
						TotalShotsHit:              "303",
						TotalShotsFired:            "304",
						TotalKillsHeadshot:         "305",
						TotalKD:                    "306",
						LastMatchContributionScore: "307",
						LastMatchDamage:            "308",
						LastMatchDeaths:            "309",
						LastMatchKills:             "3010",
						LastMatchRounds:            "3011",
						LastMatchKD:                "3012",
						LastMatchADR:               "3013",
						HitRatio:                   "3014",
						Playtime2Weeks:             "3015",
					},
					{
						Time:                       "3003",
						TotalKills:                 "3013",
						TotalADR:                   "3023",
						TotalShotsHit:              "3033",
						TotalShotsFired:            "3043",
						TotalKillsHeadshot:         "3053",
						TotalKD:                    "3063",
						LastMatchContributionScore: "3073",
						LastMatchDamage:            "3083",
						LastMatchDeaths:            "3093",
						LastMatchKills:             "30103",
						LastMatchRounds:            "30113",
						LastMatchKD:                "30123",
						LastMatchADR:               "30133",
						HitRatio:                   "30143",
						Playtime2Weeks:             "30153",
					},
					{
						Time:                       "30034",
						TotalKills:                 "30134",
						TotalADR:                   "30234",
						TotalShotsHit:              "30334",
						TotalShotsFired:            "30434",
						TotalKillsHeadshot:         "30534",
						TotalKD:                    "30634",
						LastMatchContributionScore: "30734",
						LastMatchDamage:            "30834",
						LastMatchDeaths:            "30934",
						LastMatchKills:             "301034",
						LastMatchRounds:            "301134",
						LastMatchKD:                "301234",
						LastMatchADR:               "301334",
						HitRatio:                   "301434",
						Playtime2Weeks:             "301534",
					},
				},
			},
			wantErr: false,
		},

		{
			name:    "Retrieve PlayerHistory from fixtures (ID: 4)",
			steamID: "4",
			want: steamclient.PlayerHistory{
				SteamID: "4",
				Data: []steamclient.PlayerHistoryEntry{
					{
						Time:                       "400",
						TotalKills:                 "401",
						TotalADR:                   "402",
						TotalShotsHit:              "403",
						TotalShotsFired:            "404",
						TotalKillsHeadshot:         "405",
						TotalKD:                    "406",
						LastMatchContributionScore: "407",
						LastMatchDamage:            "408",
						LastMatchDeaths:            "409",
						LastMatchKills:             "4010",
						LastMatchRounds:            "4011",
						LastMatchKD:                "4012",
						LastMatchADR:               "4013",
						HitRatio:                   "4014",
						Playtime2Weeks:             "4015",
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			prepareDB()
			got, err := db.GetPlayerHistory(tt.steamID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.GetPlayerHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataStorage.GetPlayerHistory() = \n%v \nwant:\n%v", spew.Sdump(got), spew.Sdump(tt.want))
			}
		})
	}
}

func TestDataStorage_UpdatePlayerHistory(t *testing.T) {
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
			if err := ds.UpdatePlayerHistory(tt.args.pi); (err != nil) != tt.wantErr {
				t.Errorf("DataStorage.UpdatePlayerHistory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
