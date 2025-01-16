package connections

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/TorchofFire/uRelay-adventurer/internal/models"
	"github.com/TorchofFire/uRelay-adventurer/internal/profile"
)

type userGotten struct {
	ID        int    `json:"id"`
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

	var personalID *int
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

func updateChannels(secure bool, serverAddress string) error {
	body, err := httpGetRequest(secure, serverAddress, "channels")
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	var channels []models.GuildChannels
	err = json.Unmarshal(body, &channels)
	if err != nil {
		return fmt.Errorf("failed to parse http response body to JSON: %v", err)
	}

	ServersMu.Lock()
	defer ServersMu.Unlock()

	for _, channel := range channels {
		Servers[serverAddress].Channels[channel.ID] = ChannelData{
			Channel: channel,
		}
	}

	return nil
}
