package templates

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

var templates = make(map[string]interface{})

// Parse function
func Parse(tpl string, data interface{}) {
	templates[tpl] = data
}

// ExecuteTemplate function
func ExecuteTemplate(tpl string, w http.ResponseWriter) {

	data := templates[tpl]

	pattern := strings.Join([]string{"entities", tpl, "*.html"}, "/")
	t := template.Must(template.ParseGlob(pattern))
	pattern = strings.Join([]string{"entities", "common", "*.html"}, "/")
	t = template.Must(t.ParseGlob(pattern))

	err := t.Execute(w, data)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

}
