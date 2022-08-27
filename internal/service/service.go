package service

type Service struct {
	UserService
	TracksService
}

func New(ur UserRepo, tr TracksRepo) *Service {
	return &Service{
		UserService:   newUser(ur),
		TracksService: newTracks(tr),
	}
}
