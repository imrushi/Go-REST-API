package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type KeyProduct struct{}

func UpdateProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	prod := Product{}

	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}
	// prod := r.Context().Value(KeyProduct{}).(Product)

	err = UpdateProduct(id, &prod)
	if err == ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}
