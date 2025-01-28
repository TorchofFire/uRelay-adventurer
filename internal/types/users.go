package types

type Status string

const (
	Online Status = "online"
	// Idle    Status = "idle"
	// DnD     Status = "dnd"
	Offline Status = "offline"
)

type Users struct {
	ID        uint64 `json:"id"`
	PublicKey string `json:"public_key"`
	Name      string `json:"name"`
	Status    Status `json:"status"`
}
