package httprouter

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/scigno/webframework/logger"
)

var hostname, _ = os.Hostname()

// Entity struct
type Entity struct {
	entity string
}

// Event struct
type Event struct {
	entity string
}

// Identity struct
type Identity struct {
	entity string
}

// From struct
type From struct {
	entity string
}

// To struct
type To struct {
	entity string
}

// Relation struct
type Relation struct {
	event Event
}

// Subject struct
type Subject struct {
	identity Identity
}

// RED struct
type RED struct {
	from     From
	to       To
	relation Relation
	subject  Subject
}

type regexFuncWrapper struct {
	handler      func(w http.ResponseWriter, r *http.Request)
	params       []string
	paramsFilter *regexp.Regexp
	customWriter CustomWriter
}

func createREDMsg(from string, to string, event string, subject string, smap interface{}) interface{} {
	return map[string]interface{}{
		"meta": map[string]interface{}{
			"version": "1.0",
			"time":    time.Now(),
		},
		"from": map[string]interface{}{
			"entity": from,
		},
		"to": map[string]interface{}{
			"entity": to,
		},
		"relation": map[string]interface{}{
			"event": event,
		},
		"subject": map[string]interface{}{
			"identity":   "httpd",
			"properties": smap,
		},
	}
}

func (f regexFuncWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// reqLog, _ := json.MarshalIndent(req, "", "  ")
	reqLog, _ := json.Marshal(createREDMsg(r.Host, hostname, "request", "httpd", map[string]interface{}{"method": r.Method, "url": r.RequestURI}))
	logger.Info("%s", string(reqLog))

	ctx := r.Context()
	path := r.URL.Path
	values := f.paramsFilter.Split(path, -1)
	// logger.Info("Params Filter: %s", f.paramsFilter.String())

	// Remove empty elements from array
	arr := []string{}
	for _, v := range values {
		if v != "" {
			arr = append(arr, v)
		}
	}
	// Add key=values to context
	for i, v := range arr {
		ctx = context.WithValue(ctx, contextKey(f.params[i]), v)
	}

	f.customWriter.ResponseWriter = w
	f.handler(&f.customWriter, r.WithContext(ctx))
	statusCode := f.customWriter.statusCode

	// respLog, _ := json.MarshalIndent(resp, "", "  ")
	respLog, _ := json.Marshal(createREDMsg(hostname, r.Host, "response", "httpd", map[string]interface{}{"method": r.Method, "url": r.RequestURI, "status": strconv.Itoa(statusCode)}))
	logger.Info("%s", string(respLog))

	// logger.Info("| completed | %s | %s | %s | Status: %d", r.Method, r.RemoteAddr, r.RequestURI, statusCode)
}

////////////////////////////////////////////////////////////////////////////////////////////

type regexHandlerWrapper struct {
	handler      http.Handler
	params       []string
	paramsFilter *regexp.Regexp
	customWriter CustomWriter
}

func (h regexHandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Info("| serving | %s | %s | %s", r.Method, r.RemoteAddr, r.RequestURI)
	path := r.URL.Path
	params := h.paramsFilter.Split(path, -1)
	for i, v := range params {
		if v == "" {
			params = append(params[:i], params[i+1:]...)
		}
	}

	h.customWriter.ResponseWriter = w
	h.handler.ServeHTTP(&h.customWriter, r)
	statusCode := h.customWriter.statusCode

	// h.handler.ServeHTTP(w, r)
	logger.Info("| completed | %s | %s | %s | Status: %d", r.Method, r.RemoteAddr, r.RequestURI, statusCode)
}

////////////////////////////////////////////////////////////////////////////////////////////

type handlerWrapper struct {
	handler      http.Handler
	customWriter CustomWriter
}

func (h handlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Info("| serving | %s | %s | %s", r.Method, r.RemoteAddr, r.RequestURI)
	h.customWriter.ResponseWriter = w
	h.handler.ServeHTTP(&h.customWriter, r)
	statusCode := h.customWriter.statusCode
	// h.handler.ServeHTTP(w, r)
	logger.Info("| completed | %s | %s | %s | Status: %d", r.Method, r.RemoteAddr, r.RequestURI, statusCode)
}

////////////////////////////////////////////////////////////////////////////////////////////

type funcWrapper struct {
	handler      func(w http.ResponseWriter, r *http.Request)
	customWriter CustomWriter
}

func (f funcWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// logger.Error("| serving | %s | %s | %s ", r.Method, r.RemoteAddr, r.RequestURI)
	// f.customWriter.ResponseWriter = w
	// f.handler(&f.customWriter, r)
	// statusCode := f.customWriter.statusCode
	f.handler(w, r)
	logger.Info("| COMPLETED | %s | %s | %s", r.Method, r.RemoteAddr, r.RequestURI)
	// logger.Info("| funcWrapper [%v]|", statusCode)
}

////////////////////////////////////////////////////////////////////////////////////////////

// CustomWriter struc
type CustomWriter struct {
	http.ResponseWriter
	statusCode int
}

// NewCustomWriter function
func NewCustomWriter(w http.ResponseWriter) CustomWriter {
	return CustomWriter{w, http.StatusOK}
}

// WriteHeader method
func (c *CustomWriter) WriteHeader(status int) {
	c.statusCode = status
	c.ResponseWriter.WriteHeader(c.statusCode)
}
