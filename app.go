package main

import (
	"context"
	"fmt"

	"github.com/TorchofFire/uRelay-adventurer/internal/connections"
	"github.com/TorchofFire/uRelay-adventurer/internal/profile"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	profile.Init()
	go connections.NewConnection(ctx, false, "localhost:8080")
}

func (a *App) SendMessage(serverId, message string, channelId int) {
	server, exists := connections.Servers[serverId]
	if !exists {
		fmt.Printf("Connection not found for serverId: %s\n", serverId)
		return
	}
	connections.SendMessage(server.Conn, message, *connections.Servers[serverId].PersonalID, channelId)
	// TODO: return error
}
