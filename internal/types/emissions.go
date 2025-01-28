package types

type Severity string

const (
	Info    Severity = "info"
	Warning Severity = "warning"
	Danger  Severity = "danger"
)

type GuildMessageEmission struct {
	GuildID    string `json:"guild_id"`
	ID         uint64 `json:"id"`
	SenderID   uint64 `json:"sender_id"`
	SenderName string `json:"sender_name"`
	Message    string `json:"message"`
	ChannelID  uint64 `json:"channel_id"`
	SentAt     uint32 `json:"sent_at"`
}

type SystemMessageEmission struct {
	GuildID   string   `json:"guild_id"`
	Severity  Severity `json:"severity"`
	Message   string   `json:"message"`
	ChannelId uint64   `json:"channel_id"`
}
