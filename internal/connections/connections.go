package connections

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/TorchofFire/uRelay-adventurer/internal/emitters"
	"github.com/TorchofFire/uRelay-adventurer/internal/packets"
	"github.com/TorchofFire/uRelay-adventurer/internal/profile"
	"github.com/TorchofFire/uRelay-adventurer/internal/types"
	"github.com/gorilla/websocket"
)

type Service struct {
	profile   *profile.Service
	emitters  *emitters.Service
	packets   *packets.Service
	servers   map[string]*ServerData
	serversMu sync.Mutex
}

func NewService(profile *profile.Service, emitters *emitters.Service, packets *packets.Service) *Service {
	s := &Service{profile: profile, emitters: emitters, packets: packets}
	s.servers = make(map[string]*ServerData)
	return s
}

func (s *Service) NewConnection(ctx context.Context, secure bool, serverAddress string) {
	wsProtocol := "ws"
	if secure {
		wsProtocol = "wss"
	}
	fullWsAddress := fmt.Sprintf("%s://%s", wsProtocol, serverAddress)

	dialer := websocket.DefaultDialer
	conn, _, err := dialer.Dial(fullWsAddress, nil)
	if err != nil {
		log.Printf("Failed to connect to WebSocket server: %v", err)
		// TODO: send error to fontend
		return
	}

	s.addNewConnection(serverAddress, conn)
	defer func() {
		conn.Close()
		s.removeConnection(serverAddress)
	}()

	fmt.Println("Connected to WebSocket server at", fullWsAddress)

	err = s.handshake(conn, serverAddress)
	if err != nil {
		log.Println(err)
		conn.Close()
	}

	s.serversMu.Lock()
	s.servers[serverAddress].Secure = secure
	s.serversMu.Unlock()
	s.updateUsers(secure, serverAddress)
	s.updateChannelsAndCategories(secure, serverAddress)

	for {
		messageType, packet, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Failed to read message: %v", err)
			return
			// TODO: check that the connection closes when forcibly closed by host
		}
		if messageType != websocket.TextMessage {
			return
		}
		deserializedPacket, err := s.packets.DeserializePacket(packet)
		if err != nil {
			log.Println(err)
			return
		}
		switch p := deserializedPacket.(type) {
		case packets.GuildMessage:
			dataToEmit, err := s.turnMsgPacketToEmit(p, serverAddress)
			if err != nil {
				return // TODO: handle error
			}
			s.emitters.EmitGuildMessage(ctx, dataToEmit)
		case packets.Handshake:
			// do nothing
		case packets.SystemMessage:
			dataToEmit := types.SystemMessageEmission{
				GuildID:   serverAddress,
				Severity:  p.Severity,
				Message:   p.Message,
				ChannelId: p.ChannelId,
			}
			s.emitters.EmitSystemMessage(ctx, dataToEmit)
		case packets.User:
			// TODO: emit
		default:
			log.Fatal("A deserialized and known packet was not handled")
		}
	}
}
