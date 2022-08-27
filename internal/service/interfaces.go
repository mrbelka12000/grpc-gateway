package service

import "github.com/mrbelka12000/grpc-gateway/internal/models"

type (
	UserService interface {
		Login()
		SaveLoginData()
	}
	UserRepo interface {
		SaveLoginData()
	}
	Session interface {
		Get()
		Update()
	}
	TracksRepo interface {
		SaveTracks()
	}
	TracksService interface {
		GetAll() (*models.Tracks, error)
	}
)
