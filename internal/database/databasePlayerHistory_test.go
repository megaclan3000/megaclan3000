package database

import (
	"reflect"
	"testing"

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
			name:    "Retrieve PlayerHistory from fixtures",
			steamID: "123456789",
			want: steamclient.PlayerHistory{
				SteamID:    123456789,
				Time:       1,
				TotalKills: 2,
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
				t.Errorf("DataStorage.GetPlayerHistory() = %v, want %v", got, tt.want)
			}
		})
	}
}
