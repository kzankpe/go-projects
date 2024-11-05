package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kzankpe/unit-converter/internal"
)

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)

	http.HandleFunc("/form", internal.FormHandler)

	fmt.Println("Server is listening on port 8090")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
