package emitters

import (
	"context"

	"github.com/TorchofFire/uRelay-adventurer/internal/types"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Service struct {
}

func NewService() *Service {
	s := &Service{}
	return s
}

func (s *Service) EmitGuildMessage(ctx context.Context, data types.GuildMessageEmission) {
	runtime.EventsEmit(ctx, "guild_message", data)
}

func (s *Service) EmitUser(ctx context.Context, data types.Users) {
	runtime.EventsEmit(ctx, "user", data)
}

func (s *Service) EmitSystemMessage(ctx context.Context, data types.SystemMessageEmission) {
	runtime.EventsEmit(ctx, "system_message", data)
}
