package repository

import (
	"fmt"
	config2 "github.com/tttinh/go-rest-api-with-gin/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	Group GroupRepository
}

var db *gorm.DB

func New(dbConfig *config2.DatabaseConfig) Repository {
	initDB(dbConfig)
	return Repository{
		Group: makeGroupRepository(),
	}
}

// initDB initializes the database instance
func initDB(dbConfig *config2.DatabaseConfig) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Name,
	)

	mysqlDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to init database, err: %v", err)
	}

	db = mysqlDb
}
