package connections

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/TorchofFire/uRelay-adventurer/internal/packets"
	"github.com/TorchofFire/uRelay-adventurer/internal/types"
)

func (s *Service) httpGetRequest(secure bool, serverAddress, route string) ([]byte, error) {
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

func (s *Service) updateUsers(secure bool, serverAddress string) error {
	body, err := s.httpGetRequest(secure, serverAddress, "users")
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	var users []types.Users
	err = json.Unmarshal(body, &users)
	if err != nil {
		return fmt.Errorf("failed to parse http response body to JSON: %v", err)
	}

	s.serversMu.Lock()
	defer s.serversMu.Unlock()

	var personalID uint64
	for _, user := range users {
		s.servers[serverAddress].Users[user.ID] = types.Users{
			ID:        user.ID,
			PublicKey: user.PublicKey,
			Name:      user.Name,
			Status:    user.Status,
		}

		if user.PublicKey == s.profile.Profile.PublicKey {
			personalID = user.ID
		}
	}

	s.servers[serverAddress].PersonalID = personalID

	if personalID == 0 {
		return fmt.Errorf("could not find self in fetched users")
	}

	return nil
}

func (s *Service) updateChannelsAndCategories(secure bool, serverAddress string) error {
	body, err := s.httpGetRequest(secure, serverAddress, "channels")
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	var channelsAndCategories types.ChannelsAndCategories
	err = json.Unmarshal(body, &channelsAndCategories)
	if err != nil {
		return fmt.Errorf("failed to parse http response body to JSON: %v", err)
	}

	s.serversMu.Lock()
	defer s.serversMu.Unlock()

	for _, channel := range channelsAndCategories.Channels {
		s.servers[serverAddress].Channels[channel.ID] = ChannelData{
			Channel:  channel,
			Messages: make(map[uint64]types.GuildMessageEmission),
		}
	}
	for _, category := range channelsAndCategories.Categories {
		s.servers[serverAddress].Categories[category.ID] = category
	}

	return nil
}

func (s *Service) GetMessagesFromTextChannel(serverAddress string, channelId, msgId uint64) ([]types.GuildMessageEmission, error) {
	s.serversMu.Lock()
	server, serverExists := s.servers[serverAddress]
	if !serverExists {
		s.serversMu.Unlock()
		return nil, fmt.Errorf("server connection not found")
	}
	secure := server.Secure

	cachedMsgs := server.Channels[channelId].Messages

	s.serversMu.Unlock()
	if len(cachedMsgs) > 0 && func() bool {
		var lowestId uint64 = ^uint64(0)
		for id := range cachedMsgs {
			if lowestId == 0 || id < lowestId {
				lowestId = id
			}
		}
		return msgId != lowestId
	}() {
		msgsSlice := make([]types.GuildMessageEmission, 0, len(cachedMsgs))
		for _, msg := range cachedMsgs {
			msgsSlice = append(msgsSlice, msg)
		}
		return msgsSlice, nil
	}

	route := fmt.Sprintf("text-channel/%d", channelId)
	if msgId != 0 {
		if msgId == 1 {
			return []types.GuildMessageEmission{}, nil
		}
		route = fmt.Sprintf("%s?msg=%d", route, msgId-1)
	}

	body, err := s.httpGetRequest(secure, serverAddress, route)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	var messages []types.GuildMessages
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
		msg, err := s.turnMsgPacketToEmit(pMessage, serverAddress)
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

func (s *Service) GetUsersSliceFromServer(serverAddress string) ([]types.Users, error) {
	s.serversMu.Lock()
	defer s.serversMu.Unlock()
	server, exists := s.servers[serverAddress]
	if !exists {
		return nil, fmt.Errorf("server not found: %s", serverAddress)
	}

	var usersSlice []types.Users
	if server.Users != nil {
		usersSlice = make([]types.Users, 0, len(server.Users))
		for _, user := range server.Users {
			usersSlice = append(usersSlice, user)
		}
	}
	return usersSlice, nil
}

func (s *Service) GetChannelsAndCategories(serverAddress string) (types.ChannelsAndCategories, error) {
	s.serversMu.Lock()
	defer s.serversMu.Unlock()
	server, exists := s.servers[serverAddress]
	if !exists {
		return types.ChannelsAndCategories{}, fmt.Errorf("server not found: %s", serverAddress)
	}

	var channelsSlice []types.GuildChannels
	if server.Channels != nil {
		channelsSlice = make([]types.GuildChannels, 0, len(server.Channels))
		for _, channel := range server.Channels {
			channelsSlice = append(channelsSlice, channel.Channel)
		}
	}

	var categoriesSlice []types.GuildCategories
	if server.Categories != nil {
		categoriesSlice = make([]types.GuildCategories, 0, len(server.Categories))
		for _, category := range server.Categories {
			categoriesSlice = append(categoriesSlice, category)
		}
	}

	return types.ChannelsAndCategories{
		Channels:   channelsSlice,
		Categories: categoriesSlice,
	}, nil
}
