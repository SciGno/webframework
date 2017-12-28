package httprouter

import (
	"net/http"
	"regexp"
	"strings"
)

var buildRegEx, _ = regexp.Compile("{(.*?)}")

// ContextKey type
type ContextKey string

// Router struct
type Router struct {
	simplePath     map[string]map[string]http.Handler
	regexPath      map[string]map[string]http.Handler
	static         http.Handler
	staticPath     string
	statusResponse map[int]http.Handler
}

var statusHandlers = map[int]http.Handler{
	404: Func2Handler(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}),
}

// New returns a new Router
func New() Router {
	// logger.Info("--------------------------------------")
	// logger.Info("[Router] New ")
	return Router{
		simplePath: make(map[string]map[string]http.Handler),
		regexPath:  make(map[string]map[string]http.Handler),
		statusResponse: map[int]http.Handler{
			404: Func2Handler(func(w http.ResponseWriter, r *http.Request) {
				http.NotFound(w, r)
			}),
		},
	}
}

// ServeHTTP function
func (rtr Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// logger.Info("--------------------------------------")
	// logger.Info("[Router] ServeHTTP ")

	// b, _ := json.MarshalIndent(rtr.simplePath, "", "  ")
	// println("Simple Path Map", string(b))
	// c, _ := json.MarshalIndent(rtr.regexPath, "", "  ")
	// println("Regexp Path Map", string(c))

	path := r.URL.Path

	if path[len(path)-1:] == "/" {
		path = path[:len(path)-1]
	}

	if rtr.staticPath != "" && strings.Contains(path, rtr.staticPath) {
		rtr.static.ServeHTTP(w, r)
		return
	}

	// Removing the / at the bigining of path
	if len(path) > 0 {
		path = path[1:]
	}

	if simpleHandler, ok := rtr.simplePath[path]; ok {

		// logger.Info("[Router] simplePath ")
		handler, ok := simpleHandler[r.Method]
		if ok {
			handler.ServeHTTP(w, r)
			return
		}
	} else {

		// logger.Info("[Router] regexPath ")
		for k, v := range rtr.regexPath {
			matchRegexp, _ := regexp.Compile("^" + k + "$")
			lenRegexp := strings.Count(k, "/")
			lenPath := strings.Count(path, "/")

			// Check if path matches pattern and compare the lengths to guarantee
			// the path priority based on length
			if matchRegexp.Match([]byte(path)) && lenRegexp == lenPath {
				handler, ok := v[r.Method]
				if ok {
					handler.ServeHTTP(w, r)
					return
				}
			}
		}
	}
	GetStatusHandler(http.StatusNotFound).ServeHTTP(w, r)
}

////////////////////////////////////
// Handler process section
////////////////////////////////////

// func (rtr *Router) processFuncHandler(p string, method string, h func(w http.ResponseWriter, r *http.Request)) {
// 	logger.Info("--------------------------------------")
// 	logger.Info("[Router] processFuncHandler ")
// 	p = rtr.validatePath(p)

// 	if p[len(p)-1:] == "/" {
// 		p = p[:len(p)-1]
// 	}

// 	if buildRegEx.Match([]byte(p)) { // Path has named parameters
// 		paramsArr, regexPattern, filterRegexp := rtr.getRegExpData(p)
// 		if rm, ok := rtr.regexPath[regexPattern]; ok { // Path in map
// 			rm[method] = rtr.wrapRegexpFunction(paramsArr, filterRegexp, h)
// 		} else { // Path not in map
// 			m := make(map[string]http.Handler)
// 			m[method] = rtr.wrapRegexpFunction(paramsArr, filterRegexp, h)
// 			rtr.regexPath[regexPattern] = m
// 		}
// 	} else { // Path is normal
// 		if rm, ok := rtr.simplePath[p]; ok { // Path in map
// 			rm[method] = rtr.wrapFunction(h)
// 		} else { // Path not in map
// 			m := make(map[string]http.Handler)
// 			m[method] = rtr.wrapFunction(h)
// 			rtr.simplePath[p] = m
// 		}
// 	}
// }

