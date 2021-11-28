package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func StartServer() {
	r := NewRouter()
	server := &http.Server{
		Addr:    ":4000",
		Handler: r,
	}
	log.Printf("Server is running on port : %s", server.Addr)
	server.ListenAndServe()
}

func WriteJsonData(w http.ResponseWriter, data interface{}, statusHeader int) {
	//Write out the response
	jsonData, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusHeader)
	w.Write(jsonData)
}
