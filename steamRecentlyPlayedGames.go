package main

// https://developer.valvesoftware.com/wiki/Steam_Web_API#GetRecentlyPlayedGames_.28v0001.29
// http://api.steampowered.com/IPlayerService/GetRecentlyPlayedGames/v0001/?key=XXXXXXXXXXXXXXXXX&steamid=76561197960434622&format=json

type RecentlyPlayedGamesData struct {
	Response struct {

		// the total number of unique games the user has played in the last
		// two weeks. This is mostly significant if you opted to return a
		// limited number of games with the count input parameter
		TotalCount int `json:"total_count"`

		// A games array, with the following contents:
		Games []struct {

			// Unique identifier for the game
			Appid int `json:"appid"`

			// The name of the game
			Name string `json:"name"`

			// playtime_2weeks The total number of minutes played in the last 2 weeks
			Playtime2Weeks int `json:"playtime_2weeks"`

			// playtime_forever The total number of minutes played "on record", since Steam began tracking total playtime in early 2009.
			PlaytimeForever int `json:"playtime_forever"`

			// img_icon_url, img_logo_url - these are the filenames of various
			// images for the game. To construct the URL to the image, use this
			// format:
			// http://media.steampowered.com/steamcommunity/public/images/apps/{appid}/{hash}.jpg
			ImgIconURL string `json:"img_icon_url"`
			ImgLogoURL string `json:"img_logo_url"`

			//Playtime on different operating systems
			PlaytimeWindowsForever int `json:"playtime_windows_forever"`
			PlaytimeMacForever     int `json:"playtime_mac_forever"`
			PlaytimeLinuxForever   int `json:"playtime_linux_forever"`
		} `json:"games"`
	} `json:"response"`
}
