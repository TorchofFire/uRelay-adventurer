package packets

import (
	"encoding/json"

	"github.com/TorchofFire/uRelay-adventurer/internal/types"
)

type BasePacket struct {
	Type types.BasePacket `json:"type"`
	Data json.RawMessage  `json:"data"`
}

type Handshake struct { // type: handshake
	Name      string `json:"name"`
	PublicKey string `json:"public_key"`
	Proof     string `json:"proof"`
	/*
		proof is timestamp with server id signed.
		example of decoded: unixtimestamp|ipOrDomain(:port)
		"|" being a delimiter.
		The server id is important so the handshake is server specific and cannot be replayed elsewhere.
	*/
}

type GuildMessage struct { // type: guild_message
	ChannelId int    `json:"channel_id"`
	SenderId  int    `json:"sender_id"`
	Message   string `json:"message"`
	Id        int    `json:"id"`
}

type SystemMessage struct { // type: system_message
	Severity  types.Severity `json:"severity"`
	Message   string         `json:"message"`
	ChannelId int            `json:"channel_id"`
}
