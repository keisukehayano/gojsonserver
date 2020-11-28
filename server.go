package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// root ただのjsonを返す
	//http.HandleFunc("/", handleIndexRequest)

	http.HandleFunc("/", handleRequest)

	// サーバー起動
	server.ListenAndServe()
}
