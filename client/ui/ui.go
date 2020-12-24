package ui

import (
	"context"

	"github.com/Hoongeun/gogoatalk/client/socket"
	pb "github.com/Hoongeun/gogoatalk/protobuf"
	"github.com/rivo/tview"
)

const pageCount = 5

var (
	pages *tview.Pages
)

type UI struct {
	app   *tview.Application
	pages *tview.Pages
	lp    *LoginForm
	mp    *Mainframe
	s     *socket.Socket
}

func NewUI(ctx context.Context, s *socket.Socket) *UI {
	app := tview.NewApplication()
	pages := tview.NewPages()
	loginform := NewLoginForm(ctx, app, s, pages)
	mainframe := NewMainframe(ctx, app, s)
	mainframe.init()
	pages.AddPage("loginpage",
		loginform.view, true, true)
	pages.AddPage("mainframe",
		mainframe.view, true, false)
	return &UI{
		s:     s,
		app:   app,
		pages: pages,
		lp:    loginform,
		mp:    mainframe,
	}
}

func (self *UI) Init() {
	if err := self.app.SetRoot(self.pages, true).SetFocus(self.pages).Run(); err != nil {
		panic(err)
	}
}

func (self *UI) OnUserLogin(userid string, username string, presents []*pb.Present) {
	self.mp.OnUserLogin(userid, username, presents)
	go self.app.Draw()
}

func (self *UI) OnOtherUserLogin(userid string, username string) {
	self.mp.OnOtherUserLogin(userid, username)
	go self.app.Draw()
}

func (self *UI) OnUserLogout(userid string, username string) {
	self.mp.OnUserLogout(userid, username)
	go self.app.Draw()
}

func (self *UI) OnReadLatest(messages []*pb.Message) {
	self.mp.OnReadLatest(messages)
	go self.app.Draw()
}

func (self *UI) OnReadMore(messages []*pb.Message) {
	self.mp.OnReadMore(messages)
	go self.app.Draw()
}

func (self *UI) OnSendMessage(m pb.Message) {
	self.mp.OnSendMessage(m)
	go self.app.Draw()
}

func (self *UI) OnServerShutdown() {
	self.mp.OnServerShutdown()
	go self.app.Stop()
}
