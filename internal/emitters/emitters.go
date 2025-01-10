package emitters

import (
	"context"

	"github.com/TorchofFire/uRelay-adventurer/internal/packets"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func EmitGuildMessage(ctx context.Context, packet packets.GuildMessage) {
	runtime.EventsEmit(ctx, "guild_message", packet)
}

func EmitSystemMessage(ctx context.Context, packet packets.SystemMessage) {
	runtime.EventsEmit(ctx, "system_message", packet)
}
