package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pinpox/megaclan3000/internal/steamclient"
)

// Public data query methods

// GetPlayerSummary returns a PlayerSummary object by fetching the values from
// the database using a prepared statement.
func (ds *DataStorage) GetPlayerSummary(steamID string) (steamclient.PlayerSummary, error) {

	ps := steamclient.PlayerSummary{}
	var err error

	if rows, err := ds.statements["select_player_summary"].Query(); err == nil {
		rows.Scan(
			&ps.SteamID,
			&ps.Communityvisibilitystate,
			&ps.Profilestate,
			&ps.Personaname,
			&ps.Profileurl,
			&ps.Avatar,
			&ps.Avatarmedium,
			&ps.Avatarfull,
			&ps.Lastlogoff,
			&ps.Personastate,
			&ps.Primaryclanid,
			&ps.Timecreated,
		)
	}
	return ps, err
}

func (ds *DataStorage) GetAllPlayers() ([]steamclient.PlayerInfo, error) {
	var players []steamclient.PlayerInfo
	var rows *sql.Rows
	var err error

	if rows, err = ds.statements["select_all_player_ids"].Query(); err != nil {
		return players, err
	}

	var steamID string

	for rows.Next() {
		if err = rows.Scan(&steamID); err == nil {
			if pi, err := ds.GetPlayerInfoBySteamID(steamID); err == nil {
				players = append(players, pi)
			}
		}
	}

	rows.Close() //good habit to close
	return players, nil
}

