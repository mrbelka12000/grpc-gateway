package repo

import "fmt"

type User struct{}

func NewUser() *User {
	return &User{}
}

func (u *User) SaveLoginData() {
	fmt.Println("посетили логин репо")
}
