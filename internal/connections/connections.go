package connections

import (
	"context"
	"fmt"
	"log"

	"github.com/TorchofFire/uRelay-adventurer/internal/emitters"
	"github.com/TorchofFire/uRelay-adventurer/internal/packets"
	"github.com/TorchofFire/uRelay-adventurer/internal/types"
	"github.com/gorilla/websocket"
)

func NewConnection(ctx context.Context, secure bool, serverAddress string) {
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

	addNewConnection(serverAddress, conn)
	defer func() {
		conn.Close()
		removeConnection(serverAddress)
	}()

	fmt.Println("Connected to WebSocket server at", fullWsAddress)

	err = Handshake(conn, serverAddress)
	if err != nil {
		log.Println(err)
		conn.Close()
	}

	ServersMu.Lock()
	Servers[serverAddress].Secure = secure
	ServersMu.Unlock()
	updateUsers(secure, serverAddress)
	updateChannelsAndCategories(secure, serverAddress)

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
		deserializedPacket, err := packets.DeserializePacket(packet)
		if err != nil {
			log.Println(err)
			return
		}
		switch p := deserializedPacket.(type) {
		case packets.GuildMessage:
			dataToEmit, err := turnMsgPacketToEmit(p, serverAddress)
			if err != nil {
				return // TODO: handle error
			}
			emitters.EmitGuildMessage(ctx, dataToEmit)
		case packets.Handshake:
			// do nothing
		case packets.SystemMessage:
			dataToEmit := types.SystemMessageEmission{
				GuildID:   serverAddress,
				Severity:  p.Severity,
				Message:   p.Message,
				ChannelId: p.ChannelId,
			}
			emitters.EmitSystemMessage(ctx, dataToEmit)
		default:
			log.Fatal("A deserialized and known packet was not handled")
		}
	}
}
