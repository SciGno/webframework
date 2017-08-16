package httprouter

import (
	"net/http"
	"regexp"
	"strings"
)

var buildRegEx, _ = regexp.Compile("{(.*?)}")

type contextKey string

// Router struct
type Router struct {
	simplePath map[string]map[string]http.Handler
	regexPath  map[string]map[string]http.Handler
	static     http.Handler
	staticPath string
}

// New returns a new Router
func New() Router {
	return Router{
		simplePath: make(map[string]map[string]http.Handler),
		regexPath:  make(map[string]map[string]http.Handler),
	}
}

// ServeHTTP function
func (rtr Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// b, _ := json.MarshalIndent(rtr.simplePath, "", "  ")
	// println("Simple Path Map", string(b))
	// c, _ := json.MarshalIndent(rtr.regexPath, "", "  ")
	// println("Regexp Path Map", string(c))

	path := r.URL.Path

	// logger.Info("Path: %v", path)

	if path[len(path)-1:] == "/" {
		path = path[:len(path)-1]
	}

	// logger.Info("New Path: %v", path)
	// logger.Info("staticPath: %v", rtr.staticPath)

	if rtr.staticPath != "" && strings.Contains(path, rtr.staticPath) {
		// logger.Info("Path: %v, Static: %v, Contains: %v", path, rtr.staticPath, strings.Contains(path, rtr.staticPath))
		rtr.static.ServeHTTP(w, r)
		return
	}

	// Removing the / at the bigining of path
	if len(path) > 0 {
		path = path[1:]
	}

	if simpleHandler, ok := rtr.simplePath[path]; ok {
		// logger.Info("[ServeHTTP] Simple path...")
		handler, ok := simpleHandler[r.Method]
		if ok {
			handler.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	} else {
		// logger.Info("[ServeHTTP] Regexp path...")
		for k, v := range rtr.regexPath {
			// logger.Info("[ServeHTTP] all: %v, %v", k, path)
			// logger.Info("[ServeHTTP] Regexp: %v", regexp)
			matchRegexp, _ := regexp.Compile("^" + k + "$")
			lenRegexp := strings.Count(k, "/")
			lenPath := strings.Count(path, "/")

			// logger.Info("[ServeHTTP] Length: %v, %v", lenRegexp, lenPath)
			// Check if path matches pattern and compare the lengths to guarantee
			// the path priority based on length
			if matchRegexp.Match([]byte(path)) && lenRegexp == lenPath {
				handler, ok := v[r.Method]
				if ok {
					// logger.Info("[ServeHTTP] Regexp path: %v", path)
					handler.ServeHTTP(w, r)
				} else {
					http.NotFound(w, r)
				}
				return
			}
		}
		http.NotFound(w, r)
	}
}

////////////////////////////////////
// Handler process section
////////////////////////////////////

func (rtr *Router) processFuncHandler(p string, method string, h func(w http.ResponseWriter, r *http.Request)) {
	p = rtr.validatePath(p)

	if p[len(p)-1:] == "/" {
		p = p[:len(p)-1]
	}

	if buildRegEx.Match([]byte(p)) { // Path has named parameters
		paramsArr, regexPattern, filterRegexp := rtr.getRegExpData(p)
		if rm, ok := rtr.regexPath[regexPattern]; ok { // Path in map
			rm[method] = rtr.wrapRegexpFunction(paramsArr, filterRegexp, h)
		} else { // Path not in map
			m := make(map[string]http.Handler)
			m[method] = rtr.wrapRegexpFunction(paramsArr, filterRegexp, h)
			rtr.regexPath[regexPattern] = m
		}
	} else { // Path is normal
		if rm, ok := rtr.simplePath[p]; ok { // Path in map
			rm[method] = rtr.wrapFunction(h)
		} else { // Path not in map
			m := make(map[string]http.Handler)
			m[method] = rtr.wrapFunction(h)
			rtr.simplePath[p] = m
		}
	}
}

func (rtr *Router) processHandler(p string, method string, h http.Handler) {
	p = rtr.validatePath(p)

	if p[len(p)-1:] == "/" {
		p = p[:len(p)-1]
	}

	if buildRegEx.Match([]byte(p)) { // Path has named parameters
		// logger.Info("[processHandler] Regexp path: %v", p)
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

func (rtr Router) wrapRegexpFunction(parameters []string, filter *regexp.Regexp, h func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return regexFuncWrapper{
		handler:      h,
		params:       parameters,
		paramsFilter: filter,
		customWriter: NewCustomWriter(nil),
	}
}

func (rtr Router) wrapFunction(h func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return funcWrapper{
		handler:      h,
		customWriter: NewCustomWriter(nil),
	}
}

////////////////////////////////////
// Handler wrapper section
////////////////////////////////////

func (rtr Router) wrapRegexpHandler(parameters []string, filter *regexp.Regexp, h http.Handler) http.Handler {
	return regexHandlerWrapper{
		handler:      h,
		params:       parameters,
		paramsFilter: filter,
		customWriter: NewCustomWriter(nil),
	}
}

// handlerWrapper function
func (rtr Router) wrapHandler(h http.Handler) http.Handler {
	return handlerWrapper{
		handler:      h,
		customWriter: NewCustomWriter(nil),
	}
}

////////////////////////////////////
// Other functions
////////////////////////////////////

func (rtr *Router) validatePath(p string) string {
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

	// logger.Info("Path: %v", p)
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

	// logger.Info("%v %v %v", paramsArr, regexPattern, filterRegexp)

	return paramsArr, regexPattern, filterRegexp
}

// Vars returns a mapped value from the request
func Vars(v string, r *http.Request) interface{} {
	ctx := r.Context()
	return ctx.Value(contextKey(v))
}
