// This is the main package
package main

import (
	"log"
	"net/http"

	"github.com/SciGno/webframework/httprouter"
	"github.com/scigno/webframework/examples/semantic_ui/relations"
)

// This is the main function
func main() {

	s := httprouter.New()
	fs := http.FileServer(http.Dir("static"))
	s.HandleStatic("/static", http.StripPrefix("/static", fs))
	s.HandleFuncGET("/", relations.Index)

	log.Fatal(http.ListenAndServe(":8080", s))

}
