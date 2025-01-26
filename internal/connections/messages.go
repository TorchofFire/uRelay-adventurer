package connections

import (
	"encoding/json"

	"github.com/TorchofFire/uRelay-adventurer/internal/packets"
	"github.com/TorchofFire/uRelay-adventurer/internal/types"
	"github.com/gorilla/websocket"
)

func (s *Service) handshake(conn *websocket.Conn, serverAddress string) error {
	proof, err := s.signMessage(s.profile.Profile.PrivateKey, serverAddress)
	if err != nil {
		return err
	}
	handshakePacket := packets.Handshake{
		Name:      s.profile.Profile.Name,
		PublicKey: s.profile.Profile.PublicKey,
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

func (s *Service) SendMessage(conn *websocket.Conn, message string, userId, channelId uint64) error {
	msg, err := s.signMessage(s.profile.Profile.PrivateKey, message)
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

func (s *Service) turnMsgPacketToEmit(p packets.GuildMessage, serverAddress string) (types.GuildMessageEmission, error) {
	s.serversMu.Lock()
	pk := s.servers[serverAddress].Users[p.SenderId].PublicKey
	name := s.servers[serverAddress].Users[p.SenderId].Name
	s.serversMu.Unlock()
	plainMsg, time, err := s.unlockSignedMessage(pk, p.Message)
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

	s.serversMu.Lock()
	defer s.serversMu.Unlock()
	s.servers[serverAddress].Channels[dataToEmit.ChannelID].Messages[dataToEmit.ID] = dataToEmit

	return dataToEmit, nil
}
