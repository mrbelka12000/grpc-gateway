package repo

import "fmt"

type User struct{}

func newUser() *User {
	return &User{}
}

func (u *User) SaveLoginData() {
	fmt.Println("посетили логин репо")
}
