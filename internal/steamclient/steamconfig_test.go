package steamclient

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestNewSteamConfig(t *testing.T) {
	tests := []struct {
		name       string
		configPath string
		want       SteamConfig
		wantErr    bool
	}{
		{
			name:       "Create config from existing path",
			configPath: "../../test/steamclient/testconfig.json",
			want: SteamConfig{

				SteamAPIKey:     "fakekey",
				UpdateInterval:  4,
				HistoryInterval: 20,
				SteamIDs: []string{
					"76561197962156894",
					"76561197967611281",
					"76561198217140904",
					"76561198962966497",
					"76561198881047143",
					"76561197978562286",
					"76561197978015984",
					"76561198092006615",
					"76561198104947907",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSteamConfig(tt.configPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSteamConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("NewSteamConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
