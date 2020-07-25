package demoparser

import (
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

func TestMyParser_Parse(t *testing.T) {
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
				ID: "TODO",

				// TODO find out how to get the time when the match was played
				Time: time.Date(2020, time.July, 1, 3, 4, 5, 6, time.UTC),
				Map:  "de_mirage",

				// Match in the demo took 22 rounds
				// TODO fill in correct values
				Rounds: []Round{
					{
						Number:       1,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
						},
					},
					{
						Number:       2,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
						},
					},
					{
						Number:       3,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
						},
					},
					{
						Number:       4,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
						},
					},
					{
						Number:       5,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
						},
					},
					{
						Number:       6,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
						},
					},
					{
						Number:       7,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
						},
					},
					{
						Number:       8,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
						},
					},
					{
						Number:       9,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
						},
					},
					{
						Number:       10,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
						},
					},
					{
						Number:       11,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
						},
					},
					{
						Number:       12,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
						},
					},
					{
						Number:       13,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
						},
					},
					{
						Number:       14,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
						},
					},
					{
						Number:       15,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
						},
					},
					{
						Number:       16,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
						},
					},
					{
						Number:       17,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
						},
					},
					{
						Number:       18,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
						},
					},
					{
						Number:       19,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
						},
					},
					{
						Number:       20,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
						},
					},
					{
						Number:       21,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
						},
					},
					{
						Number:       22,
						TimeStart:    1,
						TimeEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players: []Player{
							{
								ID:          "TODO",
								Team:        "TODO",
								Won:         true,
								Mvp:         true,
								MvpReason:   "TODO",
								BombPlanted: true,
								BombDefused: true,
								ShotsFired:  1,
								ShotsHit:    1,
								Headshots:   1,
								Kills: []Kill{
									{
										Tick:         1,
										Assist:       true,
										VictimID:     "TODO",
										AssistID:     "TODO",
										AssistedID:   "TODO",
										WeaponUsedID: 1,
									},
								},
								Victims: []DamageVictim{
									{
										DamagedID:          "TODO",
										MostDamageWeaponID: 1,
										Amount:             1,
									},
								},
							},
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
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("MyParser.Parse() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
