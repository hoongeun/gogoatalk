package core

import (
	"sync"
)

type presents struct {
	usermap map[string]string
	myid    string
}

var (
	instance *presents
	once     sync.Once
)

func GetPresents() *presents {
	once.Do(func() {
		instance = &presents{
			usermap: make(map[string]string),
		}
	})
	return instance
}

func (self *presents) SetMyId(userid string) {
	self.myid = userid
}

func (self *presents) FindUsername(username string) string {
	if v, ok := self.usermap[username]; ok {
		return v
	}
	return ""
}

func (self *presents) IsMe(userid string) bool {
	return self.myid == userid
}

func (self *presents) GetUsernames() []string {
	u := make([]string, 0, len(self.usermap))
	for _, username := range self.usermap {
		u = append(u, username)
	}
	return u
}

func (self *presents) AddUser(userid string, username string) {
	self.usermap[userid] = username
}

func (self *presents) RemoveUser(userid string) {
	delete(self.usermap, userid)
}
