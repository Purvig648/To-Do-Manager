package database

import (
	"log"

	dbmodel "github.com/todo_manager/pkg/model/db_model"
	"github.com/todo_manager/pkg/util"
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

	PopulateUserTableColumnMap()
	PopulateTaskStatusColumnMMap()

	db.AutoMigrate(&dbmodel.User{}, &dbmodel.Task{})
	return db, nil
}

func PopulateUserTableColumnMap() {
	util.Choices["username"] = "username"
	util.Choices["email_id"] = "email_id"
}

func PopulateTaskStatusColumnMMap() {
	util.TaskChoices["In Progres"] = "In Progress"
	util.TaskChoices["Completed"] = "Completed"
	util.TaskChoices["Skip"] = "Skip"
}
