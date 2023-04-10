package infrustructure

import (
	"context"
	"log"

	"github.com/Nau077/cassandra-golang-sv/internal/infrustructure/config"
	"github.com/Nau077/cassandra-golang-sv/internal/infrustructure/libs"
	"github.com/Nau077/cassandra-golang-sv/internal/infrustructure/repos"
	"github.com/Nau077/cassandra-golang-sv/internal/services/post"
)

type Provider struct {
	staticPath string
	config     *config.Config
	repos      *repos.Container
	db         *libs.DbClient
	router     *libs.RouterClient
	ctx        context.Context
	service    *post.Service
}

func NewProvider(ctx context.Context, staticPath string) *Provider {
	return &Provider{
		staticPath: staticPath,
		ctx:        ctx,
	}
}

func (p *Provider) GetConfig() *config.Config {
	if p.config == nil {
		config, err := config.NewConfig(p.staticPath)
		if err != nil {
			log.Fatalf("failed to get config: %s", err.Error())
		}

		p.config = config
	}

	return p.config
}

func (p *Provider) GetDb() *libs.DbClient {
	if p.db == nil {
		db := libs.NewDbClient(p.config.DB.DSN)

		p.db = db
	}

	return p.db
}

func (p *Provider) GetRepos() *repos.Container {
	if p.repos == nil {
		r := repos.NewContainer()

		p.repos = r
	}

	return p.repos
}

func (p *Provider) GetService() *post.Service {
	if p.service == nil {
		s := post.NewService(p.GetRepos())

		p.service = s
	}

	return p.service
}

func (p *Provider) GetRouter() *libs.RouterClient {
	if p.router == nil {
		r := libs.NewRouterClient(p.config.HTTP.Port, p.GetService())

		p.router = r
	}

	return p.router
}
