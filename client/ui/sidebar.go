package ui

import (
	"github.com/Hoongeun/gogoatalk/client/core"
	pb "github.com/Hoongeun/gogoatalk/protobuf"
	"github.com/rivo/tview"
)

type Sidebar struct {
	view  *tview.List
	users []string
}

func NewSidebar() *Sidebar {
	return &Sidebar{
		view: tview.NewList(),
	}
}

func (self *Sidebar) init() {
	self.update()
}

func (self *Sidebar) update() {
	self.users = core.GetPresents().GetUsernames()
	self.view.Clear()
	self.onUpdate()
}

func (self *Sidebar) onUpdate() {
	for i, u := range self.users {
		self.view.AddItem(u, "", rune(i+1+'0'), nil)
	}
}

func (self *Sidebar) initPresents(userid string, username string, presents []*pb.Present) {
	for _, p := range presents {
		core.GetPresents().AddUser(p.Userid, p.Username)
	}
	core.GetPresents().SetMyId(userid)
	self.update()
}

func (self *Sidebar) addPresent(userid string, username string) {
	core.GetPresents().AddUser(userid, username)
	self.update()
}

func (self *Sidebar) removePresent(userid string, username string) {
	core.GetPresents().RemoveUser(userid)
	self.update()
}
