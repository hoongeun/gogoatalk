package app

import (
	"context"

	"github.com/Hoongeun/gogoatalk/client/socket"
	"github.com/Hoongeun/gogoatalk/client/ui"
)

type App struct {
	s   *socket.Socket
	ui  *ui.UI
	ctx context.Context
}

func NewApp(ctx context.Context) *App {
	return &App{
		s:   socket.NewSocket(ctx),
		ctx: ctx,
	}
}

func (a *App) Start(host string) error {
	err := a.s.Connect(a.ctx, host)
	if err != nil {
		return err
	}
	a.ui = ui.NewUI(a.ctx, a.s)
	a.s.RegisterSocketListener(a.ui)
	a.ui.Init()
	return nil
}
