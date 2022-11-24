package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Law struct {
	Name       string `json:"name,omitempty"`
	Definition string `json:"definition,omitempty"`
}

var HackerLaws = []Law{
	{
		Name:       "Amdahl's Law",
		Definition: "Amdahl's Law is a formula which shows the potential speedup of a computational task which can be achieved by increasing the resources of a system.",
	},
}

func GetRandomLaw(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(HackerLaws[0])
	if err != nil {
		http.Error(w, "couldn't retrieve random hacker law", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, bytes.NewReader(j))
}
