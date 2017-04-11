// This is the main package
package main

import (
	"log"
	"net/http"

	"bitbucket.com/scigno/webframework/config"
)

// This is the main function
func main() {

	config.Init()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
