package service

import "github.com/mrbelka12000/grpc-gateway/internal/models"

type Tracks struct {
	tr TracksRepo
}

func newTracks(tr TracksRepo) *Tracks {
	return &Tracks{
		tr: tr,
	}
}

func (t *Tracks) GetAll() (*models.Tracks, error) {
	return nil, nil
}
