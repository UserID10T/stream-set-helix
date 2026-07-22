package twitch

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

type HelixChannelResponse struct {
	Data []HelixChannel `json:"data"`
}
