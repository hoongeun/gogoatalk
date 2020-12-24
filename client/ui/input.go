package ui

import (
	"context"
	"log"
	"strings"

	"github.com/Hoongeun/gogoatalk/client/socket"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type messageInput struct {
	view *tview.InputField
}

func NewMessageInput(s *socket.Socket) *messageInput {
	var text string = ""
	view := tview.NewInputField().
		SetLabel("Message: ").
		SetPlaceholder("Press enter to send message").
		SetFieldWidth(0).
		SetChangedFunc(func(t string) {
			text = t
		})
	view.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			if len(strings.TrimSpace(text)) == 0 {
				text = ""
				view.SetText("")
				return
			}
			err := s.SendMessage(context.Background(), text)
			if err != nil {
				log.Fatalf("SendMessage(): %s", err)
				return
			}
			text = ""
			view.SetText("")
		}
	})
	return &messageInput{
		view: view,
	}
}
