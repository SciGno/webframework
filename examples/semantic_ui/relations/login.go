// Package relations Index relation
// and more
package relations

import (
	"fmt"
	"net/http"
	"time"

	"github.com/scigno/webframework/auth"

	"github.com/scigno/webframework/logger"

	"github.com/scigno/webframework/templates"
)

type loginData struct {
	Code  int
	Error int
}

// Login handler.
func Login(w http.ResponseWriter, r *http.Request) {

	// for k, v := range r.Header {
	// 	logger.Info("%v : %v", k, v)
	// }
	// logger.Info("------- LOGIN")
	logger.Info("CookieFound: %+v", auth.JWTContextValue(auth.JWTCookieFound, r))
	logger.Info("TokenFound: %+v", auth.JWTContextValue(auth.JWTTokenFound, r))
	logger.Info("TokenValid: %+v", auth.JWTContextValue(auth.JWTTokenValid, r))
	logger.Info("TokenClaims: %+v", auth.JWTContextValue(auth.JWTClaim, r))

	d := loginData{}
	d.Code = 200

	if auth.JWTContextValue(auth.JWTTokenValid, r).(bool) {
		d.Code = 201
	}

	if r.Method == http.MethodPost {

		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusRequestEntityTooLarge)
			return
		}

		for key, values := range r.PostForm {
			logger.Info("%v : %v", key, values)
		}
		email, eok := r.PostForm["email"]
		pass, pok := r.PostForm["password"]

		if eok && pok && email[0] == "leandro.lopez@gmail.com" && pass[0] == "password" {
			m := map[string]interface{}{
				"admin":  true,
				"groups": []string{"developer", "operations"},
			}
			t, _ := auth.NewJWSToken("", time.Duration(10)*time.Second, "", "leandro", time.Now(), "testing", m)
			token := fmt.Sprintf("%s", t)
			logger.Info("Token: %s", token)

			http.SetCookie(w, &http.Cookie{
				Name:     "access_token",
				Value:    token,
				Path:     "/",
				Expires:  time.Now().Add(10 * time.Second),
				HttpOnly: true,
				// Secure:   true,
			})

			// w.Header().Set("Authorization", "Bearer "+token)
			// w.Header().Set("WWW-Authenticate", "Basic realm=\"user\"")
			http.Redirect(w, r, "/settings", http.StatusFound)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		d.Error = 401
	}
	templates.ExecuteTemplate(d, "login", w, r)
}