// GetUserStatsForGame returns a UserStatsForGame object by fetching the values from
// the database using a prepared statement.
func (ds *DataStorage) GetUserStatsForGame(steamID string) (steamclient.UserStatsForGame, error) {

	usfg := steamclient.UserStatsForGame{}
	var err error

	if rows, err := ds.statements["select_player_stats"].Query(steamID); err == nil {
		rows.Scan(

			&usfg.SteamID,
			&usfg.Stats.TotalKills,
			&usfg.Stats.TotalDeaths,
			&usfg.Stats.TotalTimePlayed,
			&usfg.Stats.TotalPlantedBombs,
			&usfg.Stats.TotalDefusedBombs,
			&usfg.Stats.TotalWins,
			&usfg.Stats.TotalDamageDone,
			&usfg.Stats.TotalMoneyEarned,
			&usfg.Stats.TotalKillsKnife,
			&usfg.Stats.TotalKillsHegrenade,
			&usfg.Stats.TotalKillsGlock,
			&usfg.Stats.TotalKillsDeagle,
			&usfg.Stats.TotalKillsElite,
			&usfg.Stats.TotalKillsFiveseven,
			&usfg.Stats.TotalKillsXm1014,
			&usfg.Stats.TotalKillsMac10,
			&usfg.Stats.TotalKillsUmp45,
			&usfg.Stats.TotalKillsP90,
			&usfg.Stats.TotalKillsAwp,
			&usfg.Stats.TotalKillsAk47,
			&usfg.Stats.TotalKillsAug,
			&usfg.Stats.TotalKillsFamas,
			&usfg.Stats.TotalKillsG3sg1,
			&usfg.Stats.TotalKillsM249,
			&usfg.Stats.TotalKillsHeadshot,
			&usfg.Stats.TotalKillsEnemyWeapon,
			&usfg.Stats.TotalWinsPistolround,
			&usfg.Stats.TotalWinsMapCsAssault,
			&usfg.Stats.TotalWinsMapDeDust2,
			&usfg.Stats.TotalWinsMapDeInferno,
			&usfg.Stats.TotalWinsMapDeTrain,
			&usfg.Stats.TotalWeaponsDonated,
			&usfg.Stats.TotalKillsEnemyBlinded,
			&usfg.Stats.TotalKillsKnifeFight,
			&usfg.Stats.TotalKillsAgainstZoomedSniper,
			&usfg.Stats.TotalDominations,
			&usfg.Stats.TotalDominationOverkills,
			&usfg.Stats.TotalRevenges,
			&usfg.Stats.TotalShotsHit,
			&usfg.Stats.TotalShotsFired,
			&usfg.Stats.TotalRoundsPlayed,
			&usfg.Stats.TotalShotsDeagle,
			&usfg.Stats.TotalShotsGlock,
			&usfg.Stats.TotalShotsElite,
			&usfg.Stats.TotalShotsFiveseven,
			&usfg.Stats.TotalShotsAwp,
			&usfg.Stats.TotalShotsAk47,
			&usfg.Stats.TotalShotsAug,
			&usfg.Stats.TotalShotsFamas,
			&usfg.Stats.TotalShotsG3sg1,
			&usfg.Stats.TotalShotsP90,
			&usfg.Stats.TotalShotsMac10,
			&usfg.Stats.TotalShotsUmp45,
			&usfg.Stats.TotalShotsXm1014,
			&usfg.Stats.TotalShotsM249,
			&usfg.Stats.TotalHitsDeagle,
			&usfg.Stats.TotalHitsGlock,
			&usfg.Stats.TotalHitsElite,
			&usfg.Stats.TotalHitsFiveseven,
			&usfg.Stats.TotalHitsAwp,
			&usfg.Stats.TotalHitsAk47,
			&usfg.Stats.TotalHitsAug,
			&usfg.Stats.TotalHitsFamas,
			&usfg.Stats.TotalHitsG3sg1,
			&usfg.Stats.TotalHitsP90,
			&usfg.Stats.TotalHitsMac10,
			&usfg.Stats.TotalHitsUmp45,
			&usfg.Stats.TotalHitsXm1014,
			&usfg.Stats.TotalHitsM249,
			&usfg.Stats.TotalRoundsMapCsAssault,
			&usfg.Stats.TotalRoundsMapDeDust2,
			&usfg.Stats.TotalRoundsMapDeInferno,
			&usfg.Stats.TotalRoundsMapDeTrain,
			&usfg.Stats.LastMatchTWins,
			&usfg.Stats.LastMatchCtWins,
			&usfg.Stats.LastMatchWins,
			&usfg.Stats.LastMatchMaxPlayers,
			&usfg.Stats.LastMatchKills,
			&usfg.Stats.LastMatchDeaths,
			&usfg.Stats.LastMatchMvps,
			&usfg.Stats.LastMatchFavweaponID,
			&usfg.Stats.LastMatchFavweaponShots,
			&usfg.Stats.LastMatchFavweaponHits,
			&usfg.Stats.LastMatchFavweaponKills,
			&usfg.Stats.LastMatchDamage,
			&usfg.Stats.LastMatchMoneySpent,
			&usfg.Stats.LastMatchDominations,
			&usfg.Stats.LastMatchRevenges,
			&usfg.Stats.TotalMvps,
			&usfg.Stats.TotalRoundsMapDeLake,
			&usfg.Stats.TotalRoundsMapDeSafehouse,
			&usfg.Stats.TotalRoundsMapDeBank,
			&usfg.Stats.TotalTRPlantedBombs,
			&usfg.Stats.TotalGunGameRoundsWon,
			&usfg.Stats.TotalGunGameRoundsPlayed,
			&usfg.Stats.TotalWinsMapDeBank,
			&usfg.Stats.TotalWinsMapDeLake,
			&usfg.Stats.TotalMatchesWonBank,
			&usfg.Stats.TotalMatchesWon,
			&usfg.Stats.TotalMatchesPlayed,
			&usfg.Stats.TotalGgMatchesWon,
			&usfg.Stats.TotalGgMatchesPlayed,
			&usfg.Stats.TotalProgressiveMatchesWon,
			&usfg.Stats.TotalTrbombMatchesWon,
			&usfg.Stats.TotalContributionScore,
			&usfg.Stats.LastMatchContributionScore,
			&usfg.Stats.LastMatchRounds,
			&usfg.Stats.TotalKillsHkp2000,
			&usfg.Stats.TotalShotsHkp2000,
			&usfg.Stats.TotalHitsHkp2000,
			&usfg.Stats.TotalHitsP250,
			&usfg.Stats.TotalKillsP250,
			&usfg.Stats.TotalShotsP250,
			&usfg.Stats.TotalKillsSg556,
			&usfg.Stats.TotalShotsSg556,
			&usfg.Stats.TotalHitsSg556,
			&usfg.Stats.TotalHitsScar20,
			&usfg.Stats.TotalKillsScar20,
			&usfg.Stats.TotalShotsScar20,
			&usfg.Stats.TotalShotsSsg08,
			&usfg.Stats.TotalHitsSsg08,
			&usfg.Stats.TotalKillsSsg08,
			&usfg.Stats.TotalShotsMp7,
			&usfg.Stats.TotalHitsMp7,
			&usfg.Stats.TotalKillsMp7,
			&usfg.Stats.TotalKillsMp9,
			&usfg.Stats.TotalShotsMp9,
			&usfg.Stats.TotalHitsMp9,
			&usfg.Stats.TotalHitsNova,
			&usfg.Stats.TotalKillsNova,
			&usfg.Stats.TotalShotsNova,
			&usfg.Stats.TotalHitsNegev,
			&usfg.Stats.TotalKillsNegev,
			&usfg.Stats.TotalShotsNegev,
			&usfg.Stats.TotalShotsSawedoff,
			&usfg.Stats.TotalHitsSawedoff,
			&usfg.Stats.TotalKillsSawedoff,
			&usfg.Stats.TotalShotsBizon,
			&usfg.Stats.TotalHitsBizon,
			&usfg.Stats.TotalKillsBizon,
			&usfg.Stats.TotalKillsTec9,
			&usfg.Stats.TotalShotsTec9,
			&usfg.Stats.TotalHitsTec9,
			&usfg.Stats.TotalShotsMag7,
			&usfg.Stats.TotalHitsMag7,
			&usfg.Stats.TotalKillsMag7,
			&usfg.Stats.TotalGunGameContributionScore,
			&usfg.Stats.LastMatchGgContributionScore,
			&usfg.Stats.TotalKillsM4a1,
			&usfg.Stats.TotalKillsGalilar,
			&usfg.Stats.TotalKillsMolotov,
			&usfg.Stats.TotalKillsTaser,
			&usfg.Stats.TotalShotsM4a1,
			&usfg.Stats.TotalShotsGalilar,
			&usfg.Stats.TotalShotsTaser,
			&usfg.Stats.TotalHitsM4a1,
			&usfg.Stats.TotalHitsGalilar,
			&usfg.Stats.TotalMatchesWonTrain,
			&usfg.Stats.TotalMatchesWonLake,
			&usfg.Stats.GILessonCsgoInstrExplainBuymenu,
			&usfg.Stats.GILessonCsgoInstrExplainBuyarmor,
			&usfg.Stats.GILessonCsgoInstrExplainPlantBomb,
			&usfg.Stats.GILessonCsgoInstrExplainBombCarrier,
			&usfg.Stats.GILessonBombSitesA,
			&usfg.Stats.GILessonDefusePlantedBomb,
			&usfg.Stats.GILessonCsgoInstrExplainFollowBomber,
			&usfg.Stats.GILessonCsgoInstrExplainPickupBomb,
			&usfg.Stats.GILessonCsgoInstrExplainPreventBombPickup,
			&usfg.Stats.GILessonCsgoCycleWeaponsKb,
			&usfg.Stats.GILessonCsgoInstrExplainZoom,
			&usfg.Stats.GILessonCsgoInstrExplainReload,
			&usfg.Stats.GILessonTrExplainPlantBomb,
			&usfg.Stats.GILessonBombSitesB,
			&usfg.Stats.GILessonVersionNumber,
			&usfg.Stats.GILessonFindPlantedBomb,
			&usfg.Stats.GILessonCsgoHostageLeadToHrz,
			&usfg.Stats.GILessonCsgoInstrRescueZone,
			&usfg.Stats.GILessonCsgoInstrExplainInspect,
			&usfg.Stats.SteamStatXpearnedgames,
		)
	}
	return usfg, err
}

