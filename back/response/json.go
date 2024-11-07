package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data interface{}) {

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		http.Error(w, "Error al procesar la respuesta JSON", http.StatusInternalServerError)
	}
}
