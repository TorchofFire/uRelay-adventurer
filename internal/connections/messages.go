package connections

import (
	"encoding/json"

	"github.com/TorchofFire/uRelay-adventurer/internal/packets"
	"github.com/TorchofFire/uRelay-adventurer/internal/profile"
	"github.com/TorchofFire/uRelay-adventurer/internal/types"
	"github.com/gorilla/websocket"
)

func handshake(conn *websocket.Conn, serverAddress string) error {
	proof, err := signMessage(profile.Profile.PrivateKey, serverAddress)
	if err != nil {
		return err
	}
	handshakePacket := packets.Handshake{
		Name:      profile.Profile.Name,
		PublicKey: profile.Profile.PublicKey,
		Proof:     proof,
	}
	handshakePacketJSON, err := json.Marshal(handshakePacket)
	if err != nil {
		return err
	}
	packet := packets.BasePacket{
		Type: types.Handshake,
		Data: handshakePacketJSON,
	}
	if err := conn.WriteJSON(packet); err != nil {
		conn.Close()
		return err
	}
	return nil
}

func SendMessage(conn *websocket.Conn, message string, userId, channelId int) error {
	msg, err := signMessage(profile.Profile.PrivateKey, message)
	if err != nil {
		return err
	}
	messagePacket := packets.GuildMessage{
		ChannelId: channelId,
		SenderId:  userId,
		Message:   msg,
		Id:        0,
	}
	messagePacketJSON, err := json.Marshal(messagePacket)
	if err != nil {
		return err
	}
	packet := packets.BasePacket{
		Type: types.GuildMessage,
		Data: messagePacketJSON,
	}
	if err := conn.WriteJSON(packet); err != nil {
		conn.Close()
		return err
	}
	return nil
}
