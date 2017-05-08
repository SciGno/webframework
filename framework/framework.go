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
	servers map[string]HTTPServer
}

// New function
func New() Framework {
	return Framework{
		servers: make(map[string]HTTPServer),
	}
}

// NewServer Method
func (f *Framework) NewServer(server string, addr string) error {
	f.servers[server] = NewHTTPServer(server, addr)
	return nil
}

// NewTLSServer Method
func (f *Framework) NewTLSServer(server string, addr string, cert string, key string) error {
	f.servers[server] = NewTLSHTTPServer(server, addr, cert, key)
	return nil
}

// Servers Method
func (f *Framework) Servers() string {
	var buffer bytes.Buffer
	for k := range f.servers {
		buffer.WriteString(fmt.Sprintf("%s\n", k))
	}
	return fmt.Sprintf(buffer.String())
}

// Handle Method
func (f *Framework) Handle(server string, pattern string, handler http.Handler) {
	f.servers[server].mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		requestLogger(r)
		handler.ServeHTTP(w, r)
	})
}

// HandleFunc Method
func (f *Framework) HandleFunc(server string, pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	f.servers[server].mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		requestLogger(r)
		handler(w, r)
	})
}

// HandleGETFunc Method
func (f *Framework) HandleGETFunc(server string, pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	f.servers[server].mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			requestLogger(r)
			handler(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
}

// Start Method
func (f *Framework) Start(server string) {
	err := http.ListenAndServe(f.servers[server].addr, f.servers[server].mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// requestLogger function
func requestLogger(r *http.Request) {
	logger.Info("| %s | %s | %s ", r.Method, r.RemoteAddr, r.RequestURI)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}
