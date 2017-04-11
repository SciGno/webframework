// Package relations Index relation
// and more
package relations

import (
	"net/http"

	"bitbucket.com/scigno/webframework/auth"
	"bitbucket.com/scigno/webframework/templates"
)

type data struct {
	Data map[string]interface{}
}

var local = data{make(map[string]interface{})}

// Index handler.
func Index(w http.ResponseWriter, r *http.Request) {

	// fmt.Println("method:", r.Method) // get request method

	local.Data["local"] = "Welcome Message"
	local.Data["number"] = 1
	local.Data["User"] = auth.User()
	local.Data["Auth"] = auth.Auth()

	templates.Parse("index", local)
	templates.ExecuteTemplate("index", w)
}
