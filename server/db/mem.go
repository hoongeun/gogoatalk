package db

import (
	"errors"
	"fmt"
	"time"

	c "github.com/Hoongeun/gogoatalk/common"
	"github.com/Hoongeun/gogoatalk/common/util"
)

type MemStorage struct {
	m []c.Message
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		m: make([]c.Message, 0),
	}
}

var ErrCantFindMessage = errors.New("Cannot find the message")

func (m *MemStorage) ReadLatest() []c.Message {
	start := util.MathMax(0, len(m.m)-50)
	return m.m[start:]
}

func (m *MemStorage) ReadMore(id string, more int) ([]c.Message, error) {
	for i, msg := range m.m {
		if msg.Id == id {
			start := util.MathMax(0, i-more-1)
			end := i
			return m.m[start:end], nil
		}
	}
	return nil, ErrCantFindMessage
}

func (m *MemStorage) Append(userid string, text string) *c.Message {
	now := time.Now().Unix()
	id := fmt.Sprintf("%d%s", now, util.RandString(6))
	newMessage := c.Message{
		Id:        id,
		Text:      text,
		Userid:    userid,
		CreatedAt: now,
		UpdatedAt: now,
	}
	m.m = append(m.m, newMessage)
	return &newMessage
}

func (m *MemStorage) Remove(userid string, id string) (*c.Message, error) {
	for i, msg := range m.m {
		if msg.Id == id {
			if userid != m.m[i].Userid {
				return nil, errors.New("Cannot remove message: Not an author")
			}
			m.m = m.m[:i+copy(m.m[i:], m.m[i+1:])]
			return &m.m[i], nil
		}
	}
	return nil, ErrCantFindMessage
}

func (m *MemStorage) Update(userid string, id string, text string) (*c.Message, error) {
	for i, msg := range m.m {
		if msg.Id == id {
			if userid != m.m[i].Userid {
				return nil, errors.New("Cannot update message: Not an author")
			}
			now := time.Now().Unix()
			m.m[i].Text = text
			m.m[i].UpdatedAt = now
			return &m.m[i], nil
		}
	}
	return nil, ErrCantFindMessage
}
