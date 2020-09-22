package steamclient

// LastMatch is a helper function to access the dynamically caluclated values
// from the templates
func (pi PlayerInfo) LastMatch() LastMatch {

	outcome := "DRAW"

	if pi.UserStatsForGame.Stats.LastMatchWins > (pi.UserStatsForGame.Stats.LastMatchRounds / 2) {
		outcome = "WON"
	} else if pi.UserStatsForGame.Stats.LastMatchWins < (pi.UserStatsForGame.Stats.LastMatchRounds / 2) {
		outcome = "LOST"
	}

	favWeapon := getWeaponByID(pi.UserStatsForGame.Stats.LastMatchFavweaponID)

	return LastMatch{
		Outcome:           outcome,
		FavWeaponIconPath: favWeapon.IconPath,
		FavWeaponName:     favWeapon.Name,
		FavWeaponAccuracy: divideNoZero(pi.UserStatsForGame.Stats.LastMatchFavweaponHits, pi.UserStatsForGame.Stats.LastMatchFavweaponShots),
	}
}

// LastMatch hold additional stats calculated to be viewed in the templates
type LastMatch struct {
	Outcome           string
	FavWeaponIconPath string
	FavWeaponName     string
	FavWeaponAccuracy string
}
