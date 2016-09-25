package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type JSON map[string]interface{}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().String(), ":", "Got an HTTP", r.Method, "Request on", r.URL.Path)
	fmt.Println(time.Now().String(), ":", "Content-Type:", r.Header.Get("Content-Type"))
	decoder := json.NewDecoder(r.Body)
	var reqObj JSON
	err := decoder.Decode(&reqObj)
	if err != nil {
		fmt.Println(time.Now().String(), ":", "JSON Error:", err)
		return
	}
	fmt.Println(time.Now().String(), ":", "JSON Body:", reqObj)
	fmt.Fprintf(w, "{\"status\":\"ok\"")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println(time.Now().String(), ":", "Starting HTTP Server")
	http.ListenAndServe(":3000", nil)
}
