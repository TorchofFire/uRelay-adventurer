package emitters

import (
	"context"

	"github.com/TorchofFire/uRelay-adventurer/internal/types"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func EmitGuildMessage(ctx context.Context, data types.GuildMessageEmission) {
	runtime.EventsEmit(ctx, "guild_message", data)
}

func EmitSystemMessage(ctx context.Context, data types.SystemMessageEmission) {
	runtime.EventsEmit(ctx, "system_message", data)
}
