package infrustructure

import (
	"context"

	"github.com/Nau077/cassandra-golang-sv/internal/infrustructure/libs"
	"github.com/Nau077/cassandra-golang-sv/internal/infrustructure/repos"
	"honnef.co/go/tools/config"
)

type Provider struct {
	staticPath string
	config     *config.Config
	repos      *repos.Container
	db         *libs.NewDbClient
	ctx        context.Context
}

func NewProvider(staticPath string, ctx context.Context) *Provider {
	return &Provider{
		staticPath: staticPath,
		ctx:        ctx,
	}
}
