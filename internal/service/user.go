package service

import "fmt"

type User struct {
	ur UserRepo
}

func newUser(ur UserRepo) *User {
	return &User{
		ur: ur,
	}
}

func (u *User) Login() (string, error) {
	fmt.Println("zaebka ebanaya")
	return "", nil
}

func (u *User) SaveLoginData() {
}
