// Package relations Index relation
// and more
package relations

import (
	"net/http"

	"github.com/scigno/webframework/templates"
)

type data struct {
	Data map[string]interface{}
}

var local = data{make(map[string]interface{})}

// Index handler.
func Index(w http.ResponseWriter, r *http.Request) {

	// for k, v := range r.Header {
	// 	logger.Info("%v : %v", k, v)
	// }
	// logger.Info("------- INDEX")

	// fmt.Println("method:", r.Method) // get request method

	local.Data["Message"] = "Welcome Message"
	local.Data["User"] = "Someone"
	local.Data["Auth"] = "Authorized"

	// templates.Init("index", local)
	templates.ExecuteTemplate(local, "index", w, r)
}
