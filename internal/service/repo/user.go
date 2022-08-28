package repo

import (
	"sync"
)

type User struct {
	mp   map[string]string
	code map[string]string
	mw   sync.Mutex
}

func NewUser() *User {
	return &User{
		mp:   make(map[string]string),
		code: make(map[string]string),
	}
}

func (u *User) SaveLoginCode(ip, code string) {
	u.code[ip] = code
}

func (u *User) GetLoginCode(ip string) string {
	return u.code[ip]
}

func (u *User) SaveUsersState(ip, state string) {
	u.mw.Lock()
	u.mp[ip] = state
	u.mw.Unlock()

}

func (u *User) CheckStateByIP(ip, state string) bool {
	u.mw.Lock()
	defer u.mw.Unlock()
	savedState, ok := u.mp[ip]
	if !ok {
		return false
	}

	return savedState == state
}
