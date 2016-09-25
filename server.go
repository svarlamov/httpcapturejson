package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type JSON map[string]interface{}


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got an HTTP", r.Method, "Request on", r.URL.Path)
	fmt.Println("Content-Type:", r.Header.Get("Content-Type"))
	decoder := json.NewDecoder(r.Body)
	var reqObj JSON
	err := decoder.Decode(&reqObj)
	if err != nil {
		fmt.Println("JSON Error:", err)
		return
	}
	fmt.Println("JSON Body:", reqObj)
	fmt.Fprintf(w, "{\"status\":\"ok\"")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting HTTP Server")
	http.ListenAndServe(":3000", nil)
}