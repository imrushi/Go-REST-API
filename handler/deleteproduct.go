package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ErrorStruct struct {
	Message string `json:"errMsg"`
}

func DeleteProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// prod := Product{}

	// err = prod.FromJSON(r.Body)
	// if err != nil {
	// 	http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	// }
	// prod := r.Context().Value(KeyProduct{}).(Product)

	err = DeleteProduct(id)
	if err == ErrProductNotFound {
		fmt.Printf("[ERROR] deleting record id does not exist")
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		fmt.Printf("[ERROR] deleting record", err)
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
