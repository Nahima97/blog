package handlers

import (
	"blog/config"
	"blog/middleware"
	"blog/models"
	"blog/services"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type PostHandler struct {
	Service *services.PostService
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var req models.Post
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userRole, err := middleware.GetUserRole(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	permission := config.RolePermission(userRole, "create:post")
	if !permission {
		http.Error(w, "access denied", http.StatusUnauthorized)
		return
	}

	userID, err := middleware.ExtractUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	user_uuid, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.AuthorID = user_uuid

	err = h.Service.CreatePost(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}

func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.Service.GetAllPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}


func (h *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	post, err := h.Service.GetPostByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}


func (h *PostHandler) UpdateOwnPost(w http.ResponseWriter, r *http.Request) {
	var req models.Post
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := mux.Vars(r)["id"]

	userID, err := middleware.ExtractUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	post, err := h.Service.GetPostByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	user_uuid, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	
	if user_uuid != post.AuthorID {
		http.Error(w, "cannot update this post as you are not the author", http.StatusUnauthorized)
		return
	}
	
	post.Title = req.Title
	post.Content = req.Content

	err = h.Service.UpdateOwnPost(post, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(req)
}


func (h *PostHandler) DeleteOwnPost(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

	userID, err := middleware.ExtractUserID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	post, err := h.Service.GetPostByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return 
	}

	user_uuid, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if user_uuid != post.AuthorID {
		http.Error(w, "cannot delete this post as you are not the author", http.StatusUnauthorized)
		return
	}

	err = h.Service.DeleteOwnPost(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}


func (h *PostHandler) PostsByID_URL(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		h.GetPostByID(w, r)
	case http.MethodPut:
		h.UpdateOwnPost(w, r)
	case http.MethodDelete:
		h.DeleteOwnPost(w, r)
	}
}

func (h *PostHandler) Posts_URL(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		h.GetAllPosts(w, r)
	case http.MethodPost:
		h.CreatePost(w, r)
	}
}
