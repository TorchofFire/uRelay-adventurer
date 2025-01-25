package connections

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/TorchofFire/uRelay-adventurer/internal/models"
	"github.com/TorchofFire/uRelay-adventurer/internal/packets"
	"github.com/TorchofFire/uRelay-adventurer/internal/profile"
	"github.com/TorchofFire/uRelay-adventurer/internal/types"
)

type userGotten struct {
	ID        uint64 `json:"id"`
	PublicKey string `json:"public_key"`
	Name      string `json:"name"`
}

func httpGetRequest(secure bool, serverAddress, route string) ([]byte, error) {
	httProtocol := "http"
	if secure {
		httProtocol = "https"
	}
	request := fmt.Sprintf("%s://%s/%s", httProtocol, serverAddress, route)
	resp, err := http.Get(request)
	if err != nil {
		return nil, fmt.Errorf("failed http GET request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read http response body: %v", err)
	}
	return body, nil
}

func updateUsers(secure bool, serverAddress string) error {
	body, err := httpGetRequest(secure, serverAddress, "users")
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	var users []userGotten
	err = json.Unmarshal(body, &users)
	if err != nil {
		return fmt.Errorf("failed to parse http response body to JSON: %v", err)
	}

	ServersMu.Lock()
	defer ServersMu.Unlock()

	var personalID *uint64
	for _, user := range users {
		Servers[serverAddress].Users[user.ID] = models.Users{
			ID:        user.ID,
			PublicKey: user.PublicKey,
			Name:      user.Name,
		}

		if user.PublicKey == profile.Profile.PublicKey {
			personalID = &user.ID
		}
	}

	Servers[serverAddress].PersonalID = personalID

	if personalID == nil {
		return fmt.Errorf("could not find self in fetched users")
	}

	return nil
}

type channelsAndCategories struct {
	Channels   []models.GuildChannels   `json:"channels"`
	Categories []models.GuildCategories `json:"categories"`
}

func updateChannelsAndCategories(secure bool, serverAddress string) error {
	body, err := httpGetRequest(secure, serverAddress, "channels")
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	var channelsAndCategories channelsAndCategories
	err = json.Unmarshal(body, &channelsAndCategories)
	if err != nil {
		return fmt.Errorf("failed to parse http response body to JSON: %v", err)
	}

	ServersMu.Lock()
	defer ServersMu.Unlock()

	for _, channel := range channelsAndCategories.Channels {
		Servers[serverAddress].Channels[channel.ID] = ChannelData{
			Channel:  channel,
			Messages: make(map[uint64]types.GuildMessageEmission),
		}
	}

	return nil
}

func GetMessagesFromTextChannel(serverAddress string, channelId, msgId uint64) ([]types.GuildMessageEmission, error) {
	ServersMu.Lock()
	secure := Servers[serverAddress].Secure
	ServersMu.Unlock()

	route := fmt.Sprintf("text-channel/%d", channelId)
	if msgId != 0 {
		route = fmt.Sprintf("%s?msg=%d", route, msgId)
	}

	body, err := httpGetRequest(secure, serverAddress, route)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	var messages []models.GuildMessages
	err = json.Unmarshal(body, &messages)
	if err != nil {
		return nil, fmt.Errorf("failed to parse http response body to JSON: %v", err)
	}

	var dataToEmit []types.GuildMessageEmission

	for _, message := range messages {
		pMessage := packets.GuildMessage{
			ChannelId: message.ChannelID,
			SenderId:  message.SenderID,
			Message:   message.Message,
			Id:        message.ID,
		}
		msg, err := turnMsgPacketToEmit(pMessage, serverAddress)
		if err != nil {
			continue
		}
		dataToEmit = append(dataToEmit, msg)
	}
	if len(dataToEmit) == 0 {
		dataToEmit = []types.GuildMessageEmission{}
	}
	return dataToEmit, nil
}
