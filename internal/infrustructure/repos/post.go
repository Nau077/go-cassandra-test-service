package repos

import (
	"context"

	model "github.com/Nau077/cassandra-golang-sv/internal/infrustructure/models"
	"github.com/Nau077/cassandra-golang-sv/internal/presentation/model_adapter"
)

// добавить dbClient внутрь репозитория Post
type PostRepo struct {
	db DbClient.Client
}

type PostSender interface {
	AddPost(ctx context.Context, post model_adapter.Post) (int64, error)
	GetPost(ctx context.Context, id int64) (*model.Post, error)
}

func (p *PostRepo) AddPost(ctx context.Context, post model_adapter.Post) (int64, error) {
	result := db.Create(&post)

	// user.ID
	// result.Error
	// result.RowsAffected
	return result
}

func (p *PostRepo) GetPost(ctx context.Context, id int64) (*model.Post, error) {

}
