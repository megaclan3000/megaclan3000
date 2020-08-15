package demoparser

// Score holds the information about one players score. It normally represents
// the summarized information of all matches the player has played
type Score struct {

	//ID of the player
	SteamId string

	//Points earned from won matches
	ScoreMatch int

	//Points earned from won rounds
	ScoreRound int

	//Points earned from planting bomb
	ScorePlant int

	//Points earned from defusing bomb
	ScoreDefuse int

	//Points earned from MVP stars
	ScoreMvp int

	//Points earned from kills
	ScoreKill int

	//Points earned from assist
	ScoreAssist int

	//Points earned from 3-kills
	Score3k int

	//Points earned from 4-kills
	Score4k int

	//Points earned from aces
	Score5k int

	//Points earned from entryfrag
	ScoreEntryfrag int

	//Points earned from clutches v1
	ScoreClutchv1 int

	//Points earned from clutches v2
	ScoreClutchv2 int

	//Points earned from clutches v3
	ScoreClutchv3 int

	//Points earned from clutches v4
	ScoreClutchv4 int

	//Points earned from clutches v5
	ScoreClutchv5 int

	//Points lost due to teamkills
	ScoreTeamkill int

	//Points lost due to teamflashes
	ScoreTeamflash int

	//Points lost due to teamdamage
	ScoreTeamdamage int

	//Points lost due to kick
	ScoreKick int

	//Points lost due to dying with HE/knive
	ScoreNoWeapon int
}
