package routes

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type CheckResponse struct {
	Success bool `json:"success"`
}

func TestApiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	test := make(map[string]bool)
	test["success"] = true

	j, _ := json.Marshal(test)
	w.Write(j)
}

func CheckApiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	successChance := rand.Intn(101)
	var checkResponse bool

	if successChance <= 80 {
		checkResponse = true
	} else {
		checkResponse = false
	}

	response := &CheckResponse{
		Success: checkResponse,
	}

	j, _ := json.Marshal(response)
	w.Write(j)
}
