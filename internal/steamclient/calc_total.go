package steamclient

// Calculations holds additional values to be shown in the templates, that are
// calculated on-the-fly
type Calculations struct {
	PercentWin      string
	PercentAccuracy string
	PercentHeadshot string
}

// Calc is a helper function to access the dynamically caluclated values
// from the templates
func (pi PlayerInfo) Calc() Calculations {

	return Calculations{
		PercentWin:      percentStringFloats(pi.UserStatsForGame.Stats.TotalMatchesWon, pi.UserStatsForGame.Stats.TotalMatchesPlayed),
		PercentAccuracy: percentStringFloats(pi.UserStatsForGame.Stats.TotalShotsHit, pi.UserStatsForGame.Stats.TotalShotsFired),
		PercentHeadshot: percentStringFloats(pi.UserStatsForGame.Stats.TotalKillsHeadshot, pi.UserStatsForGame.Stats.TotalKills),
	}
}
