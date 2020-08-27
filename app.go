package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

func HandleMessenger(resp http.ResponseWriter, request *http.Request) {
	secretKey := "secret_token"
	if request.Method == "GET" {
		u, _ := url.Parse(request.RequestURI)
		values, _ := url.ParseQuery(u.RawQuery)
		token := values.Get("hub.verify_token")
		if token == secretKey {
			resp.WriteHeader(200)
			resp.Write([]byte(values.Get("hub.challenge")))
			return
		}
		resp.WriteHeader(400)
		resp.Write([]byte(`Bad token`))
		return
	}
}

// Initialize request
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HandleMessenger).Methods("POST", "GET")
	port := ":8000"
	log.Printf("Server started on %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
