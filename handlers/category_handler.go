package handlers

import (
	"bookstore/models"
	"encoding/json"
	"net/http"
)

var category []models.Category
var nextIDCategory = 1

func GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//for _, cate := range category {
	//	category = append(category, cate)
	//}
	json.NewEncoder(w).Encode(category)
}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	var categ models.Category
	if err := json.NewDecoder(r.Body).Decode(&categ); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	categ.ID = nextIDCategory
	nextIDCategory++
	category = append(category, categ)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(categ)
}
