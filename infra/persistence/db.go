package persistence

import (
	"fmt"
	"github.com/tttinh/go-rest-api-with-gin/infra/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

// NewDB initializes the database instance
func NewDB(appSetting config.Config) *gorm.DB {
	dbConfig := appSetting.Database
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Name,
	)

	gormCfg := &gorm.Config{}
	if appSetting.Server.Mode == "release" {
		gormCfg.Logger = logger.Default.LogMode(logger.Silent)
	}

	db, err := gorm.Open(mysql.Open(dsn), gormCfg)

	if err != nil {
		log.Fatalf("Failed to init database, err: %v", err)
	}

	return db
}
