// Package framework API relation
// and more
package framework

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/scigno/webframework/logger"
)

var buildRegEx, _ = regexp.Compile("\\{(.*?)\\}")

// MuxHandler function
type MuxHandler struct {
	path     map[string]map[string]http.Handler
	regexVal map[string]string
	log      bool
}

// NewMuxHandler creates a handler to be user gor the mux
func NewMuxHandler() *MuxHandler {
	return &MuxHandler{
		path:     make(map[string]map[string]http.Handler),
		regexVal: make(map[string]string),
		log:      false,
	}
}

func (m *MuxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// logger.Info("[MuxHandler] Path: %v", r.URL.Path)
	// logger.Info("[MuxHandler] Map: %v", m.path)
	path := r.URL.Path
	for k, v := range m.path {
		// matchRegEx, err := regexp.Compile(k)
		if matchRegEx, err := regexp.Compile(k); matchRegEx.Match([]byte(path)) && err == nil {
			if _, ok := v[r.Method]; ok {
				(m.path[k][r.Method]).ServeHTTP(w, r)
				// return
				// logger.Info("[MuxHandler] Full match!!! %v", m.path[k][r.Method])
				return
			}
		}
		// logger.Info("[MuxHandler] Path Matches: %v", matchRegEx.Match([]byte(path)))
		// logger.Info("[MuxHandler] RegExPatter: %v", k)
		// logger.Info("[MuxHandler] Value: %v", v)
	}

	// if _, ok := m.path[r.URL.Path]; ok {
	// 	// logger.Info("[MuxHandler] Path1: %v", r.URL.Path)
	// 	m.serveHTTP(w, r, r.URL.Path)
	// 	return
	// }

	// trimed := r.URL.Path[0 : strings.LastIndex(r.URL.Path, "/")+1]
	// if _, ok2 := m.path[trimed]; ok2 {
	// 	// logger.Info("[MuxHandler] Path2: %v", trimed)
	// 	m.serveHTTP(w, r, trimed)
	// 	return
	// }

	if m.log {
		requestErrorLogger(w, r)
	}
	http.NotFound(w, r)

}

func (m *MuxHandler) serveHTTP(w http.ResponseWriter, r *http.Request, path string) {
	if handler := m.path[path][r.Method]; handler != nil {
		// if m.log {
		// 	requestLogger(w, r)
		// }
		handler.ServeHTTP(w, r)
	} else {
		// if m.log {
		// 	requestErrorLogger(w, r)
		// }
		http.NotFound(w, r)
	}
}

// HTTPHandler struct
type HTTPHandler struct {
	muxHandler *MuxHandler
}

// NewHTTPHandler creates a handler to be user gor the mux
func NewHTTPHandler() *HTTPHandler {
	return &HTTPHandler{
		muxHandler: NewMuxHandler(),
	}
}

// GET method
func (h *HTTPHandler) GET(pattern string, handler http.Handler) {
	h.addHandlerToMap(pattern, handler, HTTPGet)
}

// POST method
func (h *HTTPHandler) POST(pattern string, handler http.Handler) {
	h.addHandlerToMap(pattern, handler, HTTPPost)
}

// PUT method
func (h *HTTPHandler) PUT(pattern string, handler http.Handler) {
	h.addHandlerToMap(pattern, handler, HTTPPut)
}

// DELETE method
func (h *HTTPHandler) DELETE(pattern string, handler http.Handler) {
	h.addHandlerToMap(pattern, handler, HTTPDelete)
}

// ALL method
func (h *HTTPHandler) ALL(pattern string, handler http.Handler) {
	h.GET(pattern, handler)
	h.POST(pattern, handler)
	h.PUT(pattern, handler)
	h.DELETE(pattern, handler)
}

func (h *HTTPHandler) addHandlerToMap(pattern string, handler http.Handler, method string) {
	valueSplitPattern := buildRegEx.Split(pattern, -1)
	arr := []string{}
	for i := 0; i < len(valueSplitPattern)-1; i++ {
		//fmt.Println(regexPattern2[i])
		arr = append(arr, "("+valueSplitPattern[i]+")")
	}
	splitValuesRegEx := strings.Join(arr, "|")
	pattern = "^" + buildRegEx.ReplaceAllString(pattern, ".*") + "$"
	h.muxHandler.regexVal[pattern] = splitValuesRegEx

	logger.Info("RegEx Pattern: %v", pattern)
	logger.Info("SplitValuesRegEx: %v", splitValuesRegEx)
	_, ok := h.muxHandler.path[pattern]
	if !ok {
		mm := make(map[string]http.Handler)
		mm[method] = handler
		h.muxHandler.path[pattern] = mm
	} else {
		h.muxHandler.path[pattern][method] = handler
	}
}
