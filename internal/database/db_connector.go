package database

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	mysqlgorm "gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"qonto/internal/config"
	"qonto/pkg/models"
)

func newDatabaseConnection() *gorm.DB {
	var err error
	conn, err := getConnection()

	if err != nil {
		log.Fatal("failed to connect database")
		panic(err)
	}
	log.Info("Database connected successfully")
	err = conn.AutoMigrate(&models.BankAccount{}, &models.Transfer{})
	if err != nil {
		log.Fatal("failed to migrate models")
		panic(err)
	}
	if config.EnvConfigs.App.Environment == "dev" {
		loadDummyData(conn)
	}
	return conn
}

func getConnection() (*gorm.DB, error) {
	fmt.Println(config.EnvConfigs.Db.DatabaseUri)
	if config.EnvConfigs.Db.Driver == "mysql" {
		return gorm.Open(mysqlgorm.Open(config.EnvConfigs.Db.DatabaseUri), &gorm.Config{})
	}
	return gorm.Open(sqlite.Open(config.EnvConfigs.Db.DatabaseUri), &gorm.Config{})
}

func loadDummyData(conn *gorm.DB) {
	conn.Exec("INSERT INTO bank_accounts VALUES (1, 'ACME Corp', 10000000, 'FR10474608000002006107XXXXX', 'OIVUSCLQXXX')")
	conn.Exec("INSERT INTO transfers VALUES (1, 'ACME Corp. Main Account', 'EE382200221020145685', 'CCOPFRPPXXX', 11000000, 1, 'Treasury management'),(2, 'Bip Bip', 'EE383680981021245685', 'CRLYFRPPTOU', 1000000, 1, 'Bip Bip Salary')")
}
