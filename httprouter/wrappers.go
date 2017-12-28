package httprouter

import (
	"context"
	"net/http"
	"os"
	"regexp"
)

var hostname, _ = os.Hostname()

// type regexFuncWrapper struct {
// 	handler      func(w http.ResponseWriter, r *http.Request)
// 	params       []string
// 	paramsFilter *regexp.Regexp
// 	customWriter CustomResponseWriter
// }

// func (f regexFuncWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	logger.Info("--------------------------------------")
// 	logger.Info("[regexFuncWrapper] ServeHTTP ")
// 	// logger.Info("| serving | %s | %s | %s", r.Method, r.RemoteAddr, r.RequestURI)
// 	ctx := r.Context()
// 	path := r.URL.Path
// 	values := f.paramsFilter.Split(path, -1)
// 	// logger.Info("Params Filter: %s", f.paramsFilter.String())

// 	// Remove empty elements from array
// 	arr := []string{}
// 	for _, v := range values {
// 		if v != "" && v != "/" {
// 			arr = append(arr, v)
// 		}
// 	}

// 	// logger.Info("Array: %v", arr)
// 	// logger.Info("Len: %v", len(arr))
// 	// Add key=values to context
// 	for i, v := range arr {
// 		logger.Info("[regexFuncWrapper] Index: %v", i)
// 		ctx = context.WithValue(ctx, ContextKey(f.params[i]), v)
// 	}

// 	f.customWriter.ResponseWriter = w
// 	f.handler(&f.customWriter, r.WithContext(ctx))
// 	// statusCode := f.customWriter.statusCode
// 	// logger.Info("| completed | %s | %s | %s | Status: %d", r.Method, r.RemoteAddr, r.RequestURI, statusCode)
// }

////////////////////////////////////////////////////////////////////////////////////////////

type regexHandlerWrapper struct {
	handler      http.Handler
	params       []string
	paramsFilter *regexp.Regexp
	customWriter CustomResponseWriter
}

func (h regexHandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// logger.Info("| serving | %s | %s | %s", r.Method, r.RemoteAddr, r.RequestURI)
	ctx := r.Context()
	path := r.URL.Path
	params := h.paramsFilter.Split(path, -1)

	// Remove empty elements from array
	arr := []string{}
	for _, v := range params {
		if v != "" && v != "/" {
			arr = append(arr, v)
		}
	}
	for i, v := range arr {
		ctx = context.WithValue(ctx, ContextKey(h.params[i]), v)
	}
	h.customWriter.ResponseWriter = w
	h.handler.ServeHTTP(&h.customWriter, r.WithContext(ctx))
	// statusCode := h.customWriter.statusCode
	// logger.Info("| completed | %s | %s | %s | Status: %d", r.Method, r.RemoteAddr, r.RequestURI, statusCode)
}

////////////////////////////////////////////////////////////////////////////////////////////

type handlerWrapper struct {
	handler      http.Handler
	customWriter CustomResponseWriter
}

func (h handlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// logger.Info("--------------------------------------")
	// logger.Info("[handlerWrapper] ServeHTTP ")
	// logger.Info("| serving | %s | %s | %s", r.Method, r.RemoteAddr, r.RequestURI)
	h.customWriter.ResponseWriter = w
	h.handler.ServeHTTP(&h.customWriter, r)
	if sh := GetStatusHandler(h.customWriter.statusCode); sh != nil {
		sh.ServeHTTP(w, r)
	}
	// statusCode := h.customWriter.statusCode
	// logger.Info("| completed | %s | %s | %s | Status: %d", r.Method, r.RemoteAddr, r.RequestURI, statusCode)
}

////////////////////////////////////////////////////////////////////////////////////////////

// type funcWrapper struct {
// 	handler      func(w http.ResponseWriter, r *http.Request)
// 	customWriter CustomResponseWriter
// }

// func (f funcWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	logger.Info("--------------------------------------")
// 	logger.Info("[funcWrapper] ServeHTTP ")
// 	// logger.Info("| serving | %s | %s | %s", r.Method, r.RemoteAddr, r.RequestURI)
// 	f.customWriter.ResponseWriter = w
// 	f.handler(&f.customWriter, r)
// 	if sh := GetStatusHandler(f.customWriter.statusCode); sh != nil {
// 		sh.ServeHTTP(w, r)
// 	}
// 	// statusCode := f.customWriter.statusCode
// 	// logger.Info("| completed | %s | %s | %s | Status: %d", r.Method, r.RemoteAddr, r.RequestURI, statusCode)
// }

////////////////////////////////////////////////////////////////////////////////////////////

// CustomResponseWriter struc
type CustomResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// NewCustomWriter function
func NewCustomWriter(w http.ResponseWriter) CustomResponseWriter {
	// logger.Info("--------------------------------------")
	// logger.Info("[NewCustomWriter] () ")
	return CustomResponseWriter{w, http.StatusOK}
}

// Header method
func (c *CustomResponseWriter) Header() http.Header {
	// logger.Info("--------------------------------------")
	// logger.Info("[NewCustomWriter] Header ")
	return c.ResponseWriter.Header()
}

// WriteHeader method
func (c *CustomResponseWriter) WriteHeader(status int) {
	// logger.Info("--------------------------------------")
	// logger.Info("[NewCustomWriter] WriteHeader ")
	c.statusCode = status
	c.ResponseWriter.WriteHeader(c.statusCode)
}

////////////////////////////////////////////////////////////////////////////////////////////

// SimpleFuncWrapper struc for function to handler wrapping
type SimpleFuncWrapper struct {
	Handler func(w http.ResponseWriter, r *http.Request)
}

func (f SimpleFuncWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// logger.Info("--------------------------------------")
	// logger.Info("[SimpleFuncWrapper] ServeHTTP ")
	f.Handler(w, r)
}

// Func2Handler wrapper
func Func2Handler(f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	// logger.Info("--------------------------------------")
	// logger.Info("[Func2Handler] () ")
	return SimpleFuncWrapper{
		f,
	}
}
