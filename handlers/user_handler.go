package handlers

import (
	"blog/middleware"
	"blog/models"
	"blog/services"
	"encoding/json"
	"net/http"

)

type UserHandler struct {
	Service *services.UserService
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req models.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Service.RegisterUser(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)

}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.Service.Login(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

func (h *UserHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {

userID, err := middleware.ExtractUserID(r)
if err != nil {
	http.Error(w, err.Error(), http.StatusUnauthorized)
	return 
}

user, err := h.Service.GetUserInfo(userID)
if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return 
}

w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(user)
}

