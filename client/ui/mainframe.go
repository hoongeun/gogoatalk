package ui

import (
	"context"
	"fmt"

	"github.com/Hoongeun/gogoatalk/client/socket"
	pb "github.com/Hoongeun/gogoatalk/protobuf"
	"github.com/rivo/tview"
)

type Mainframe struct {
	sidebar          *Sidebar
	messageContainer *messageContainer
	messageInputBar  *messageInput
	view             *tview.Grid
}

func NewMainframe(ctx context.Context, app *tview.Application, s *socket.Socket) *Mainframe {
	sidebar := NewSidebar()
	messageContainer := NewMessageContainer(s)
	messageInputBar := NewMessageInput(s)
	view := NewLayout(sidebar, messageContainer, messageInputBar)
	return &Mainframe{
		sidebar,
		messageContainer,
		messageInputBar,
		view,
	}
}

func (self *Mainframe) init() {
	self.sidebar.init()
	self.messageContainer.init()
}

func (self *Mainframe) OnUserLogin(userid string, username string, presents []*pb.Present) {
	self.sidebar.initPresents(userid, username, presents)
}

func (self *Mainframe) OnOtherUserLogin(userid string, username string) {
	self.sidebar.addPresent(userid, username)
	self.messageContainer.printNotification(fmt.Sprintf("%s enter the chat", username))
}

func (self *Mainframe) OnUserLogout(userid string, username string) {
	self.sidebar.removePresent(userid, username)
	self.messageContainer.printNotification(fmt.Sprintf("%s leave the chat", username))
}

func (self *Mainframe) OnReadLatest(messages []*pb.Message) {
	self.messageContainer.prependMessages(messages)
}

func (self *Mainframe) OnReadMore(messages []*pb.Message) {
	self.messageContainer.prependMessages(messages)
}

func (self *Mainframe) OnSendMessage(message pb.Message) {
	self.messageContainer.appendMessage(message)
}

func (self *Mainframe) OnServerShutdown() {
	self.messageContainer.printCriticalSysMessage("Server is shutdown")
}
