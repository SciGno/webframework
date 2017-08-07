package framework

import (
	"context"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/scigno/webframework/logger"
)

// ContextKey type
type ContextKey string

var matches, _ = regexp.Compile("\\{(.*?)\\}")

// HandlerWrapper struct converts a function into an http.Handler
type HandlerWrapper struct {
	handler http.Handler
}

// ServeHTTP method
func (h HandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}

// RegExHandlerWrapper struct converts a function into an http.Handler
type RegExHandlerWrapper struct {
	handler http.Handler
}

// ServeHTTP method
func (h RegExHandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}

// FuncHandlerWrapper struct converts a function into an http.Handler
type FuncHandlerWrapper struct {
	handler func(w http.ResponseWriter, r *http.Request)
}

// ServeHTTP method
func (h FuncHandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler(w, r)
}

// RegExFuncHandlerWrapper struct converts a function into an http.Handler
// and parses path named parameters
type RegExFuncHandlerWrapper struct {
	handler    func(w http.ResponseWriter, r *http.Request)
	muxHandler *MuxHandler
}

// ServeHTTP method
func (h RegExFuncHandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// path := r.URL.Path
	// setRequestValues(path)
	// logger.Info("[RegExFuncHandlerWrapper] Request: %+v", r.Context())
	sm := r.Context().Value(http.ServerContextKey)

	logger.Info("[RegExFuncHandlerWrapper]: %+v", sm.(*http.Server).Handler)
	p := strings.Split(r.URL.Path, "/")
	userid := p[len(p)-1:][0]
	contextKeyAuthtoken := "auth-token"
	contextKeyAnother := userid
	ctx := r.Context()
	ctx = context.WithValue(ctx, contextKeyAuthtoken, contextKeyAnother)
	h.handler(w, r.WithContext(ctx))
}

// func setRequestValues(p string) {

// 	for k, v := range MuxHandler.path {
// 		// matchRegEx, err := regexp.Compile(k)
// 		if matchRegEx, err := regexp.Compile(k); matchRegEx.Match([]byte(path)) && err == nil {
// 			if _, ok := v[r.Method]; ok {
// 				(m.path[k][r.Method]).ServeHTTP(w, r)
// 				// return
// 				// logger.Info("[MuxHandler] Full match!!! %v", m.path[k][r.Method])
// 				return
// 			}
// 		}
// 		// logger.Info("[MuxHandler] Path Matches: %v", matchRegEx.Match([]byte(path)))
// 		// logger.Info("[MuxHandler] RegExPatter: %v", k)
// 		// logger.Info("[MuxHandler] Value: %v", v)
// 	}

// 	logger.Info("regexPattern2 Pattern: %v", p)
// 	regexPattern2 := buildRegEx.Split(p, -1)
// 	logger.Info("regexPattern2 Pattern: %v", regexPattern2)
// 	// splitValsRegEx, _ := regexp.Compile(regexTagSplit)
// }

// HTTPRouter struct
type HTTPRouter struct {
	handler *HTTPHandler
	mux     *http.ServeMux
	address string
}

// NewServer fucntion
func NewServer(address string) HTTPRouter {
	m := http.NewServeMux()
	r := HTTPRouter{
		handler: NewHTTPHandler(),
		address: address,
		mux:     m,
	}
	// Let muxHandler decide what handler serves a specific pattern
	r.mux.Handle("/", r.handler.muxHandler)
	// logger.Info("Handler: %v", r.handler)
	return r
}

// Log method
func (r *HTTPRouter) Log(log bool) {
	r.handler.muxHandler.log = log
	// logger.Info("[HTTPTouter] Log: %v", r.handler.muxHandler.Log)
}

// Handle method
func (r *HTTPRouter) Handle(pattern string, handler http.Handler) *HTTPRouter {
	r.handler.ALL(pattern, handler)
	return r
}

// HandleFunc method
func (r *HTTPRouter) HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) *HTTPRouter {
	h := FuncHandlerWrapper{
		handler: handler,
	}
	r.handler.ALL(pattern, h)
	return r
}

// HandleGET method
func (r *HTTPRouter) HandleGET(pattern string, handler http.Handler) *HTTPRouter {
	r.handler.GET(pattern, handler)
	return r
}

// HandlePOST method
func (r *HTTPRouter) HandlePOST(pattern string, handler http.Handler) *HTTPRouter {
	r.handler.POST(pattern, handler)
	return r
}

// HandlePUT method
func (r *HTTPRouter) HandlePUT(pattern string, handler http.Handler) *HTTPRouter {
	r.handler.PUT(pattern, handler)
	return r
}

// HandleDELETE method
func (r *HTTPRouter) HandleDELETE(pattern string, handler http.Handler) *HTTPRouter {
	r.handler.DELETE(pattern, handler)
	return r
}

// HandleFuncGET method
func (r *HTTPRouter) HandleFuncGET(pattern string, handler func(w http.ResponseWriter, r *http.Request)) *HTTPRouter {
	r.handler.GET(pattern, r.getFuncHandlerWrapper(pattern, handler))
	return r
}

// HandleFuncPOST method
func (r *HTTPRouter) HandleFuncPOST(pattern string, handler func(w http.ResponseWriter, r *http.Request)) *HTTPRouter {
	h := FuncHandlerWrapper{
		handler: handler,
	}
	r.handler.POST(pattern, h)
	return r
}

// HandleFuncPUT method
func (r *HTTPRouter) HandleFuncPUT(pattern string, handler func(w http.ResponseWriter, r *http.Request)) *HTTPRouter {
	h := FuncHandlerWrapper{
		handler: handler,
	}
	r.handler.PUT(pattern, h)
	return r
}

// HandleFuncDELETE method
func (r *HTTPRouter) HandleFuncDELETE(pattern string, handler func(w http.ResponseWriter, r *http.Request)) *HTTPRouter {
	h := FuncHandlerWrapper{
		handler: handler,
	}
	r.handler.DELETE(pattern, h)
	return r
}

///////////////////
// HandlerWrapper getters
///////////////////
func (r *HTTPRouter) getFuncHandlerWrapper(pattern string, handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	// fmt.Println("Requested Path matches:", matches.Match([]byte(pattern)))
	var h http.Handler
	if matches.Match([]byte(pattern)) {
		h = RegExFuncHandlerWrapper{
			handler:    handler,
			muxHandler: r.handler.muxHandler,
		}
	} else {
		h = FuncHandlerWrapper{
			handler: handler,
		}
	}
	return h
}

func (r *HTTPRouter) getHandlerWrapper(pattern string, handler http.Handler) http.Handler {
	// fmt.Println("Requested Path matches:", matches.Match([]byte(pattern)))
	var h http.Handler
	if matches.Match([]byte(pattern)) {
		h = RegExHandlerWrapper{
			handler: handler,
		}
	} else {
		h = HandlerWrapper{
			handler: handler,
		}
	}
	return h
}

// Start method
func (r *HTTPRouter) Start() error {
	err := http.ListenAndServe(r.address, r.mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	return err
}

// GetContextKey method
// func GetContextKey(key string, r *http.Request) interface{} {
// 	cKey := contextKey(key)
// 	ctx := r.Context()
// 	val := ctx.Value(cKey)
// 	return val
// }
