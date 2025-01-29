package main

import (
	"context"
	"fmt"

	"github.com/TorchofFire/uRelay-adventurer/internal/connections"
	"github.com/TorchofFire/uRelay-adventurer/internal/packets"
	"github.com/TorchofFire/uRelay-adventurer/internal/types"
)

// App struct
type App struct {
	ctx         context.Context
	connections *connections.Service
}

// NewApp creates a new App application struct
func NewApp(connections *connections.Service) *App {
	a := &App{connections: connections}
	return a
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	go a.connections.NewConnection(ctx, false, "localhost:8080")
}

func (a *App) SendMessage(serverId, message string, channelId uint64) error {

	server, err := a.connections.GetServer(serverId)
	if err != nil {
		fmt.Printf("%v", err)
	}

	err = a.connections.SendMessage(server.Conn, message, *server.PersonalID, channelId)

	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}
	return nil
}

func (a *App) GetMessages(serverId string, channelId, msgId uint64) ([]types.GuildMessageEmission, error) {
	return a.connections.GetMessagesFromTextChannel(serverId, channelId, msgId)
}

func (a *App) GetUsers(serverId string) ([]packets.User, error) {
	return a.connections.GetUsersSliceFromServer(serverId)
}
