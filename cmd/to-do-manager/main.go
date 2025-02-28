package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/todo_manager/pkg/auth"
	"github.com/todo_manager/pkg/database"
	"github.com/todo_manager/pkg/handler"
	"github.com/todo_manager/pkg/middleware"
	"github.com/todo_manager/pkg/repository"
	"github.com/todo_manager/pkg/server"
	"github.com/todo_manager/pkg/service"
)

func main() {
	err := godotenv.Load("pkg/env/auth.env")
	if err != nil {
		return
	}
	a := auth.NewAuth(os.Getenv("Secret_Key"))

	mid := middleware.NewMiddleware(a)
	db, err := database.ConnectToDatabase()
	if err != nil {
		return
	}

	repo := repository.NewRepoLayer(db)

	svc := service.NewServiceLayer(repo, a)

	handler := handler.NewHandlerLayer(svc)

	server.StartApp(handler, mid)

}
