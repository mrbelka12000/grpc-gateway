package repo

type Repo struct {
	*User
	*Tracks
}

func NewRepo() *Repo {
	return &Repo{
		User:   newUser(),
		Tracks: newTracks(),
	}
}
