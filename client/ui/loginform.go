package ui

import (
	"context"
	"log"

	"github.com/Hoongeun/gogoatalk/client/socket"
	"github.com/rivo/tview"
)

type LoginForm struct {
	view *tview.Form
}

func NewLoginForm(ctx context.Context, app *tview.Application, s *socket.Socket, pages *tview.Pages) *LoginForm {
	var (
		Username string
		Password string
	)
	form := tview.NewForm().
		AddInputField("UserName", "", 60, nil, func(text string) {
			Username = text
		}).
		AddPasswordField("Password", "", 60, '*', func(text string) {
			Password = text
		}).
		AddButton("Login", func() {
			err := s.Login(ctx, Username, Password)
			if err != nil {
				log.Fatalf("Failed to Login")
			}
			pages.SwitchToPage("mainframe")
			s.OnEnterChatroom(ctx)
			s.ReadLatest(ctx)
		}).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("User Login").SetTitleAlign(tview.AlignLeft)
	return &LoginForm{
		view: form,
	}
}
