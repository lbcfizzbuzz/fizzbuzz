package datastore

import (
	"errors"
	cfg "github.com/samyy321/fizzbuzz/config"
	"github.com/samyy321/fizzbuzz/models"
)

// Datastore represents the methods that has to be implemented by any new datastore system
type Datastore interface {
	FindByMostUsedQueryString() (models.Request, error)
	Init() error
	Store(request *models.Request) error
}

// GetDatastore returns a datastore given the type as parameter.
func GetDatastore(config cfg.Configuration) (Datastore, error) {
	switch config.DbType {
	case "mysql":
		return &MySQLDatastore{Dsn: config.Dsn}, nil
	default:
		return nil, errors.New("This datastore type is not implemented")
	}
}