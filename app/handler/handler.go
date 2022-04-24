package handler

import (
	"encoding/json"
	"main-1/models"
	"net/http"
)

func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := &models.Response{Status: "OK"}
		resp, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(resp)
	}
}
