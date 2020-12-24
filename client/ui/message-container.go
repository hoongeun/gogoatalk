package ui

import (
	"bytes"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Hoongeun/gogoatalk/client/core"
	"github.com/Hoongeun/gogoatalk/client/socket"
	"github.com/Hoongeun/gogoatalk/common"
	"github.com/Hoongeun/gogoatalk/common/util"
	pb "github.com/Hoongeun/gogoatalk/protobuf"
	"github.com/rivo/tview"
)

func msgFmt(m common.Message) string {
	if m.Userid == "system" {
		fmt.Sprintf("%s\n%s - %s", m.Text, core.GetPresents().FindUsername(m.Userid), util.ToLiteralTime(m.CreatedAt))
	} else if core.GetPresents().IsMe(m.Userid) {
		return fmt.Sprintf("%s\n[green]%s(Me)[white:black] - %s", m.Text, core.GetPresents().FindUsername(m.Userid), util.ToLiteralTime(m.CreatedAt))
	}
	return fmt.Sprintf("%s\n%s - %s", m.Text, core.GetPresents().FindUsername(m.Userid), util.ToLiteralTime(m.CreatedAt))
}

func convertPbMessage(m pb.Message) common.Message {
	return common.Message{
		Id:        m.Id,
		Userid:    m.Userid,
		Text:      m.Text,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

type messageContainer struct {
	view     *tview.TextView
	messages []common.Message
	s        *socket.Socket
	mtx      sync.Mutex
	cancel   context.CancelFunc
}

func NewMessageContainer(s *socket.Socket) *messageContainer {
	view := tview.NewTextView().
		SetDynamicColors(true).
		SetWordWrap(true)
	view.SetTitle("Message")
	view.SetBorder(true)

	return &messageContainer{
		view: view,
		s:    s,
	}
}

func (self *messageContainer) init() {
	ctx, cancel := context.WithCancel(context.Background())
	self.cancel = cancel
	go self.startTicker(ctx)
}

func (self *messageContainer) Stop() {
	self.cancel()
}

func (self *messageContainer) Add(m common.Message, prepend bool) {
	if prepend {
		self.messages = append([]common.Message{m}, self.messages...)
	} else {
		self.messages = append(self.messages, m)
	}
}

func (self *messageContainer) prependMessages(messages []*pb.Message) {
	for i := len(messages) - 1; i >= 0; i-- {
		m := convertPbMessage(*messages[i])
		self.Add(m, true)
	}
	self.Update()
}

func (self *messageContainer) Update() {
	b := bytes.NewBuffer([]byte{})
	for _, m := range self.messages {
		b.WriteString(fmt.Sprintf("%s\n\n", msgFmt(m)))
	}
	self.view.SetText(b.String())
	self.view.ScrollToEnd()
}

func (self *messageContainer) startTicker(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			self.Update()
		}
	}
}

func (self *messageContainer) printUserLogin(userid string, username string) {
	now := time.Now().Unix()
	m := common.Message{
		Id:        fmt.Sprintf("sys%d%s", now, util.RandString(10)),
		Userid:    "system",
		Text:      fmt.Sprintf("[yellow]%s enter the chat", username),
		CreatedAt: now,
		UpdatedAt: now,
	}
	self.Add(m, false)
	self.Update()
}

func (self *messageContainer) appendMessage(m pb.Message) {
	self.Add(convertPbMessage(m), false)
	self.Update()
}

func (self *messageContainer) printNotification(text string) {
	now := time.Now().Unix()
	m := common.Message{
		Id:        fmt.Sprintf("sys%d%s", now, util.RandString(10)),
		Userid:    "system",
		Text:      fmt.Sprintf("[yellow]%s", text),
		CreatedAt: now,
		UpdatedAt: now,
	}
	self.Add(m, false)
	self.Update()
}

func (self *messageContainer) printCriticalSysMessage(text string) {
	now := time.Now().Unix()
	m := common.Message{
		Id:        fmt.Sprintf("sys%d%s", now, util.RandString(10)),
		Userid:    "system",
		Text:      fmt.Sprintf("[red]%s", text),
		CreatedAt: now,
		UpdatedAt: now,
	}
	self.Add(m, false)
	self.Update()
}
