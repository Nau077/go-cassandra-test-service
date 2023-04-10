package libs

import (
	model "github.com/Nau077/cassandra-golang-sv/internal/infrustructure/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbClient struct {
	Client *gorm.DB
	dsn    string
}

func NewDbClient(dsn string) *DbClient {
	return &DbClient{
		dsn: dsn,
	}
}

func (d *DbClient) Init() error {
	db, err := gorm.Open(postgres.Open(d.dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&model.Post{})

	d.Client = db

	if err != nil {
		return err
	}

	return nil
}
