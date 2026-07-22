package twitch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/lxn/walk"

	"github.com/asimshrestha2/stream-set/guicontroller"
	"github.com/asimshrestha2/stream-set/save"
)

type Channel struct {
	Mature                       bool   `json:"mature"`
	Status                       string `json:"status"`
	BroadcasterLanguage          string `json:"broadcaster_language"`
	DisplayName                  string `json:"display_name"`
	Game                         string `json:"game"`
	Language                     string `json:"language"`
	ID                           string `json:"_id"`
	Name                         string `json:"name"`
	CreatedAt                    string `json:"created_at"`
	UpdatedAt                    string `json:"updated_at"`
	Partner                      bool   `json:"partner"`
	Logo                         string `json:"logo"`
	VideoBanner                  string `json:"video_banner"`
	ProfileBanner                string `json:"profile_banner"`
	ProfileBannerBackgroundColor string `json:"profile_banner_background_color"`
	URL                          string `json:"url"`
	Views                        int64  `json:"views"`
	Followers                    int64  `json:"followers"`
	BroadcasterType              string `json:"broadcaster_type"`
	StreamKey                    string `json:"stream_key"`
	Email                        string `json:"email"`
}

type ChannelG struct {
	ChannelA GameC `json:"channel"`
}

type GameC struct {
	GameA string `json:"game"`
}

var (
	imageSet = false
)

func SetTwitchChannel() {
	UserChannel = GetChannelInfo()
	if guicontroller.MW.TwitchUsername != nil {
		guicontroller.MW.TwitchUsername.SetText(UserChannel.DisplayName)
	}
	if guicontroller.MW.TwitchGame != nil {
		guicontroller.MW.TwitchGame.SetText(UserChannel.Game)
	}
	if !imageSet && UserChannel.Logo != "" {
		save.Image(UserChannel.Logo, func() {
			imageSet = true
			img, err := walk.NewImageFromFile(save.ImagePathFromURL(UserChannel.Logo))
			if err != nil {
				imageSet = false
				log.Println(err)
				return
			}
			if guicontroller.MW.TwitchImage != nil {
				guicontroller.MW.TwitchImage.SetImage(img)
			}
		})
	}
}

func GetChannelInfo() Channel {
	body, err := Request("GET", TwitchAPIURL+"/channels", nil, false, false)
	if err != nil {
		fmt.Println("GetChannelInfo error:", err)
		return Channel{}
	}
	ch := Channel{}
	if err := json.Unmarshal([]byte(body), &ch); err != nil {
		fmt.Println("GetChannelInfo decode error:", err)
		return Channel{}
	}
	return ch
}

func UpdateChannelGame(game string) {
	if game != "" {
		resC := &ChannelG{
			ChannelA: GameC{
				GameA: game,
			},
		}

		res2B, _ := json.Marshal(resC)

		if _, err := Request("PATCH", TwitchAPIURL+"/channels", bytes.NewBuffer(res2B), true, true); err != nil {
			log.Println(err)
		} else {
			SetTwitchChannel()
		}
	}
}
