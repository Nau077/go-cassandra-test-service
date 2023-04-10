package internal

import (
	"context"
	"errors"
	"log"

	"github.com/Nau077/cassandra-golang-sv/internal/infrustructure"
)

type App struct {
	provider   *infrustructure.Provider
	staticPath string
}

func NewApp(ctx context.Context, staticPath string) (*App, error) {
	a := &App{
		staticPath: staticPath,
	}
	a.initDeps(ctx)

	return a, nil
}

func (a *App) initDeps(ctx context.Context) {
	a.initServiceProvider(ctx)
}

func (a *App) initServiceProvider(ctx context.Context) error {
	a.provider = infrustructure.NewProvider(ctx, a.staticPath)

	return nil
}

func (a *App) Run() error {
	// defer func() {
	// 	a.provider.GetDB().Close()
	// }()

	// cделать метод runServer() и там стартовать роутер
	// в старте сервера вложить туда service из провайдера, куда будет вложен репозиторий
	err := a.provider.GetDb().Init()
	if err != nil {
		return errors.New("failed to connect db")
	}

	err = a.provider.GetRouter().Init()
	if err != nil {
		return errors.New("failed to connect router")
	}

	return nil

}

func (a *App) RunHttpServer() error {
	err := a.provider.GetRouter().Init()
	if err != nil {
		log.Fatalf("failed to process muxer: #{err.Error()}")
	}

	return nil

}