// GetRecentlyPlayedGames returns a RecentlyPlayedGames object by fetching the values from
// the database using a prepared statement.
func (ds *DataStorage) GetRecentlyPlayedGames(steamID string) (steamclient.RecentlyPlayedGames, error) {
	rpg := steamclient.RecentlyPlayedGames{}
	var err error
	var id int

	if rows, err := ds.statements["select_recently_played"].Query(steamID); err == nil {
		rows.Scan(
			&id,
			&rpg.Playtime2Weeks,
			&rpg.PlaytimeForever,
			&rpg.PlaytimeWindowsForever,
			&rpg.PlaytimeMacForever,
			&rpg.PlaytimeLinuxForever,
		)
	}
	return rpg, err
}

// GetPlayerHistory returns a PlayerHistory object by fetching the values from
// the database using a prepared statement.
func (ds *DataStorage) GetPlayerHistory(steamID string) (steamclient.PlayerHistory, error) {
	ph := steamclient.PlayerHistory{}
	var err error

	if rows, err := ds.statements["select_player_history"].Query(steamID); err == nil {
		rows.Scan(
			&ph.SteamID,
			&ph.Time,
			&ph.TotalKills,
		)
	}
	return ph, err
}

// Private data retrieval methods
func (ds *DataStorage) UpdatePlayerSummary(ps steamclient.PlayerSummary) {

	var result sql.Result
	var err error

	if result, err = ds.statements["update_player_summary"].Exec(
		ps.Communityvisibilitystate,
		ps.Profilestate,
		ps.Personaname,
		ps.Profileurl,
		ps.Avatar,
		ps.Avatarmedium,
		ps.Avatarfull,
		ps.Lastlogoff,
		ps.Personastate,
		ps.Primaryclanid,
		ps.Timecreated,
		ps.SteamID,
	); err != nil {
		log.Fatal(err)
	}

	rows, err := result.RowsAffected()
	log.Println("Rows affected:", rows)
}

