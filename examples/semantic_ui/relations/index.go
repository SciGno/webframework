// Package relations Index relation
// and more
package relations

import (
	"net/http"

	"github.com/scigno/webframework/auth"
	"github.com/scigno/webframework/templates"
)

type indexData struct {
	Code  int
	Error int
	Data  map[string]interface{}
}

// Index handler.
func Index(w http.ResponseWriter, r *http.Request) {

	// for k, v := range r.Header {
	// 	logger.Info("%v : %v", k, v)
	// }
	// logger.Info("------- INDEX")

	// fmt.Println("method:", r.Method) // get request method

	d := indexData{}
	d.Data = make(map[string]interface{})
	d.Code = 200

	if !auth.JWTContextValue(auth.JWTTokenValid, r).(bool) {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

	d.Data["Message"] = "Welcome Message"
	d.Data["User"] = "Someone"
	d.Data["Auth"] = "Authorized"

	// templates.Init("index", local)
	templates.ExecuteTemplate(d, "index", w, r)
}
