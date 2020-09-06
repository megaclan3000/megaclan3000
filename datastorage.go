package main

import (
	"errors"

	// "github.com/megaclan3000/megaclan3000/internal/demoparser"
	"github.com/megaclan3000/megaclan3000/internal/demoparser"
	"github.com/megaclan3000/megaclan3000/internal/steamclient"
)

// DataStorage is the interface to get and set data retrieved from the steam
// API. It holds the data as in-memory cache to avoid having to pull the data
// when a request is made for better response time
type DataStorage struct {
	Players []steamclient.PlayerInfo
}

func (ds *DataStorage) Upload(match demoparser.InfoStruct) error {
	return nil
}

func (ds *DataStorage) GetMatchByID(id string) (demoparser.InfoStruct, error) {

	demoInfoFromDem, err := demoparser.GetMatchInfo("1", steamClient)
	if err != nil {
		panic(err)
	}

	return demoInfoFromDem, nil
}

// GetPlayerInfoBySteamID returns the PlayerInfo object for a given steamID
func (ds DataStorage) GetPlayerInfoBySteamID(steamID uint64) (steamclient.PlayerInfo, error) {

	for _, v := range ds.Players {
		if v.PlayerSummary.SteamID == steamID {
			return v, nil
		}
	}
	return steamclient.PlayerInfo{}, errors.New("Player not found")
}
