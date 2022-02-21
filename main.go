package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

type body struct {
	String string `json:"string"`
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	reqBody := body{}
	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rev, err := Reverse(reqBody.String)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	respBody := body{String: rev}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respBody)
}

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server started")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
