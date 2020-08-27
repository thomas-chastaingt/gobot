package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func HomeEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "hello")
}

func VerificationEndpoint(w http.ResponseWriter, r *http.Request) {
	challenge := r.URL.Query().Get("hub.challenge")
	token := r.URL.Query().Get("hub.verifiy_token")

	if token == os.Getenv("VERIFY_TOKEN") {
		w.WriteHeader(200)
		w.Write([]byte(challenge))
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Error, wrong validation token"))
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeEndpoint)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
