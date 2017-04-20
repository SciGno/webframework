package framework

import "net/http"

// HTTPServer struct
type HTTPServer struct {
	name string
	addr string
	cert string
	key  string
	mux  *http.ServeMux
}

// NewHTTPServer Method
func NewHTTPServer(name string, addr string) HTTPServer {
	return HTTPServer{
		name: name,
		addr: addr,
		cert: "",
		key:  "",
		mux:  http.NewServeMux(),
	}
}

// NewTLSHTTPServer Method
func NewTLSHTTPServer(name string, addr string, cert string, key string) HTTPServer {
	return HTTPServer{
		name: name,
		addr: addr,
		cert: cert,
		key:  key,
		mux:  http.NewServeMux(),
	}
}
