// Package relations in login
// and more
package relations

import (
	"net/http"

	"github.com/scigno/webframework/logger"
	"github.com/scigno/webframework/templates"
)

type login struct {
	Name string
}

// Login function handler
func Login(w http.ResponseWriter, r *http.Request) {

	// fmt.Println("method:", r.Method) // get request method
	logger.Info(r.Host)
	logger.Info(r.RemoteAddr)
	logger.Info(r.UserAgent())
	logger.Info(r.Method)

	data := login{"Welcome Message"}
	// temp := core.NewTemplateParser()
	templates.Parse("login", data)
	templates.ExecuteTemplate("login", w)

	// t, _ := template.ParseFiles("entities/index/index.html")
	// a := "core.GetUserIdentity()"
	// t.Execute(w, a)

}
