package framework

import "net/http"

// HandleGETFunc Method
func (f *Framework) HandleGETFunc(server string, pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	f.servers[server].mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			requestLogger(w, r)
			handler(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
}

// HandlePOSTFunc Method
func (f *Framework) HandlePOSTFunc(server string, pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	f.servers[server].mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			requestLogger(w, r)
			handler(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
}

// HandlePUTFunc Method
func (f *Framework) HandlePUTFunc(server string, pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	f.servers[server].mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			requestLogger(w, r)
			handler(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
}

// HandleDELETEFunc Method
func (f *Framework) HandleDELETEFunc(server string, pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	f.servers[server].mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			requestLogger(w, r)
			handler(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
}

// HandleHEADFunc Method
func (f *Framework) HandleHEADFunc(server string, pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	f.servers[server].mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "HEAD" {
			requestLogger(w, r)
			handler(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
}
