package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type GreetRequest struct {
	Name string `json:name`
}

type GreetResponse struct {
	Message string `json:message`
}

func main() {
	log.Println("実行")

	http.HandleFunc(
		"/",
		func(
			w http.ResponseWriter,
			r *http.Request) {
			fmt.Fprintln(w, "HelloWorld!")
		})

	http.HandleFunc(
		"/json",
		func(
			w http.ResponseWriter,
			r *http.Request) {
			res := GreetResponse{Message: "HelloGopher."}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(res); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})

	addr := ":8080"
	log.Printf("サーバー起動:http://localhost%s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
