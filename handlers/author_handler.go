package handlers

import (
	"bookstore/models"
	"encoding/json"
	"net/http"
)

var nextIDAuthors = 1
var userList []models.Author

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//for _, user := range userList {
	//	userList = append(userList, user)
	//}
	json.NewEncoder(w).Encode(userList)
}

func AddAuthor(w http.ResponseWriter, r *http.Request) {
	var user models.Author

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.ID = nextIDAuthors
	nextIDAuthors++
	userList = append(userList, user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
