package service

import "fmt"

type User struct {
	UserRepo
}

func newUser(ur UserRepo) *User {
	return &User{
		UserRepo: ur,
	}
}

func (u *User) Login() {
	u.UserRepo.SaveLoginData()
	fmt.Println("посетили логин сервис")
}

func (u *User) SaveLoginData() {
}
