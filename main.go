package main

import (
	"assignment-3/functions"
	"fmt"
	"net/http"
)

func main() {
	go functions.JsonReload()

	http.HandleFunc("/", functions.WebReload)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	fmt.Println("listening on PORT:", ":9000")
	http.ListenAndServe(":9000", nil)
}
