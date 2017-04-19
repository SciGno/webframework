package framework

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/scigno/webframework/logger"
)

// Framework struct
type Framework struct {
	mux map[string]*http.ServeMux
}

var routes = make(map[string][]string)

// New function
func New() Framework {
	f := Framework{
		mux: make(map[string]*http.ServeMux),
	}
	return f
}

// NewServer function
func (f *Framework) NewServer(server string) error {
	f.mux[server] = http.NewServeMux()
	return nil
}

// Servers function
func (f *Framework) Servers() string {
	var buffer bytes.Buffer
	for k := range f.mux {
		buffer.WriteString(fmt.Sprintf("%s\n", k))
	}
	return fmt.Sprintf(buffer.String())
}

// Routes function
func (f *Framework) Routes() string {
	var buffer bytes.Buffer
	for k := range f.mux {
		buffer.WriteString(fmt.Sprintf("%s: %s\n", k, routes[k]))
	}
	return fmt.Sprintf(buffer.String())
}

// Handle function
func (f *Framework) Handle(server string, path string, handler http.Handler) {
	routes[server] = append(routes[server], path)
	f.mux[server].HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		requestLogger(r)
		handler.ServeHTTP(w, r)
	})
}

// HandleFunc function
func (f *Framework) HandleFunc(server string, path string, handler func(w http.ResponseWriter, r *http.Request)) {
	routes[server] = append(routes[server], path)
	f.mux[server].HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		requestLogger(r)
		handler(w, r)
	})
}

// requestLogger function
func requestLogger(r *http.Request) {
	logger.Info("=== REQUEST ===")
	logger.Info("UserAgent: %s", r.UserAgent())
	logger.Info("Referer: %s", r.Referer())
	logger.Info("Host: %s", r.Host)
	logger.Info("Method: %s", r.Method)
	logger.Info("RemoteAddr: %s", r.RemoteAddr)
	logger.Info("RequestURI: %s", r.RequestURI)
	logger.Info("URL.EscapedPath(): %s", r.URL.EscapedPath())
	for k, v := range r.Header {
		logger.Info("%s: %s", k, v)
	}
}

// Start function
func (f *Framework) Start(server string) {
	err := http.ListenAndServe(":8080", f.mux[server])
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
