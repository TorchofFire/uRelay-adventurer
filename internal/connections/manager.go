package connections

import (
	"fmt"

	"github.com/TorchofFire/uRelay-adventurer/internal/types"
	"github.com/gorilla/websocket"
)

type ChannelData struct {
	Channel  types.GuildChannels
	Messages map[uint64]types.GuildMessageEmission `json:"messages"`
}

type ServerData struct {
	Conn       *websocket.Conn
	Secure     bool
	PersonalID *uint64
	Channels   map[uint64]ChannelData
	Users      map[uint64]types.Users
}

func (s *Service) addNewConnection(serverId string, conn *websocket.Conn) {
	s.serversMu.Lock()
	defer s.serversMu.Unlock()
	s.servers[serverId] = &ServerData{
		Conn:     conn,
		Channels: make(map[uint64]ChannelData),
		Users:    make(map[uint64]types.Users),
	}
}

func (s *Service) removeConnection(serverId string) {
	s.serversMu.Lock()
	defer s.serversMu.Unlock()
	delete(s.servers, serverId)
}

func (s *Service) GetServer(serverId string) (*ServerData, error) {
	s.serversMu.Lock()
	defer s.serversMu.Unlock()
	server, exists := s.servers[serverId]
	if !exists {
		return nil, fmt.Errorf("server not found for server id: %s", serverId)
	}
	return server, nil
}
