package main

import "net/http"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error

	switch r.Method {
	case "GET":
		err = handleGetCountryInfo(w, r)

	case "POST":
		err = handlePostCountryInfo(w, r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
