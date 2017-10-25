// Package relations Index relation
// and more
package relations

import (
	"net/http"

	"github.com/scigno/webframework/auth"

	"github.com/scigno/webframework/logger"

	"github.com/scigno/webframework/templates"
)

type registeredData struct {
	Code  int
	Error int
}

// Registered handler.
func Registered(w http.ResponseWriter, r *http.Request) {

	// for k, v := range r.Header {
	// 	logger.Info("%v : %v", k, v)
	// }
	// logger.Info("------- LOGIN")
	logger.Info("CookieFound: %+v", auth.JWTContextValue(auth.JWTCookieFound, r))
	logger.Info("TokenFound: %+v", auth.JWTContextValue(auth.JWTTokenFound, r))
	logger.Info("TokenValid: %+v", auth.JWTContextValue(auth.JWTTokenValid, r))
	logger.Info("TokenClaims: %+v", auth.JWTContextValue(auth.JWTClaim, r))

	d := registerData{}
	d.Code = 200

	if auth.JWTContextValue(auth.JWTTokenValid, r).(bool) {
		d.Code = 201
	}

	templates.ExecuteTemplate(d, "registered", w, r)
}
