package demoparser

// Score holds the information about one players score. It normally represents
// the summarized information of all matches the player has played
type Score struct {

	//ID of the player
	SteamId string `db:"steamid"`

	//Points earned from won matches
	ScoreMatch int `db:"score_match"`

	//Points earned from won rounds
	ScoreRound int `db:"score_round"`

	//Points earned from planting bomb
	ScorePlant int `db:"score_plant"`

	//Points earned from defusing bomb
	ScoreDefuse int `db:"score_defuse"`

	//Points earned from MVP stars
	ScoreMvp int `db:"score_mvp"`

	//Points earned from kills
	ScoreKill int `db:"score_kill"`

	//Points earned from assist
	ScoreAssist int `db:"score_assist"`

	//Points earned from 3-kills
	Score3k int `db:"score_3k"`

	//Points earned from 4-kills
	Score4k int `db:"score_4k"`

	//Points earned from aces
	Score5k int `db:"score_5k"`

	//Points earned from entryfrag
	ScoreEntryfrag int `db:"score_entryfrag"`

	//Points earned from clutches v1
	ScoreClutchv1 int `db:"score_clutchv1"`

	//Points earned from clutches v2
	ScoreClutchv2 int `db:"score_clutchv2"`

	//Points earned from clutches v3
	ScoreClutchv3 int `db:"score_clutchv3"`

	//Points earned from clutches v4
	ScoreClutchv4 int `db:"score_clutchv4"`

	//Points earned from clutches v5
	ScoreClutchv5 int `db:"score_clutchv5"`

	//Points lost due to teamkills
	ScoreTeamkill int `db:"score_teamkill"`

	//Points lost due to teamflashes
	ScoreTeamflash int `db:"score_teamflash"`

	//Points lost due to teamdamage
	ScoreTeamdamage int `db:"score_teamdamage"`

	//Points lost due to kick
	ScoreKick int `db:"score_kick"`

	//Points lost due to dying with HE/knive
	ScoreNoWeapon int `db:"score_noweapon"`
}
