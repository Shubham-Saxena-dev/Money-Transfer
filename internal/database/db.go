package database

import (
	"gorm.io/gorm"
	"sync"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

type Database interface {
	InitialiseDbConnection()
}

type database struct{}

func NewDatabase() Database {
	return &database{}
}

func (db *database) InitialiseDbConnection() {
	if dbInstance == nil {
		once.Do(
			func() {
				dbInstance = newDatabaseConnection()
			})
	}
}

func GetInstance() *gorm.DB {
	return dbInstance
}