func (ds *DataStorage) UpdateUserStatsForGame(stats steamclient.UserStatsForGame) {

	var result sql.Result
	var err error

	if result, err = ds.statements["update_player_stats"].Exec(
		stats.SteamID,
		stats.Stats.TotalKills,
		stats.Stats.TotalDeaths,
		stats.Stats.TotalTimePlayed,
		stats.Stats.TotalPlantedBombs,
		stats.Stats.TotalDefusedBombs,
		stats.Stats.TotalWins,
		stats.Stats.TotalDamageDone,
		stats.Stats.TotalMoneyEarned,
		stats.Stats.TotalKillsKnife,
		stats.Stats.TotalKillsHegrenade,
		stats.Stats.TotalKillsGlock,
		stats.Stats.TotalKillsDeagle,
		stats.Stats.TotalKillsElite,
		stats.Stats.TotalKillsFiveseven,
		stats.Stats.TotalKillsXm1014,
		stats.Stats.TotalKillsMac10,
		stats.Stats.TotalKillsUmp45,
		stats.Stats.TotalKillsP90,
		stats.Stats.TotalKillsAwp,
		stats.Stats.TotalKillsAk47,
		stats.Stats.TotalKillsAug,
		stats.Stats.TotalKillsFamas,
		stats.Stats.TotalKillsG3sg1,
		stats.Stats.TotalKillsM249,
		stats.Stats.TotalKillsHeadshot,
		stats.Stats.TotalKillsEnemyWeapon,
		stats.Stats.TotalWinsPistolround,
		stats.Stats.TotalWinsMapCsAssault,
		stats.Stats.TotalWinsMapDeDust2,
		stats.Stats.TotalWinsMapDeInferno,
		stats.Stats.TotalWinsMapDeTrain,
		stats.Stats.TotalWeaponsDonated,
		stats.Stats.TotalKillsEnemyBlinded,
		stats.Stats.TotalKillsKnifeFight,
		stats.Stats.TotalKillsAgainstZoomedSniper,
		stats.Stats.TotalDominations,
		stats.Stats.TotalDominationOverkills,
		stats.Stats.TotalRevenges,
		stats.Stats.TotalShotsHit,
		stats.Stats.TotalShotsFired,
		stats.Stats.TotalRoundsPlayed,
		stats.Stats.TotalShotsDeagle,
		stats.Stats.TotalShotsGlock,
		stats.Stats.TotalShotsElite,
		stats.Stats.TotalShotsFiveseven,
		stats.Stats.TotalShotsAwp,
		stats.Stats.TotalShotsAk47,
		stats.Stats.TotalShotsAug,
		stats.Stats.TotalShotsFamas,
		stats.Stats.TotalShotsG3sg1,
		stats.Stats.TotalShotsP90,
		stats.Stats.TotalShotsMac10,
		stats.Stats.TotalShotsUmp45,
		stats.Stats.TotalShotsXm1014,
		stats.Stats.TotalShotsM249,
		stats.Stats.TotalHitsDeagle,
		stats.Stats.TotalHitsGlock,
		stats.Stats.TotalHitsElite,
		stats.Stats.TotalHitsFiveseven,
		stats.Stats.TotalHitsAwp,
		stats.Stats.TotalHitsAk47,
		stats.Stats.TotalHitsAug,
		stats.Stats.TotalHitsFamas,
		stats.Stats.TotalHitsG3sg1,
		stats.Stats.TotalHitsP90,
		stats.Stats.TotalHitsMac10,
		stats.Stats.TotalHitsUmp45,
		stats.Stats.TotalHitsXm1014,
		stats.Stats.TotalHitsM249,
		stats.Stats.TotalRoundsMapCsAssault,
		stats.Stats.TotalRoundsMapDeDust2,
		stats.Stats.TotalRoundsMapDeInferno,
		stats.Stats.TotalRoundsMapDeTrain,
		stats.Stats.LastMatchTWins,
		stats.Stats.LastMatchCtWins,
		stats.Stats.LastMatchWins,
		stats.Stats.LastMatchMaxPlayers,
		stats.Stats.LastMatchKills,
		stats.Stats.LastMatchDeaths,
		stats.Stats.LastMatchMvps,
		stats.Stats.LastMatchFavweaponID,
		stats.Stats.LastMatchFavweaponShots,
		stats.Stats.LastMatchFavweaponHits,
		stats.Stats.LastMatchFavweaponKills,
		stats.Stats.LastMatchDamage,
		stats.Stats.LastMatchMoneySpent,
		stats.Stats.LastMatchDominations,
		stats.Stats.LastMatchRevenges,
		stats.Stats.TotalMvps,
		stats.Stats.TotalRoundsMapDeLake,
		stats.Stats.TotalRoundsMapDeSafehouse,
		stats.Stats.TotalRoundsMapDeBank,
		stats.Stats.TotalTRPlantedBombs,
		stats.Stats.TotalGunGameRoundsWon,
		stats.Stats.TotalGunGameRoundsPlayed,
		stats.Stats.TotalWinsMapDeBank,
		stats.Stats.TotalWinsMapDeLake,
		stats.Stats.TotalMatchesWonBank,
		stats.Stats.TotalMatchesWon,
		stats.Stats.TotalMatchesPlayed,
		stats.Stats.TotalGgMatchesWon,
		stats.Stats.TotalGgMatchesPlayed,
		stats.Stats.TotalProgressiveMatchesWon,
		stats.Stats.TotalTrbombMatchesWon,
		stats.Stats.TotalContributionScore,
		stats.Stats.LastMatchContributionScore,
		stats.Stats.LastMatchRounds,
		stats.Stats.TotalKillsHkp2000,
		stats.Stats.TotalShotsHkp2000,
		stats.Stats.TotalHitsHkp2000,
		stats.Stats.TotalHitsP250,
		stats.Stats.TotalKillsP250,
		stats.Stats.TotalShotsP250,
		stats.Stats.TotalKillsSg556,
		stats.Stats.TotalShotsSg556,
		stats.Stats.TotalHitsSg556,
		stats.Stats.TotalHitsScar20,
		stats.Stats.TotalKillsScar20,
		stats.Stats.TotalShotsScar20,
		stats.Stats.TotalShotsSsg08,
		stats.Stats.TotalHitsSsg08,
		stats.Stats.TotalKillsSsg08,
		stats.Stats.TotalShotsMp7,
		stats.Stats.TotalHitsMp7,
		stats.Stats.TotalKillsMp7,
		stats.Stats.TotalKillsMp9,
		stats.Stats.TotalShotsMp9,
		stats.Stats.TotalHitsMp9,
		stats.Stats.TotalHitsNova,
		stats.Stats.TotalKillsNova,
		stats.Stats.TotalShotsNova,
		stats.Stats.TotalHitsNegev,
		stats.Stats.TotalKillsNegev,
		stats.Stats.TotalShotsNegev,
		stats.Stats.TotalShotsSawedoff,
		stats.Stats.TotalHitsSawedoff,
		stats.Stats.TotalKillsSawedoff,
		stats.Stats.TotalShotsBizon,
		stats.Stats.TotalHitsBizon,
		stats.Stats.TotalKillsBizon,
		stats.Stats.TotalKillsTec9,
		stats.Stats.TotalShotsTec9,
		stats.Stats.TotalHitsTec9,
		stats.Stats.TotalShotsMag7,
		stats.Stats.TotalHitsMag7,
		stats.Stats.TotalKillsMag7,
		stats.Stats.TotalGunGameContributionScore,
		stats.Stats.LastMatchGgContributionScore,
		stats.Stats.TotalKillsM4a1,
		stats.Stats.TotalKillsGalilar,
		stats.Stats.TotalKillsMolotov,
		stats.Stats.TotalKillsTaser,
		stats.Stats.TotalShotsM4a1,
		stats.Stats.TotalShotsGalilar,
		stats.Stats.TotalShotsTaser,
		stats.Stats.TotalHitsM4a1,
		stats.Stats.TotalHitsGalilar,
		stats.Stats.TotalMatchesWonTrain,
		stats.Stats.TotalMatchesWonLake,
		stats.Stats.GILessonCsgoInstrExplainBuymenu,
		stats.Stats.GILessonCsgoInstrExplainBuyarmor,
		stats.Stats.GILessonCsgoInstrExplainPlantBomb,
		stats.Stats.GILessonCsgoInstrExplainBombCarrier,
		stats.Stats.GILessonBombSitesA,
		stats.Stats.GILessonDefusePlantedBomb,
		stats.Stats.GILessonCsgoInstrExplainFollowBomber,
		stats.Stats.GILessonCsgoInstrExplainPickupBomb,
		stats.Stats.GILessonCsgoInstrExplainPreventBombPickup,
		stats.Stats.GILessonCsgoCycleWeaponsKb,
		stats.Stats.GILessonCsgoInstrExplainZoom,
		stats.Stats.GILessonCsgoInstrExplainReload,
		stats.Stats.GILessonTrExplainPlantBomb,
		stats.Stats.GILessonBombSitesB,
		stats.Stats.GILessonVersionNumber,
		stats.Stats.GILessonFindPlantedBomb,
		stats.Stats.GILessonCsgoHostageLeadToHrz,
		stats.Stats.GILessonCsgoInstrRescueZone,
		stats.Stats.GILessonCsgoInstrExplainInspect,
		stats.Stats.SteamStatXpearnedgames,
	); err != nil {
		log.Fatal(err)
	}

	rows, err := result.RowsAffected()
	log.Println("Rows affected:", rows)

}

