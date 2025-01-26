package main

import (
	"embed"

	"github.com/TorchofFire/uRelay-adventurer/internal/connections"
	"github.com/TorchofFire/uRelay-adventurer/internal/emitters"
	"github.com/TorchofFire/uRelay-adventurer/internal/packets"
	"github.com/TorchofFire/uRelay-adventurer/internal/profile"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	profileService := profile.NewService()
	profileService.Init()
	emittersService := emitters.NewService()
	packetsService := packets.NewService()
	connectionsService := connections.NewService(profileService, emittersService, packetsService)
	app := NewApp(connectionsService)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "uRelay Adventurer",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 9, G: 9, B: 9, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
