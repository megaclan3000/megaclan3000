package main

import (
	// "encoding/binary"
	"encoding/json"
	"errors"

	// "strconv"

	// "github.com/megaclan3000/megaclan3000/internal/demoparser"
	"github.com/megaclan3000/megaclan3000/internal/steamclient"
	log "github.com/sirupsen/logrus"

	"context"
	"time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

// DataStorage is the interface to get and set data retrieved from the steam
// API. It holds the data as in-memory cache to avoid having to pull the data
// when a request is made for better response time
type DataStorage struct {
	Players []steamclient.PlayerInfo

	ctx               context.Context
	database          *mongo.Database
	matchesCollection *mongo.Collection
	killsCollection   *mongo.Collection
	roundsCollection  *mongo.Collection
	client            *mongo.Client
	cancel            context.CancelFunc
}

func (ds *DataStorage) Close() {
	defer ds.cancel()
	defer ds.client.Disconnect(ds.ctx)

}

func NewDataStorage(mongoURI, dbName string) (*DataStorage, error) {

	var err error
	ds := DataStorage{}

	// ds.ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	// ds.client, err = mongo.Connect(ds.ctx, options.Client().ApplyURI(mongoURI))

	ds.ctx, ds.cancel = context.WithTimeout(context.Background(), 10*time.Second)

	ds.client, err = mongo.Connect(ds.ctx, options.Client().ApplyURI(mongoURI))

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		return &ds, err
	}

	err = ds.client.Ping(ds.ctx, nil)

	if err != nil {
		return &ds, err
	}

	log.Debug("Ping to succeeded!", dbName)

	ds.database = ds.client.Database(dbName)
	ds.matchesCollection = ds.database.Collection("matches")
	ds.killsCollection = ds.database.Collection("kills")
	ds.roundsCollection = ds.database.Collection("rounds")

	return &ds, nil

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

type matchstat struct {
	Time         time.Time
	ScoreClan    int
	ScoreEnemy   int
	PlayersClan  []string
	PlayersEnemy []string
}

func (ms matchstat) MarshalJSON() ([]byte, error) {

	return json.Marshal(&struct {
		Time         time.Time `json:"time"`
		ScoreClan    int       `json:"score_clan"`
		ScoreEnemy   int       `json:"score_enemy"`
		PlayersClan  []string  `json:"players_clan"`
		PlayersEnemy []string  `json:"players_enemy"`
	}{

		Time:         ms.Time,
		ScoreClan:    ms.ScoreClan,
		ScoreEnemy:   ms.ScoreEnemy,
		PlayersClan:  ms.PlayersClan,
		PlayersEnemy: ms.PlayersEnemy,
	})

}

func (ds DataStorage) GetMatches() []matchstat {

	return []matchstat{
		{
			Time:         time.Now(),
			ScoreClan:    1,
			ScoreEnemy:   2,
			PlayersClan:  []string{"player1", "player2"},
			PlayersEnemy: []string{"player3", "player4"},
		},
		{
			Time:         time.Now(),
			ScoreClan:    1,
			ScoreEnemy:   2,
			PlayersClan:  []string{"player1", "player2"},
			PlayersEnemy: []string{"player3", "player4"},
		},

		{
			Time:         time.Now(),
			ScoreClan:    1,
			ScoreEnemy:   2,
			PlayersClan:  []string{"player1", "player2"},
			PlayersEnemy: []string{"player3", "player4"},
		},
	}
}

//func (ds *DataStorage) SaveMatch(match demoparser.Match) error {

//	// , err := primitive.ObjectIDFromHex(strconv.FormatUint(match.Hash, 10))

//	// bMatch := demoparser.BsonMatch{
//	// 	ObjectID:   primitive.NewObjectID(),
//	// 	UploadTime: match.UploadTime,
//	// 	Map:        match.Map,
//	// }

//	// var roundPlayerIDs []primitive.ObjectID
//	// var matchPlayerIDs []primitive.ObjectID
//	// // Save all players using steamID64 as objectID
//	// for _, r := range match.Rounds {
//	// 	for _, rp := range r.Players {
//	// 		if insertResult, err := ds.playersCollection.InsertOne(ds.ctx, demoparser.NewBsonPlayer(rp)); err != nil {
//	// 			return err
//	// 		} else {
//	// 			log.Debug("Inserted player: ", insertResult.InsertedID)
//	// 			roundPlayerIDs = append(roundPlayerIDs, insertResult.InsertedID.(primitive.ObjectID))
//	// 			matchPlayerIDs = append(matchPlayerIDs, insertResult.InsertedID.(primitive.ObjectID))
//	// 		}
//	// 	}
//	// }

//	// Save the rounds
//	// for _, r := range match.Rounds {
//	// 	var roundKills []primitive.ObjectID

//	// // Save kills of the round
//	// for _, rk := range r.Kills {
//	// 	bkill := demoparser.NewBsonKill(rk)
//	// 	if insertResult, err := ds.killsCollection.InsertOne(ds.ctx, bkill); err != nil {
//	// 		return err
//	// 	} else {
//	// 		log.Debug("Inserted kill: ", insertResult.InsertedID)
//	// 		roundKills = append(roundKills, insertResult.InsertedID.(primitive.ObjectID))
//	// 	}
//	// }

//	// bRound := demoparser.BsonRound{

//	// 	ObjectID:     primitive.NewObjectID(),
//	// 	TimeStart:    r.TimeStart,
//	// 	TimeEnd:      r.TimeEnd,
//	// 	TeamWon:      r.TeamWon,
//	// 	BombPlanted:  r.BombPlanted,
//	// 	BombDefused:  r.BombDefused,
//	// 	BombExploded: r.BombExploded,
//	// 	// Players:      roundPlayerIDs,
//	// 	Kills: roundKills,
//	// }

//	// }

//	var rounds []primitive.ObjectID
//	players := make(map[uint64]uint64)

//	for _, r := range match.Rounds {

//		if insertResult, err := ds.roundsCollection.InsertOne(ds.ctx, demoparser.NewBsonRound(*r)); err != nil {
//			return err
//		} else {
//			log.Debug("Inserted round: ", insertResult.InsertedID)
//			rounds = append(rounds, insertResult.InsertedID.(primitive.ObjectID))
//			for _, p := range r.Players {
//				players[p.SteamID64] = p.SteamID64
//			}
//		}
//	}

//	uniquePlayers := make([]uint64, 0, len(players))

//	for k := range players {
//		uniquePlayers = append(uniquePlayers, k)
//	}

//	// bMatch.Players = matchPlayerIDs
//	bMatch := demoparser.BsonMatch{
//		ObjectID:   primitive.NewObjectID(),
//		UploadTime: match.UploadTime,
//		Map:        match.Map,
//		Players:    uniquePlayers,
//		Rounds:     rounds,
//	}

//	if insertResult, err := ds.matchesCollection.InsertOne(ds.ctx, bMatch); err != nil {
//		return err
//	} else {
//		log.Debug("Inserted match: ", insertResult.InsertedID)
//	}

//	return nil
//}

//func (ds *DataStorage) GetPlayerObjectBySteamID(steamID uint64) (primitive.ObjectID, error) {
//	//TODO implement
//	return primitive.NewObjectID(), nil
//}

//func objectIDFromUint64(i uint64) primitive.ObjectID {
//	var ID primitive.ObjectID
//	binary.LittleEndian.PutUint64(ID[:], uint64(i))
//	return ID
//}
