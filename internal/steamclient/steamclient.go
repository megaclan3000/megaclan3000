package steamclient

type SteamClient struct {
	config SteamConfig
}

func NewSteamClient() *SteamClient {

	var config SteamConfig

	config = readConfig()
	config.Refresh()
	return &SteamClient{config}
}
