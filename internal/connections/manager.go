package connections

import (
	"sync"

	"github.com/TorchofFire/uRelay-adventurer/internal/models"
	"github.com/TorchofFire/uRelay-adventurer/internal/types"
	"github.com/gorilla/websocket"
)

type ChannelData struct {
	Channel  models.GuildChannels
	Messages map[uint64]types.GuildMessageEmission `json:"messages"`
}

type ServerData struct {
	Conn       *websocket.Conn
	Secure     bool
	PersonalID *uint64
	Channels   map[uint64]ChannelData
	Users      map[uint64]models.Users
}

var (
	Servers   = make(map[string]*ServerData)
	ServersMu sync.Mutex
)

func addNewConnection(serverId string, conn *websocket.Conn) {
	ServersMu.Lock()
	defer ServersMu.Unlock()
	Servers[serverId] = &ServerData{
		Conn:     conn,
		Channels: make(map[uint64]ChannelData),
		Users:    make(map[uint64]models.Users),
	}
}

func removeConnection(serverId string) {
	ServersMu.Lock()
	defer ServersMu.Unlock()
	delete(Servers, serverId)
}
