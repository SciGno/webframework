// Package relations Index relation
// and more
package relations

import (
	"net/http"

	"bitbucket.com/scigno/webframework/templates"
)

type api struct {
	Version string
}

// API This is the API function for the handler.
func API(w http.ResponseWriter, r *http.Request) {

	// fmt.Println("method:", r.Method) // get request method

	data := api{"API v1.0"}
	// temp := core.NewTemplateParser()
	templates.Parse("api", data)
	templates.ExecuteTemplate("api", w)

	// t, _ := template.ParseFiles("entities/index/index.html")
	// a := "core.GetUserIdentity()"
	// t.Execute(w, a)

}
