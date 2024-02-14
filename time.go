package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func serverHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	if request.Method != http.MethodGet {
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	currentTime := getCurrentTime()
	resp := map[string]string{"time": currentTime}
	ResponseInJson, err := json.Marshal(resp)
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	writer.Write(ResponseInJson)
}

func getCurrentTime() string {
	currentTime := time.Now().Format(time.RFC3339)
	return currentTime
}

func main() {
	http.HandleFunc("/time", serverHandler)

	port := 8795
	println("Server is running on port:", fmt.Sprint(port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
