package connections

import (
	"encoding/json"

	"github.com/TorchofFire/uRelay-adventurer/internal/packets"
	"github.com/TorchofFire/uRelay-adventurer/internal/profile"
	"github.com/TorchofFire/uRelay-adventurer/internal/types"
	"github.com/gorilla/websocket"
)

func Handshake(conn *websocket.Conn, serverAddress string) error {
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

func SendMessage(conn *websocket.Conn, message string, userId, channelId uint64) error {
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

func turnMsgPacketToEmit(p packets.GuildMessage, serverAddress string) (types.GuildMessageEmission, error) {
	ServersMu.Lock()
	pk := Servers[serverAddress].Users[p.SenderId].PublicKey
	name := Servers[serverAddress].Users[p.SenderId].Name
	ServersMu.Unlock()
	plainMsg, time, err := unlockSignedMessage(pk, p.Message)
	if err != nil {
		return types.GuildMessageEmission{}, err
	}

	dataToEmit := types.GuildMessageEmission{
		GuildID:    serverAddress,
		ID:         p.Id,
		ChannelID:  p.ChannelId,
		SenderID:   p.SenderId,
		SenderName: name,
		Message:    plainMsg,
		SentAt:     uint32(time.Unix()),
	}

	ServersMu.Lock()
	defer ServersMu.Unlock()
	Servers[serverAddress].Channels[dataToEmit.ChannelID].Messages[dataToEmit.ID] = dataToEmit

	return dataToEmit, nil
}
