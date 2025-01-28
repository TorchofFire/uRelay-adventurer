package types

type ChannelsAndCategories struct {
	Channels   []GuildChannels   `json:"channels"`
	Categories []GuildCategories `json:"categories"`
}

type GuildCategories struct {
	ID              uint64 `json:"id" db:"id"`
	Name            string `json:"name" db:"name"`
	DisplayPriority uint16 `json:"display_priority" db:"display_priority"`
}
