package main

import (
	"encoding/json"
	"net/http"

	repository "randapi.com/repo"
)

var (
	repo repository.Repository = repository.NewMysqlRepository()
)

func getRoot(resp http.ResponseWriter, _ *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	result, err := repo.GetRow()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error getting language row"`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}
