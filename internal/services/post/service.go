package post

import "github.com/Nau077/cassandra-golang-sv/internal/infrustructure/repos"

type Service struct {
	repos *repos.Container
}

func NewService(repos *repos.Container) *Service {
	return &Service{
		repos: repos,
	}
}
