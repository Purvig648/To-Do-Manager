package main

import (
	"github.com/todo_manager/pkg/database"
	"github.com/todo_manager/pkg/handler"
	"github.com/todo_manager/pkg/repository"
	"github.com/todo_manager/pkg/server"
	"github.com/todo_manager/pkg/service"
)

func main() {
	db, err := database.ConnectToDatabase()
	if err != nil {
		return
	}
	repo := repository.NewRepoLayer(db)

	svc := service.NewServiceLayer(repo)

	handler := handler.NewHandlerLayer(svc)

	server.StartApp(handler)

}
