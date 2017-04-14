package router

import (
	"bytes"
	"fmt"
	"net/http"
)

// Router struct
type Router struct {
	routes map[string]func(w http.ResponseWriter, r *http.Request)
}

// AddHandler function
func (r *Router) AddHandler(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.routes[path] = handler
	http.HandleFunc(path, handler)
}

// AddHandlers function
func (r *Router) AddHandlers(routers map[string]func(w http.ResponseWriter, r *http.Request)) {
	for k, v := range r.routes {
		r.AddHandler(k, v)
	}
}

// String shows router configuration
func (r *Router) String() string {
	var buffer bytes.Buffer
	for k := range r.routes {
		buffer.WriteString(fmt.Sprintf("\tPath: %s\n", k))
	}
	return fmt.Sprintf(buffer.String())
}

// New function
func New() *Router {
	r := Router{
		routes: make(map[string]func(w http.ResponseWriter, r *http.Request)),
	}
	return &r
}

var router = make(map[string]func(w http.ResponseWriter, r *http.Request))

// AddHandler function
func AddHandler(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	router[path] = handler
	http.HandleFunc(path, handler)
}

// AddHandlers function
func AddHandlers(routers map[string]func(w http.ResponseWriter, r *http.Request)) {
	for k, v := range routers {
		AddHandler(k, v)
	}
}

// String shows router configuration
func String() string {
	var buffer bytes.Buffer
	for k := range router {
		buffer.WriteString(fmt.Sprintf("\tPath: %s\n", k))
	}
	return fmt.Sprintf(buffer.String())
}
