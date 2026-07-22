package twitch

const (
	ClientID    string = "1jcuu1fyzg8nabsmoplijb826zoyte0"
	RedirectURI string = "http://localhost:8000/twitch/token/"
	// Keep the existing local redirect-based login flow.
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

type DBGame struct {
	TwitchName       string   `json:"twitchName"`
	FilePath         string   `json:"filePath"`
	AlternativeNames []string `json:"alternativeName"`
}

type UserResponse struct {
	Data []User `json:"data"`
}

type User struct {
	ID              string `json:"id"`
	Login           string `json:"login"`
	DisplayName     string `json:"display_name"`
	Type            string `json:"type"`
	BroadcasterType string `json:"broadcaster_type"`
	Description     string `json:"description"`
	ProfileImageURL string `json:"profile_image_url"`
	OfflineImageURL string `json:"offline_image_url"`
	ViewCount       int64  `json:"view_count"`
	Email           string `json:"email,omitempty"`
	CreatedAt       string `json:"created_at"`
}

type ChannelResponse struct {
	Data []Channel `json:"data"`
}

type Channel struct {
	BroadcasterID    string `json:"broadcaster_id"`
	BroadcasterLogin string `json:"broadcaster_login"`
	BroadcasterName  string `json:"broadcaster_name"`
	BroadcasterLanguage string `json:"broadcaster_language"`
	GameID           string `json:"game_id"`
	GameName         string `json:"game_name"`
	Title            string `json:"title"`
	Delay            int    `json:"delay"`
}

type SearchCategoriesResponse struct {
	Data []SearchCategory `json:"data"`
}

type SearchCategory struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BoxArtURL string `json:"box_art_url"`
}

type TopGamesResponse struct {
	Data []TopGame `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type TopGame struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BoxArtURL string `json:"box_art_url"`
}
