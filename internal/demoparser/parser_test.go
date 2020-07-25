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
				ID: "Test",

				// TODO find out how to get the time when the match was played
				Time: time.Date(2020, time.July, 1, 3, 4, 5, 6, time.UTC),
				Map:  "de_mirage",

				// Match in the demo took 22 rounds
				// TODO fill in correct values
				Rounds: []Round{
					{
						Number:       1,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
					},
					{
						Number:       2,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
					},

					{
						Number:       3,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
					},

					{
						Number:       4,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
					},

					{
						Number:       5,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
					},

					{
						Number:       6,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
					},

					{
						Number:       7,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
					},

					{
						Number:       8,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
					},

					{
						Number:       9,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
					},

					{
						Number:       10,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
					},

					{
						Number:       11,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
					},

					{
						Number:       12,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
					},

					{
						Number:       13,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
					},

					{
						Number:       14,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
					},

					{
						Number:       15,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
					},

					{
						Number:       16,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
					},

					{
						Number:       17,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
					},

					{
						Number:       18,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
					},

					{
						Number:       19,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
					},

					{
						Number:       20,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
					},

					{
						Number:       21,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
					},

					{
						Number:       22,
						TickStart:    1,
						TickEnd:      1,
						TeamWon:      "CT",
						BombPlanted:  true,
						BombDefused:  true,
						BombExploded: true,
						Players:      []Player{},
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
