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

type ChannelResponse struct {
	Data []HelixChannel `json:"data"`
}

type HelixChannel struct {
	BroadcasterID   string `json:"broadcaster_id"`
	BroadcasterName string `json:"broadcaster_name"`
	BroadcasterLogin string `json:"broadcaster_login"`
	GameID          string `json:"game_id"`
	GameName        string `json:"game_name"`
	Title           string `json:"title"`
	Language        string `json:"broadcaster_language"`
	Delay           int    `json:"delay"`
}

var (
	imageSet = false
)

func SetTwitchChannel() {
	if guicontroller.MW.TwitchUsername == nil || guicontroller.MW.TwitchGame == nil {
		return
	}

	user, err := CurrentUser()
	if err != nil || user == nil {
		log.Println("SetTwitchChannel: current user unavailable:", err)
		return
	}

	guicontroller.MW.TwitchUsername.SetText(user.DisplayName)

	channel, err := CurrentChannel(user.ID)
	if err != nil || channel == nil {
		log.Println("SetTwitchChannel: current channel unavailable:", err)
		return
	}

	guicontroller.MW.TwitchGame.SetText(channel.GameName)
	UserChannel = Channel{
		ID:          channel.BroadcasterID,
		DisplayName: user.DisplayName,
		Game:        channel.GameName,
		Name:        user.Login,
	}

	if !imageSet && user.ProfileImageURL != "" {
		save.Image(user.ProfileImageURL, func() {
			imageSet = true
			img, err := walk.NewImageFromFile(save.ImagePathFromURL(user.ProfileImageURL))
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
	user, err := CurrentUser()
	if err != nil || user == nil {
		fmt.Println("GetChannelInfo user error:", err)
		return Channel{}
	}

	channel, err := CurrentChannel(user.ID)
	if err != nil || channel == nil {
		fmt.Println("GetChannelInfo channel error:", err)
		return Channel{}
	}

	return Channel{
		DisplayName: user.DisplayName,
		Game:        channel.GameName,
		ID:          channel.BroadcasterID,
		Name:        user.Login,
		Logo:        user.ProfileImageURL,
	}
}

func UpdateChannelGame(game string) {
	if game == "" {
		return
	}

	user, err := CurrentUser()
	if err != nil || user == nil {
		log.Println("UpdateChannelGame: current user unavailable:", err)
		return
	}

	payload := map[string]string{"game_id": game}
	body, _ := json.Marshal(payload)

	if _, err := Request("PATCH", fmt.Sprintf("%s/channels?broadcaster_id=%s", TwitchAPIURL, user.ID), bytes.NewBuffer(body)); err != nil {
		log.Println(err)
		return
	}

	SetTwitchChannel()
}
