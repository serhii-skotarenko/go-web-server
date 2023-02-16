package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Server received a request to %s \n", r.URL)

	w.Header().Add("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["status"] = "OK"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	fmt.Printf("Server respond with: %s \n", jsonResp)
	w.Write(jsonResp)
	return
}

func main() {
	http.HandleFunc("/", getRoot)

	fmt.Printf("🚀Server is up and running! \n")
	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
