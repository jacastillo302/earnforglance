package repository

import (
	"context"

	domain "earnforglance/server/domain/install"
	"earnforglance/server/service/data/mongo"
)

type intallRepository struct {
	database mongo.Database
}

// NewInstallRepository creates a new instance of intallRepository
func NewInstallRepository(db mongo.Database) domain.InstallRepository {
	return &intallRepository{
		database: db,
	}
}

func (ur *intallRepository) PingDatabase(c context.Context) error {
	client := ur.database.Client()
	err := client.Ping(context.Background())
	return err
}
