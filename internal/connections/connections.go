package connections

import (
	"context"
	"fmt"
	"log"

	"github.com/TorchofFire/uRelay-adventurer/internal/emitters"
	"github.com/TorchofFire/uRelay-adventurer/internal/packets"
	"github.com/gorilla/websocket"
)

func NewConnection(ctx context.Context, secure bool, serverAddress string) {
	protocol := "ws"
	if secure {
		protocol = "wss"
	}
	fullWsAddress := fmt.Sprintf("%s://%s", protocol, serverAddress)

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

	err = handshake(conn, serverAddress)
	if err != nil {
		log.Println(err)
		conn.Close()
	}

	for {
		messageType, packet, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Failed to read message: %v", err)
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
			emitters.EmitGuildMessage(ctx, p)
		case packets.Handshake:
			// do nothing
		case packets.SystemMessage:
			emitters.EmitSystemMessage(ctx, p)
		default:
			log.Fatal("A deserialized and known packet was not handled")
		}
	}
}
