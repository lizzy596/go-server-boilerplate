package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Message struct {
	Message string
}

type Response struct {
	Status string `json:"status"`
}

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)

		if err != nil {
			log.Fatal("ERROR: ", err)
		}

		fmt.Println(string(body))
		var mes Message
		json.Unmarshal([]byte(string(body)), &mes)

		fmt.Println("Recieved Message: ", mes.Message)

		response := Response{
			Status: "created",
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(response)
	})

	print("Starting server at port 7777\n")

	err := http.ListenAndServe(":7777", nil)

	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}
