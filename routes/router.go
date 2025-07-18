package routes

import (
	"blog/handlers"
	"blog/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handlers.UserHandler, postHandler *handlers.PostHandler) *mux.Router {

	r := mux.NewRouter()

	//public routes
	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/login", userHandler.Login).Methods("POST")


	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	
	protected.HandleFunc("/me", userHandler.GetUserInfo).Methods("GET")
	protected.HandleFunc("/posts", postHandler.Posts_URL)
	protected.HandleFunc("/posts/:{id}", postHandler.PostsByID_URL)
	

	return r
}

