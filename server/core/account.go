package core

import (
	"errors"
	"os"
	"path/filepath"
	"sync"

	pb "github.com/Hoongeun/gogoatalk/protobuf"
	"github.com/spf13/viper"
)

type Present struct {
	Userid   string
	Username string
}

type Account struct {
	Userid   string `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccountManager struct {
	accounts []Account
	presents map[string]Present // map[userid] = Account
	mtx      sync.Mutex
}

func NewAccountManager() *AccountManager {
	return &AccountManager{}
}

func (am *AccountManager) LoadAccounts() error {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	cfgFile := filepath.Join(dir, "accounts.json")
	viper.SetConfigType("json")
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	am.accounts = make([]Account, 0)
	am.presents = make(map[string]Present)
	if err := viper.UnmarshalKey("accounts", &am.accounts); err != nil {
		return err
	}
	return nil
}

var (
	ErrCantFindUser  = errors.New("Cannot find user")
	ErrWrongPassword = errors.New("Password is wrong")
	ErrAlreadyLogin  = errors.New("Already Login")
)

func (am *AccountManager) ValidateUserAccount(username string, password string) (string, error) {
	for _, a := range am.accounts {
		if a.Username == username {
			if a.Password == password {
				return a.Userid, nil
			}
			return "", ErrWrongPassword
		}
	}

	return "", ErrCantFindUser
}

func (am *AccountManager) AppendPresents(userid string) error {
	if _, ok := am.presents[userid]; ok {
		return ErrAlreadyLogin
	}
	for _, a := range am.accounts {
		if a.Userid == userid {
			am.mtx.Lock()
			am.presents[userid] = Present{
				Userid:   a.Userid,
				Username: a.Username,
			}
			am.mtx.Unlock()
			return nil
		}
	}
	return ErrCantFindUser
}

func (am *AccountManager) DeletePresent(userid string) {
	am.mtx.Lock()
	delete(am.presents, userid)
	am.mtx.Unlock()
}

func (am *AccountManager) GetPresents() []*pb.Present {
	var ret []*pb.Present
	for _, a := range am.presents {
		p := &pb.Present{
			Userid:   a.Userid,
			Username: a.Username,
		}
		ret = append(ret, p)
	}
	return ret
}

func (am *AccountManager) GetAccount(userid string) (*Account, error) {
	for _, a := range am.accounts {
		if a.Userid == userid {
			return &a, nil
		}
	}
	return nil, ErrCantFindUser
}

func (am *AccountManager) GetPresent(userid string) (*Present, error) {
	if p, ok := am.presents[userid]; ok {
		return &p, nil
	}
	return nil, ErrCantFindUser
}
