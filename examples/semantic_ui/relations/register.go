// Package relations Index relation
// and more
package relations

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/scigno/webframework/auth"

	"github.com/scigno/webframework/logger"

	"github.com/scigno/webframework/templates"
)

type registerData struct {
	Code  int
	Error int
}

// Register handler.
func Register(w http.ResponseWriter, r *http.Request) {

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

	if r.Method == http.MethodPost {

		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusRequestEntityTooLarge)
			return
		}

		for key, values := range r.PostForm {
			logger.Info("%v : %v", key, values)
		}

		firstName := r.PostForm["first_name"][0]
		lastName := r.PostForm["last_name"][0]
		email := r.PostForm["email"][0]
		pass := r.PostForm["password"][0]
		pass2 := r.PostForm["password_confirm"][0]
		agree := r.PostForm["agree"][0]

		if agree != "on" {
			d.Error = 402
		}

		if emailExists(email) {
			d.Error = 400
		} else {
			if pass[0] == pass2[0] {
				logger.Info("Processing registration...")
				// Generate "hash" to store from user password
				hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
				if err != nil {
					// TODO: Properly handle error
					log.Fatal(err)
				}
				logger.Info("User: %s %s", firstName, lastName)
				logger.Info("Email: %s", email)
				logger.Info("Password: %s", pass)
				logger.Info("Hashed Password: %s", hash)
				http.Redirect(w, r, "/registered", http.StatusFound)
			} else {
				d.Error = 401
			}
		}

	}

	templates.ExecuteTemplate(d, "register", w, r)
}

func emailExists(e string) bool {
	return false
}
