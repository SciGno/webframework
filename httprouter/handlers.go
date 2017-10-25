package httprouter

import "net/http"

////////////////////////////
// Static Handler section
////////////////////////////

// HandleStatic function
func (rtr *Router) HandleStatic(p string, h http.Handler) *Router {
	rtr.static = h
	rtr.staticPath = p
	return rtr
}

////////////////////////////
// Handler section
////////////////////////////

// HandleGET function
func (rtr *Router) HandleGET(p string, h http.Handler) *Router {
	rtr.processHandler(p, http.MethodGet, h)
	return rtr
}

// HandlePOST function
func (rtr *Router) HandlePOST(p string, h http.Handler) *Router {
	rtr.processHandler(p, http.MethodPost, h)
	return rtr
}

// HandlePUT function
func (rtr *Router) HandlePUT(p string, h http.Handler) *Router {
	rtr.processHandler(p, http.MethodPut, h)
	return rtr
}

// HandleDELETE function
func (rtr *Router) HandleDELETE(p string, h http.Handler) *Router {
	rtr.processHandler(p, http.MethodDelete, h)
	return rtr
}

// Handle function
func (rtr *Router) Handle(p string, h http.Handler) *Router {
	rtr.processHandler(p, http.MethodGet, h)
	rtr.processHandler(p, http.MethodPost, h)
	rtr.processHandler(p, http.MethodPut, h)
	rtr.processHandler(p, http.MethodDelete, h)
	return rtr
}

////////////////////////////
// Function handler section
////////////////////////////

// HandleFuncGET function
func (rtr *Router) HandleFuncGET(p string, h func(w http.ResponseWriter, r *http.Request)) *Router {
	rtr.processFuncHandler(p, http.MethodGet, h)
	return rtr
}

// HandleFuncPOST function
func (rtr *Router) HandleFuncPOST(p string, h func(w http.ResponseWriter, r *http.Request)) *Router {
	rtr.processFuncHandler(p, http.MethodPost, h)
	return rtr
}

// HandleFuncPUT function
func (rtr *Router) HandleFuncPUT(p string, h func(w http.ResponseWriter, r *http.Request)) *Router {
	rtr.processFuncHandler(p, http.MethodPut, h)
	return rtr
}

// HandleFuncDELETE function
func (rtr *Router) HandleFuncDELETE(p string, h func(w http.ResponseWriter, r *http.Request)) *Router {
	rtr.processFuncHandler(p, http.MethodDelete, h)
	return rtr
}

// HandleFunc function
func (rtr *Router) HandleFunc(p string, h func(w http.ResponseWriter, r *http.Request)) *Router {
	rtr.processFuncHandler(p, http.MethodGet, h)
	rtr.processFuncHandler(p, http.MethodPost, h)
	rtr.processFuncHandler(p, http.MethodPut, h)
	rtr.processFuncHandler(p, http.MethodDelete, h)
	return rtr
}
