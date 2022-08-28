package service

import "github.com/mrbelka12000/grpc-gateway/internal/models"

type (
	UserService interface {
		GetLoginUrl(ip string) (string, error)
		SaveLoginCode(ip string, loginResp *models.LoginResp) error
		GetLoginCode(ip string) string
	}
	UserRepo interface {
		SaveLoginCode(ip, code string)
		GetLoginCode(ip string) string
		SaveUsersState(ip, state string)
		CheckStateByIP(ip, state string) bool
	}
	Session interface {
		Get(code string)
		Update()
	}
	TracksRepo interface {
		SaveTracks()
	}
	TracksService interface {
		GetAll() (*models.Tracks, error)
	}
)
