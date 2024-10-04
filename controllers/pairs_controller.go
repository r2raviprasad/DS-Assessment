package controllers

import (
	"encoding/json"
	"net/http"
)

// we are creating structs inside controller not in models because we just want to bind request into struct
type FindPairsRequest struct {
	Numbers []int `json:"numbers" validate:"required,dive,number"`
	Target  int   `json:"target" validate:"required"`
}
type FindPairsResponse struct {
	Solutions [][]int `json:"solutions"`
}

func FindPairs(w http.ResponseWriter, r *http.Request) {
	var req FindPairsRequest

	err := json.NewDecoder(r.Body).Decode(&req) // binds request to FindPairsRequest
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = validate.Struct(req) // validate request body parameters or give 400 err code
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	solutions := FindPairsSolution(req.Numbers, req.Target)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(FindPairsResponse{Solutions: solutions})
}

func FindPairsSolution(numbers []int, target int) [][]int {
	pairMap := make(map[int]int)
	var solutions [][]int

	for i, num := range numbers {
		complement := target - num
		if index, found := pairMap[complement]; found {
			solutions = append(solutions, []int{index, i}) // will append new array/slice to solutions array/slice
		}
		pairMap[num] = i
	}

	return solutions
}
