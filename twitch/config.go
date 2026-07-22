package twitch

const (
	ClientID    string = "1jcuu1fyzg8nabsmoplijb826zoyte0"
	RedirectURI string = "http://localhost:8000/twitch/token/"
	// Use the implicit flow to keep the existing local redirect-based login flow working.
	RequestTokenURL string = "" +
		"https://id.twitch.tv/oauth2/authorize" +
		"?response_type=token" +
		"&client_id=" + ClientID +
		"&redirect_uri=" + RedirectURI +
		"&force_verify=true" +
		"&scope=channel:manage:broadcast+channel:read:redemptions"

	TwitchAPIURL string = "https://api.twitch.tv/helix"
)

var (
	Token       = ""
	GameList    TopGamesResponse
	GameDB      []DBGame
	UserChannel Channel
)

type (
	DBGame struct {
		TwitchName       string   `json:"twitchName"`
		FilePath         string   `json:"filePath"`
		AlternativeNames []string `json:"alternativeName"`
	}
)
