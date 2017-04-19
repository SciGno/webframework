// This is the main package
package main

import (
	"net/http"

	"bitbucket.com/scignox/auth/relations"
	"github.com/SciGno/webframework/framework"
	"github.com/scigno/webframework/logger"
)

// This is the main function
func main() {

	// config.Init()

	wf := framework.New()
	wf.NewServer("server1")

	fs := http.FileServer(http.Dir("static"))
	wf.Handle("server1", "/static/", http.StripPrefix("/static/", fs))
	wf.HandleFunc("server1", "/", relations.Index)

	logger.Info(wf.Servers())
	logger.Info(wf.Routes())

	wf.Start("server1")

	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }

}
