package db

import (
	c "github.com/Hoongeun/gogoatalk/common"
)

type QueryMethods interface {
	ReadLatest() []c.Message
	ReadMore(id string, more int) ([]c.Message, error)
	Append(userid string, text string) *c.Message
	Remove(userid string, id string) (*c.Message, error)
	Update(userid string, id string, text string) (*c.Message, error)
}
