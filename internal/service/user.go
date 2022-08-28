package service

import (
	"fmt"
	"github.com/mrbelka12000/grpc-gateway/internal/models"
	"github.com/mrbelka12000/grpc-gateway/pkg/tools"
	"net/url"
	"os"
)

type User struct {
	ur UserRepo
}

func newUser(ur UserRepo) *User {
	return &User{
		ur: ur,
	}
}

func (u *User) GetLoginUrl(ip string) (string, error) {
	spotifyAccountHost, ok := os.LookupEnv("spotify_account_host")
	if !ok {
		return "", fmt.Errorf("no spotify_account_host in env variables")
	}
	spotifyLogin, ok := os.LookupEnv("spotify_login")
	if !ok {
		return "", fmt.Errorf("no spotify_login in env variables")
	}

	baseURL := spotifyAccountHost + spotifyLogin
	URL, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("parse baseURL error: %w", err)
	}

	state := tools.GetRandomString()
	q := url.Values{}
	q.Add("response_type", "code")
	q.Add("client_id", os.Getenv("client_id"))
	q.Add("scope", os.Getenv("scope"))
	q.Add("redirect_uri", os.Getenv("redirect_uri"))
	q.Add("state", state)
	URL.RawQuery = q.Encode()

	u.ur.SaveUsersState(ip, state)
	fmt.Println(URL.String())
	return URL.String(), nil
}

func (u *User) SaveLoginCode(ip string, loginResp *models.LoginResp) error {
	if !u.ur.CheckStateByIP(ip, loginResp.State) {
		return fmt.Errorf("state %v does not match", loginResp.State)
	}

	u.ur.SaveLoginCode(ip, loginResp.Code)
	return nil
}

func (u *User) GetLoginCode(ip string) string {
	return u.ur.GetLoginCode(ip)
}
