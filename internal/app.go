package internal

import "context"

type App struct {
	provider   *telegram.Provider
	staticPath string
}

func NewApp(ctx context.Context, staticPath string) (*App, error) {
	a := &App{
		staticPath: staticPath,
	}
	err := a.initDeps(ctx)

	return a, err
}

func (a *App) Run() error {
	defer func() {
		a.provider.GetDB().Close()
	}()

}
