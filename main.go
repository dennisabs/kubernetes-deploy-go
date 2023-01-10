package main

import (
	"fmt"
	"net/http"
	"log"
)

const webContent = "dev-ops-ninja:v120 - mostra fausto"

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webContent)
}