func (ds *DataStorage) UpdateRecentlyPlayedGames(rpg steamclient.RecentlyPlayedGames) {

	if _, err := ds.statements["update_recently_played"].Exec(
		rpg.Playtime2Weeks,
		rpg.PlaytimeForever,
		rpg.PlaytimeWindowsForever,
		rpg.PlaytimeMacForever,
		rpg.PlaytimeLinuxForever,
		rpg.SteamID,
	); err != nil {
		log.Fatal(err)
	}
}

// DataStorage is the main interface to the saved data. It provides methods for
// retrieval as well as methods to ingress new data from the API or update
// existing values
type DataStorage struct {
	db         *sql.DB
	statements map[string]*sql.Stmt
}

// GetPlayerInfoBySteamID returns a PlayerInfo from a steamID. It will try to
// get the needed values from the database and return an error if steamID
// cannot be found in it.
func (ds *DataStorage) GetPlayerInfoBySteamID(steamID string) (steamclient.PlayerInfo, error) {

	info := steamclient.PlayerInfo{}
	var err error

	if info.PlayerSummary, err = ds.GetPlayerSummary(steamID); err != nil {
		return info, err
	}

	if info.RecentlyPlayedGames, err = ds.GetRecentlyPlayedGames(steamID); err != nil {
		return info, err
	}

	if info.UserStatsForGame, err = ds.GetUserStatsForGame(steamID); err != nil {
		return info, err
	}

	if info.PlayerHistory, err = ds.GetPlayerHistory(steamID); err != nil {
		return info, err
	}

	return info, nil
}

