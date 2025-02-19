package database

import (
	"log"

	dbmodel "github.com/todo_manager/pkg/model/db_model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=Purvi@123# dbname=todo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = sqlDb.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	db.AutoMigrate(&dbmodel.User{})
	return db, nil
}
