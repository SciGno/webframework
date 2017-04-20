// This is the main package
package main

import (
	"net/http"

	"github.com/SciGno/webframework/framework"
	"github.com/scigno/webframework/examples/semantic_ui/relations"
)

// This is the main function
func main() {

	// config.Init()

	wf := framework.New()
	wf.NewServer("server1", ":8080")

	fs := http.FileServer(http.Dir("static"))
	wf.Handle("server1", "/static/", http.StripPrefix("/static/", fs))
	wf.HandleFunc("server1", "/", relations.Index)

	wf.Start("server1")

}
