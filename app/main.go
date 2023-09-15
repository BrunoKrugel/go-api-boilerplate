package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/BrunoKrugel/go-api-boilerplate/internal/api"
	"github.com/BrunoKrugel/zaplog"
)

type App struct {
	MetaServer *api.Meta
	HTTPServer *api.HTTP
}

func main() {
	ctx := context.Background()

	zaplog.StartLogger()

	app := NewApp()

	if err := app.Run(ctx); err != nil {
		zaplog.Fatal(ctx, err.Error())
	}
}

func NewApp() *App {
	return &App{
		MetaServer: api.NewMeta(),
		HTTPServer: api.NewHTTP(),
	}
}

func (app *App) Run(ctx context.Context) error {
	go func() {
		api.StartHttp(ctx, app.HTTPServer)
	}()

	go func() {
		api.StartMeta(ctx, app.MetaServer)
	}()

	handleShutdown(ctx, app)

	return nil
}

func handleShutdown(ctx context.Context, app *App) {

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	switch <-signalChannel {

	case os.Interrupt:
		zaplog.Info(context.Background(), "Received SIGINT, stopping...")
		app.MetaServer.Server.Shutdown(ctx)
		app.HTTPServer.Server.Shutdown(ctx)
	case syscall.SIGTERM:
		zaplog.Info(context.Background(), "Received SIGTERM, stopping...")
		app.MetaServer.Server.Shutdown(ctx)
		app.HTTPServer.Server.Shutdown(ctx)
	}
}
