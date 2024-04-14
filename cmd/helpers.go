package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func TranslateSubject(v *Student) string {
	gen := fmt.Sprintf("Name: %s, Goal: %v\n", v.Name, v.Goal)
	s := ""
	for k, v := range v.Materials {
		t := fmt.Sprintf("Subject: %v, Grade: %v\n", k, v.Grade)
		s += t
	}
	return gen + s
}
