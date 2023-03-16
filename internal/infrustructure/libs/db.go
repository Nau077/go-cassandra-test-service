package libs

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbClient struct {
	client *gorm.DB
}

func NewDbClient(dsn string) *dbClient {
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &dbClient{
		client: db,
	}
}