// NewDataStorage creates a new DataStorage for a given sqlite database filepath
func NewDataStorage(path string) (*DataStorage, error) {
	var err error

	// Initialize database
	storage := new(DataStorage)
	storage.statements = make(map[string]*sql.Stmt)

	log.Println("Reading", path)
	if storage.db, err = sql.Open("sqlite3", path); err != nil {
		log.Fatal("Failed to open sqlite file", err)
	}

	// Prepare CREATE statements
	if err = storage.getCreatePreparedstatements(); err != nil {
		log.Fatal("Failed to prepare CREATE statements", err)
	}

	// Create tables, if necessary
	if _, err = storage.statements["create_player_summary"].Exec(); err != nil {
		log.Fatal("Failed to create table player_summary", err)
	}

	if _, err = storage.statements["create_player_stats"].Exec(); err != nil {
		log.Fatal("Failed to create table player_stats", err)
	}
	if _, err = storage.statements["create_recently_played"].Exec(); err != nil {
		log.Fatal("Failed to create table recently_played", err)
	}
	if _, err = storage.statements["create_player_history"].Exec(); err != nil {
		log.Fatal("Failed to create table player_history", err)
	}

	// Prepare remaining statements
	if err = storage.getUpdatePreparedstatements(); err != nil {
		log.Fatal("Failed to prepare UPDATE statements", err)
	}

	if err = storage.getInsertPreparedstatements(); err != nil {
		log.Fatal("Failed to prepare INSERT statements", err)
	}

	if err = storage.getSelectPreparedstatements(); err != nil {
		log.Fatal("Failed to prepare SELECT statements", err)
	}

	// TODO fix this
	// for _, v := range config.SteamIDs {
	// 	log.Println("Updating Data for ID:", v)

	//TODO functions changed to accept data and NOT neeed a client themselves
	// 	storage.updatePlayerSummary(v)
	// 	storage.updateRecentlyPlayedGames(v)
	// 	storage.updateUserStatsForGame(v)

	// }

	return storage, nil
}
