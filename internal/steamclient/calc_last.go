package steamclient

import "strconv"

func (pi PlayerInfo) LastMatch() LastMatch {

	outcome := "DRAW"

	if wins, err := strconv.Atoi(pi.UserStatsForGame.Stats.LastMatchWins); err == nil {
		if rounds, err := strconv.Atoi(pi.UserStatsForGame.Stats.LastMatchRounds); err == nil {
			if wins > (rounds / 2) {
				outcome = "WON"
			} else if wins < (rounds / 2) {
				outcome = "LOST"
			}
		}
	}

	//TODO add correct values here!

	favWeapon := getWeaponByID(pi.UserStatsForGame.Stats.LastMatchFavweaponID)

	return LastMatch{
		Outcome:           outcome,
		FavWeaponIconPath: favWeapon.IconPath,
		FavWeaponName:     favWeapon.Name,
		FavWeaponAccuracy: divideStringFloats(pi.UserStatsForGame.Stats.LastMatchFavweaponHits, pi.UserStatsForGame.Stats.LastMatchFavweaponShots),
	}
}

type LastMatch struct {
	Outcome           string
	FavWeaponIconPath string
	FavWeaponName     string
	FavWeaponAccuracy string
}