func (rtr *Router) processHandler(p string, method string, h http.Handler) {
	// logger.Info("--------------------------------------")
	// logger.Info("[Router] processHandler ")
	p = rtr.validatePath(p)

	if p[len(p)-1:] == "/" {
		p = p[:len(p)-1]
	}

	if buildRegEx.Match([]byte(p)) { // Path has named parameters
		paramsArr, regexPattern, filterRegexp := rtr.getRegExpData(p)
		if rm, ok := rtr.regexPath[regexPattern]; ok { // Path in map
			rm[method] = rtr.wrapRegexpHandler(paramsArr, filterRegexp, h)
		} else { // Path not in map
			m := make(map[string]http.Handler)
			m[method] = rtr.wrapRegexpHandler(paramsArr, filterRegexp, h)
			rtr.regexPath[regexPattern] = m
		}
	} else { // Path is normal
		if rm, ok := rtr.simplePath[p]; ok { // Path in map
			rm[method] = rtr.wrapHandler(h)
		} else { // Path not in map
			m := make(map[string]http.Handler)
			m[method] = rtr.wrapHandler(h)
			rtr.simplePath[p] = m
		}
	}
}

////////////////////////////////////
// Function handler wrapper section
////////////////////////////////////

// func (rtr Router) wrapRegexpFunction(parameters []string, filter *regexp.Regexp, h func(w http.ResponseWriter, r *http.Request)) http.Handler {

// 	logger.Info("--------------------------------------")
// 	logger.Info("[Router] wrapRegexpFunction ")
// 	return regexFuncWrapper{
// 		handler:      h,
// 		params:       parameters,
// 		paramsFilter: filter,
// 		customWriter: NewCustomWriter(nil),
// 	}
// }

// func (rtr Router) wrapFunction(h func(w http.ResponseWriter, r *http.Request)) http.Handler {

// 	logger.Info("--------------------------------------")
// 	logger.Info("[Router] wrapFunction ")
// 	return funcWrapper{
// 		handler:      h,
// 		customWriter: NewCustomWriter(nil),
// 	}
// }

////////////////////////////////////
// Handler wrapper section
////////////////////////////////////

func (rtr Router) wrapRegexpHandler(parameters []string, filter *regexp.Regexp, h http.Handler) http.Handler {
	// logger.Info("--------------------------------------")
	// logger.Info("[Router] wrapRegexpHandler ")
	return regexHandlerWrapper{
		handler:      h,
		params:       parameters,
		paramsFilter: filter,
		customWriter: NewCustomWriter(nil),
	}
}

// handlerWrapper function
func (rtr Router) wrapHandler(h http.Handler) http.Handler {
	// logger.Info("--------------------------------------")
	// logger.Info("[Router] wrapHandler ")
	return handlerWrapper{
		handler:      h,
		customWriter: NewCustomWriter(nil),
	}
}

////////////////////////////////////
// Other functions
////////////////////////////////////

func (rtr *Router) validatePath(p string) string {
	// logger.Info("--------------------------------------")
	// logger.Info("[Router] validatePath ")
	if p[0] != '/' {
		panic("path must begin with '/' --> " + p)
	}

	if len(p) == 1 {
		return p
	}

	return p[1:]
}

// Returns all named parameters, matching pattern and pointer to regexp object
func (rtr *Router) getRegExpData(p string) ([]string, string, *regexp.Regexp) {
	// logger.Info("--------------------------------------")
	// logger.Info("[Router] getRegExpData ")
	// Build the regex for checking if requested path matches pattern
	regexPattern := buildRegEx.ReplaceAllString(p, ".*")

	// Build pattern to extract request parameters
	filter := buildRegEx.Split(p, -1)

	// logger.Info("Filter: %v", filter)
	arr := []string{}
	for i := 0; i < len(filter); i++ {
		if filter[i] != "" {
			arr = append(arr, "("+filter[i]+")")
		}
	}

	// logger.Info("Array: %v", arr)
	filterRegexp, _ := regexp.Compile(strings.Join(arr, "|"))

	// Get the parameters and add them to a []string
	params := strings.Join(buildRegEx.FindAllString(p, -1), ",")
	cleanRegEx, _ := regexp.Compile("{|}")
	paramsArr := strings.Split(cleanRegEx.ReplaceAllString(params, ""), ",")

	return paramsArr, regexPattern, filterRegexp
}

// Vars returns a mapped value from the request
func Vars(v string, r *http.Request) interface{} {
	// logger.Info("--------------------------------------")
	// logger.Info("[Router ContextKey] Vars ")
	ctx := r.Context()
	return ctx.Value(ContextKey(v))
}

// SetStatusHandler function
func SetStatusHandler(status int, h http.Handler) {
	statusHandlers[status] = h
}

// GetStatusHandler function
func GetStatusHandler(status int) http.Handler {
	return statusHandlers[status]
}
