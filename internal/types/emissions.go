package types

type GuildMessageEmission struct {
	GuildID    string `json:"guild_id"`
	ID         int    `json:"id"`
	SenderID   int    `json:"sender_id"`
	SenderName string `json:"sender_name"`
	Message    string `json:"message"`
	ChannelID  int    `json:"channel_id"`
	SentAt     int    `json:"sent_at"`
}

type SystemMessageEmission struct {
	GuildID   string   `json:"guild_id"`
	Severity  Severity `json:"severity"`
	Message   string   `json:"message"`
	ChannelId int      `json:"channel_id"`
}
