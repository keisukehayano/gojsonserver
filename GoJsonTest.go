package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func handleIndexRequest(w http.ResponseWriter, r *http.Request) (err error) {

	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")
	pass := r.URL.Query().Get("pass")

	fmt.Println("name:", name)
	fmt.Println("email:", email)
	fmt.Println("pass:", pass)

	user := User{1, name, email, pass}

	//構造体をJsonに変換
	res, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Header
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//ステータスコード 200
	w.WriteHeader(http.StatusOK)

	// Response
	w.Write(res)

	return
}
