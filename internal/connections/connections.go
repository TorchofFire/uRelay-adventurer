package connections

import (
	"fmt"
	"log"

	"github.com/TorchofFire/uRelay-adventurer/internal/packets"
	"github.com/gorilla/websocket"
)

// TODO: rework to connect to multiple guilds at startup

func NewConnection(secure bool, serverAddress string) {
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
	defer conn.Close()
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
			log.Println(p)
		case packets.Handshake:
			// do nothing
		case packets.SystemMessage:
			log.Println(p)
		default:
			log.Fatal("A deserialized and known packet was not handled")
		}
	}
}
