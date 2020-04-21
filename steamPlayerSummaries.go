package main

type PlayerSummariesData struct {
	// https://developer.valvesoftware.com/wiki/Steam_Web_API#GetPlayerSummaries_.28v0002.29

	Response struct {
		Players []struct {

			// Public Data

			// 64bit SteamID of the user
			Steamid string `json:"steamid"`

			// The player's persona name (display name)
			Personaname string `json:"personaname"`

			// The full URL of the player's Steam Community profile.
			Profileurl string `json:"profileurl"`

			// The full URL of the player's 32x32px avatar. If the user
			// hasn't configured an avatar, this will be the default ?
			// avatar.
			Avatar string `json:"avatar"`

			// The full URL of the player's 64x64px avatar. If the user
			// hasn't configured an avatar, this will be the default ?
			// avatar.
			Avatarmedium string `json:"avatarmedium"`

			// The full URL of the player's 184x184px avatar. If the
			// user hasn't configured an avatar, this will be the
			// default ? avatar.
			Avatarfull string `json:"avatarfull"`

			// 0 - Offline
			// 1 - Online
			// 2 -Busy
			// 3 - Away
			// 4 - Snooze
			// 5 - looking to trade
			// 6 - looking to play.
			// The user's current status.  If the player's profile is private,
			// this will always be "0", except if the user has set
			// their status to looking to trade or looking to play,
			// because a bug makes those status appear even if the
			// profile is private.
			Personastate int `json:"personastate"`

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
			Communityvisibilitystate int `json:"communityvisibilitystate"`

			// If set, indicates the user has a community profile
			// configured (will be set to '1')
			Profilestate int `json:"profilestate"`

			// The last time the user was online, in unix time. Only
			// available when you are friends with the requested user
			// (since Feb, 4).
			Lastlogoff int `json:"lastlogoff"`

			// If set, indicates the profile allows public comments.
			Commentpermission string `json:"commentpermission"`

			// Private Data

			// The player's "Real Name", if they have set it.
			Realname string `json:"realname"`

			// The player's primary group, as configured in their Steam Community profile.
			Primaryclanid string `json:"primaryclanid"`

			// The time the player's account was created.
			Timecreated int `json:"timecreated"`

			// If the user is currently in-game, this value will be
			// returned and set to the gameid of that game.
			Gameid string `json:"gameid"`

			//     gameserverip
			// The ip and port of the game server the user is currently
			// playing on, if they are playing on-line in a game using
			// Steam matchmaking. Otherwise will be set to "0.0.0.0:0".
			Gameserverip string `json:"gameserverip"`

			//     gameextrainfo
			// If the user is currently in-game, this will be the name
			// of the game they are playing. This may be the name of a
			// non-Steam game shortcut.
			Gameextrainfo string `json:"gameextrainfo"`

			//     cityid
			// This value will be removed in a future update (see loccityid)
			Cityid string `json:"cityid"`

			// If set on the user's Steam Community profile, The user's
			// country of residence, 2-character ISO country code
			Loccountrycode string `json:"loccountrycode"`

			//     locstatecode
			// If set on the user's Steam Community profile, The user's state of residence
			Locstatecode string `json:"locstatecode"`

			// An internal code indicating the user's city of
			// residence. A future update will provide this data in a
			// more useful way.  steam_location gem/package makes
			// player location data readable for output.
			Loccityid string `json:"loccityid"`
		} `json:"players"`
	} `json:"response"`
}
