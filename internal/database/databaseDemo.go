package database

import (
	"github.com/megaclan3000/megaclan3000/internal/demoparser"
)

func (ds *DataStorage) InsertDemoScores(scores []demoparser.Score) error {
	return ds.dbm.Insert(&scores)
}
