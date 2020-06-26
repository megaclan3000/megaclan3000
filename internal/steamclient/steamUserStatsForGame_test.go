package steamclient

import (
	"reflect"
	"testing"
)

func TestSteamClient_ParseUserStatsForGame(t *testing.T) {
	type fields struct {
		Config SteamConfig
	}
	type args struct {
		data userStatsForGameData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    UserStatsForGame
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &SteamClient{
				Config: tt.fields.Config,
			}
			got, err := sc.ParseUserStatsForGame(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("SteamClient.ParseUserStatsForGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SteamClient.ParseUserStatsForGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nilToZeroString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "String with number 123",
			input: "123",
			want:  "123",
		},
		{
			name:  "String with number 0",
			input: "0",
			want:  "0",
		},
		{
			name:  "Empty String",
			input: "",
			want:  "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nilToZeroString(tt.input); got != tt.want {
				t.Errorf("nilToZeroString() = %v, want %v", got, tt.want)
			}
		})
	}
}
