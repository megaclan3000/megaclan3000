package steamclient

type Calculations struct {
	PercentWin      string
	PercentAccuracy string
	PercentHeadshot string
}

func (pi PlayerInfo) Calc() Calculations {

	return Calculations{
		PercentWin:      percentStringFloats(pi.UserStatsForGame.Stats.TotalMatchesWon, pi.UserStatsForGame.Stats.TotalMatchesPlayed),
		PercentAccuracy: percentStringFloats(pi.UserStatsForGame.Stats.TotalShotsHit, pi.UserStatsForGame.Stats.TotalShotsFired),
		PercentHeadshot: percentStringFloats(pi.UserStatsForGame.Stats.TotalKillsHeadshot, pi.UserStatsForGame.Stats.TotalKills),
	}
}
