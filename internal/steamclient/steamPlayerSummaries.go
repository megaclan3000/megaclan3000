package steamclient

import (
	"encoding/json"
	"errors"
	"strconv"
)

type playerSummariesData struct {
	// https://developer.valvesoftware.com/wiki/Steam_Web_API#GetPlayerSummaries_.28v0002.29

	Response struct {
		Players []struct {
			Avatar                   string      `json:"avatar"`
			Avatarfull               string      `json:"avatarfull"`
			Avatarmedium             string      `json:"avatarmedium"`
			Cityid                   string      `json:"cityid"`
			Commentpermission        json.Number `json:"commentpermission"`
			Communityvisibilitystate int         `json:"communityvisibilitystate"`
			Gameextrainfo            string      `json:"gameextrainfo"`
			Gameid                   string      `json:"gameid"`
			Gameserverip             string      `json:"gameserverip"`
			Lastlogoff               int         `json:"lastlogoff"`
			Loccityid                string      `json:"loccityid"`
			Loccountrycode           string      `json:"loccountrycode"`
			Locstatecode             string      `json:"locstatecode"`
			Personaname              string      `json:"personaname"`
			Personastate             int         `json:"personastate"`
			Primaryclanid            string      `json:"primaryclanid"`
			Profilestate             int         `json:"profilestate"`
			Profileurl               string      `json:"profileurl"`
			Realname                 string      `json:"realname"`
			Steamid                  string      `json:"steamid"`
			Timecreated              int         `json:"timecreated"`
			SteamID                  uint64
		} `json:"players"`
	} `json:"response"`
}

// PlayerSummary holds the players summary data from the steam API
// endpoint GetPlayerSummaries
type PlayerSummary struct {

	// Public Data

	// 64bit SteamID of the user
	SteamID uint64 `db:"steamid"`

	// The player's persona name (display name)
	Personaname string `db:"personaname"`

	// The full URL of the player's Steam Community profile.
	Profileurl string `db:"profileurl"`

	// The full URL of the player's 32x32px avatar. If the user
	// hasn't configured an avatar, this will be the default ?
	// avatar.
	Avatar string `db:"avatar"`

	// The full URL of the player's 64x64px avatar. If the user
	// hasn't configured an avatar, this will be the default ?
	// avatar.
	Avatarmedium string `db:"avatarmedium"`

	// The full URL of the player's 184x184px avatar. If the
	// user hasn't configured an avatar, this will be the
	// default ? avatar.
	Avatarfull string `db:"avatarfull"`

	// 0 - Offline
	// 1 - Online
	// 2 - Busy
	// 3 - Away
	// 4 - Snooze
	// 5 - looking to trade
	// 6 - looking to play.
	// The user's current status. If the player's profile is private,
	// this will always be "0", except if the user has set
	// their status to looking to trade or looking to play,
	// because a bug makes those status appear even if the
	// profile is private.
	Personastate string `db:"personastate"`

	// This represents whether the profile is visible or not,
	// and if it is visible, why you are allowed to see it.
	// Note that because this WebAPI does not use
	// authentication, there are only two possible values
	// returned: 1 - the profile is not visible to you
	// (Private, Friends Only, etc), 3 - the profile is
	// "Public", and the data is visible. Mike Blaszczak's post
	// on Steam forums says, "The community visibility state
	// this API returns is different than the privacy state.
	// It's the effective visibility state from the account
	// making the request to the account being viewed given the
	// requesting account's relationship to the viewed
	// account."
	Communityvisibilitystate string `db:"communityvisibilitystate"`

	// If set, indicates the user has a community profile
	// configured (will be set to '1')
	Profilestate string `db:"profilestate"`

	// The last time the user was online, in unix time. Only
	// available when you are friends with the requested user
	// (since Feb, 4).
	Lastlogoff string `db:"lastlogoff"`

	// If set, indicates the profile allows public comments.
	Commentpermission string `db:"commentpermission"`

	// Private Data

	// This value will be removed in a future update (see loccityid)
	Cityid string `db:"cityid"`

	// If the user is currently in-game, this will be the name
	// of the game they are playing. This may be the name of a
	// non-Steam game shortcut.
	Gameextrainfo string `db:"gameextrainfo"`

	// If the user is currently in-game, this value will be
	// returned and set to the gameid of that game.
	Gameid string `db:"gameid"`

	// The ip and port of the game server the user is currently
	// playing on, if they are playing on-line in a game using
	// Steam matchmaking. Otherwise will be set to "0.0.0.0:0".
	Gameserverip string `db:"gameserverip"`

	// An internal code indicating the user's city of
	// residence. A future update will provide this data in a
	// more useful way.  steam_location gem/package makes
	// player location data readable for output.
	Loccityid string `db:"loccityid"`

	// If set on the user's Steam Community profile, The user's
	// country of residence, 2-character ISO country code
	Loccountrycode string `db:"loccountrycode"`

	// If set on the user's Steam Community profile, The user's state of residence
	Locstatecode string `db:"locstatecode"`

	// The player's primary group, as configured in their Steam Community profile.
	Primaryclanid string `db:"primaryclanid"`

	// The player's "Real Name", if they have set it.
	Realname string `db:"realname"`

	// The time the player's account was created.
	Timecreated string `db:"timecreated"`
}

func (sc *SteamClient) parsePlayerSummary(data playerSummariesData) (PlayerSummary, error) {

	if len(data.Response.Players) < 1 {
		return PlayerSummary{}, errors.New("Failed to parse PlayerSummary")
	}

	var steamID uint64
	var err error

	if steamID, err = strconv.ParseUint(data.Response.Players[0].Steamid, 10, 64); err != nil {
		return PlayerSummary{}, err
	}

	return PlayerSummary{
		Lastlogoff:               strconv.Itoa(data.Response.Players[0].Lastlogoff),
		Communityvisibilitystate: strconv.Itoa(data.Response.Players[0].Communityvisibilitystate),
		Personastate:             strconv.Itoa(data.Response.Players[0].Personastate),
		Profilestate:             strconv.Itoa(data.Response.Players[0].Profilestate),
		Timecreated:              strconv.Itoa(data.Response.Players[0].Timecreated),
		Avatar:                   data.Response.Players[0].Avatar,
		Avatarfull:               data.Response.Players[0].Avatarfull,
		Avatarmedium:             data.Response.Players[0].Avatarmedium,
		Cityid:                   data.Response.Players[0].Cityid,
		Commentpermission:        string(data.Response.Players[0].Commentpermission),
		Gameextrainfo:            data.Response.Players[0].Gameextrainfo,
		Gameid:                   data.Response.Players[0].Gameid,
		Gameserverip:             data.Response.Players[0].Gameserverip,
		Loccityid:                data.Response.Players[0].Loccityid,
		Loccountrycode:           data.Response.Players[0].Loccountrycode,
		Locstatecode:             data.Response.Players[0].Locstatecode,
		Personaname:              data.Response.Players[0].Personaname,
		Primaryclanid:            data.Response.Players[0].Primaryclanid,
		Profileurl:               data.Response.Players[0].Profileurl,
		Realname:                 data.Response.Players[0].Realname,
		SteamID:                  steamID,
	}, nil
}
