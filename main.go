package main

import (
	"blog/db"
	"blog/handlers"
	"blog/repository"
	"blog/routes"
	"blog/services"
	"fmt"
	"log"
	"net/http"
)

func main() {

	Db := db.InitDb()

	userRepo := &repository.UserRepo{Db: Db}
	postRepo := &repository.PostRepo{Db: Db}

	userService := &services.UserService{Repo: userRepo}
	postService := &services.PostService{Repo: postRepo}

	userHandler := &handlers.UserHandler{Service: userService}
	postHandler := &handlers.PostHandler{Service: postService}

	r := routes.SetupRouter(userHandler, postHandler)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("failed to start server", err)
	}
	fmt.Println("server started")
}
