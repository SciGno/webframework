package templates

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

// var templates = make(map[string]interface{})

// TemplateContext type
type TemplateContext string

// TemplateData struct
type TemplateData struct {
	Code int
	Data interface{}
}

// Init function
// func Init(tpl string, data interface{}) {
// 	templates[tpl] = data
// }

// ExecuteTemplate function
func ExecuteTemplate(data interface{}, tpl string, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	local := TemplateData{
		Data: data,
	}
	// data := templates[tpl]

	// TODO: Get other template data from context

	pattern := strings.Join([]string{"entities", tpl, "*.html"}, "/")
	t := template.Must(template.ParseGlob(pattern))
	pattern = strings.Join([]string{"entities", "common", "*.html"}, "/")
	t = template.Must(t.ParseGlob(pattern))

	// Get the status message if available
	code := ctx.Value(TemplateContext("http_status"))

	if code != nil {
		local.Code = code.(int)
	} else {
		local.Code = 200
	}

	err := t.Execute(w, local)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}
