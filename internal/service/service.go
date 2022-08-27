package service

type Service struct {
	*User
	*Tracks
}

func New(ur UserRepo, tr TracksRepo) *Service {
	return &Service{
		User:   newUser(ur),
		Tracks: newTracks(tr),
	}
}
