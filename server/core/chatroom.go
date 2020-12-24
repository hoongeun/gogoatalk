package core

import (
	"errors"

	c "github.com/Hoongeun/gogoatalk/common"
	"github.com/Hoongeun/gogoatalk/server/db"
)

type ChatRoomManager struct {
	m  []c.Message
	db db.QueryMethods
}

func NewChatRoomManager() *ChatRoomManager {
	return &ChatRoomManager{
		m:  make([]c.Message, 0),
		db: db.NewMemStorage(),
	}
}

var ErrDBUnInitialized = errors.New("DB Implementation is not initiated")

func (cm *ChatRoomManager) ReadLatest() ([]c.Message, error) {
	return cm.db.ReadLatest(), nil
}

func (cm *ChatRoomManager) ReadMore(id string, more int) ([]c.Message, error) {
	return cm.db.ReadMore(id, more)
}

func (cm *ChatRoomManager) Append(userid string, text string) (*c.Message, error) {
	m := cm.db.Append(userid, text)
	return m, nil
}

func (cm *ChatRoomManager) Remove(userid string, id string) (*c.Message, error) {
	m, err := cm.db.Remove(userid, id)
	return m, err
}

func (cm *ChatRoomManager) Update(userid string, id string, text string) (*c.Message, error) {
	m, err := cm.db.Update(userid, id, text)
	return m, err
}
