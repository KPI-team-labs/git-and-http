package main

import (
	"fmt"
	"net/http"
)

func serverHandler(writer http.ResponseWriter, response *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte("time"))
}

func main() {
	http.HandleFunc("/time", serverHandler)

	port := 8000
	println("Server is running on port:", fmt.Sprint(port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
