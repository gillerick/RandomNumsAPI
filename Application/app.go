package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main(){
	//Local http server and routes
	http.HandleFunc("/random/", randomGenerator)
	http.ListenAndServe(port(), nil)
	log.Printf("Defaulting to port %s", port())
}

func port() string{
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" +port
}

//The function that randomizes the numbers using a varying seed - the current time
func randomize() int {
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	//Range: 20 - 20000, Max: 19999, Min: 21
	return y1.Intn(19999-21) + 21
}

//The response handler
func randomGenerator(w http.ResponseWriter, r *http.Request)  {
	switch method := r.Method; method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%d", randomize())
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([] byte("Method not supported"))

	}

}


