package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/scigno/webframework/logger"
)

// ContextKey type
type ContextKey string

const (
	// JWTCookieFound is used for Request context
	JWTCookieFound = "JWTCookieFound"
	// JWTCookieReadError is used for Request context
	JWTCookieReadError = "JWTCookieReadFound"
	// JWTParseTokenError is used for Request context
	JWTParseTokenError = "JWTParseTokenError"
	// JWTTokenFound is used for Request context
	JWTTokenFound = "JWTTokenFound"
	// JWTTokenValid is used for Request context
	JWTTokenValid = "JWTTokenValid"
	// JWTClaim is used for Request context
	JWTClaim = "JWTClaim"
)

//-----------------------------------------

// JWTCookieSecuredFunc struct
type JWTCookieSecuredFunc struct {
	handler func(w http.ResponseWriter, r *http.Request)
	cookie  string
}

// ServeHTTP function
func (s JWTCookieSecuredFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// logger.Info("BEGIN: SecureFunc...")
	// c, _ := r.Cookie(s.cookie)
	// logger.Info("Cookie: %+v", c)

	// for k, v := range r.Header {
	// 	logger.Info("%v : %v", k, v)
	// }
	// logger.Info("------- PROTECTED")

	ctx := r.Context()
	ctx = context.WithValue(ctx, ContextKey(JWTTokenFound), false)
	ctx = context.WithValue(ctx, ContextKey(JWTTokenValid), false)

	tokenCookie, err := r.Cookie(s.cookie)
	if err == http.ErrNoCookie {
		logger.Error("Erro: No cookie found")
		ctx = context.WithValue(ctx, ContextKey(JWTCookieFound), false)
	} else if err != nil {
		logger.Error("Unable to read cookie")
		ctx = context.WithValue(ctx, ContextKey(JWTCookieReadError), true)
	} else {
		ctx = context.WithValue(ctx, ContextKey(JWTCookieFound), true)
		ctx = context.WithValue(ctx, ContextKey(JWTTokenFound), true)

		j, err := jws.ParseJWT([]byte(tokenCookie.Value))
		if err != nil {
			logger.Error("Parse Error: %+v", err)
			ctx = context.WithValue(ctx, ContextKey(JWTParseTokenError), true)
		}

		// Validate token
		rsaPublic, _ := crypto.ParseRSAPublicKeyFromPEM([]byte(public[0]))
		if err2 := j.Validate(rsaPublic, crypto.SigningMethodRS256); err2 != nil {
			ctx = context.WithValue(ctx, ContextKey(JWTTokenValid), false)
			logger.Error("Token Error: %+v", err2)
		} else {
			ctx = context.WithValue(ctx, ContextKey(JWTTokenValid), true)
			ctx = context.WithValue(ctx, ContextKey(JWTClaim), j.Claims())
		}

	}
	s.handler(w, r.WithContext(ctx))
	return
}

// JWTProtectedFunc function
func JWTProtectedFunc(f func(w http.ResponseWriter, r *http.Request), name string) JWTCookieSecuredFunc {
	return JWTCookieSecuredFunc{
		handler: f,
		cookie:  name,
	}
}

//-----------------------------------------

// // JWTSecureHandler struct
// type JWTSecureHandler struct {
// 	handler http.Handler
// 	claims  map[http.ResponseWriter]map[string]string
// }

// // ServeHTTP function
// func (sh JWTSecureHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	logger.Info("BEGIN: SecureHandler...")
// 	m := sh.claims[w]
// 	m["key"] = "value"
// 	sh.handler.ServeHTTP(w, r)
// 	logger.Info("Headers: %v", r.Header)
// 	logger.Info("END: SecureHandler...")
// }

// // Claims function
// func (sh JWTSecureHandler) Claims(w http.ResponseWriter) map[string]string {
// 	return sh.claims[w]
// }

// // JWTProtectedHandler function
// func JWTProtectedHandler(h http.Handler) JWTSecureHandler {
// 	return JWTSecureHandler{
// 		handler: h,
// 		claims:  map[http.ResponseWriter]map[string]string{},
// 	}
// }

//-----------------------------------------

// // JWTLoginFuncHandler struct
// type JWTLoginFuncHandler struct {
// 	handler func(w http.ResponseWriter, r *http.Request)
// 	claims  map[http.ResponseWriter]map[string]string
// }

// // ServeHTTP function
// func (l JWTLoginFuncHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	logger.Info("BEGIN: SecureFunc...")

// 	// Set custom claims
// 	claims := jwtCustomClaims{
// 		map[string]interface{}{"name": "Leandro", "admin": "true"},
// 		jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Minute * 1).UTC().Unix(),
// 			IssuedAt:  time.Now().UTC().Unix(),
// 			Issuer:    "Leandro Lopez",
// 			Subject:   "MarketBin",
// 		},
// 	}

// 	// Create token with claims
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	// Generate encoded token and send it as response.
// 	t, err := token.SignedString([]byte("secret"))
// 	if err != nil {
// 		return
// 	}

// 	w.Header().Set("Authorization", "Bearer "+t)

// 	l.handler(w, r)
// 	logger.Info("Headers: %v", w.Header())
// 	logger.Info("END: SecureFunc...")
// }

// // Claims function
// func (l JWTLoginFuncHandler) Claims(w http.ResponseWriter) map[string]string {
// 	return l.claims[w]
// }

// // JWTLoginFunc function
// func JWTLoginFunc(f func(w http.ResponseWriter, r *http.Request), v func(username string, password string) (bool, map[string]string)) JWTLoginFuncHandler {
// 	return JWTLoginFuncHandler{
// 		handler: f,
// 		claims:  map[http.ResponseWriter]map[string]string{},
// 	}
// }

//-----------------------------------------

// func keyLookupFunc(token *jwt.Token) (interface{}, error) {
// 	// Don't forget to validate the alg is what you expect:
// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 	}
// 	// Unpack key from PEM encoded PKCS8
// 	return []byte(VerifyKeys[0]), nil
// }

// // JWTFormLoginHandler struct
// type JWTFormLoginHandler struct {
// 	authenticator Authenticator
// 	validator     Validator
// 	handler       http.Handler
// 	redirect      string
// 	timeout       time.Duration
// }

// // ServeHTTP function
// func (l JWTFormLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == http.MethodGet {
// 		logger.Info("Method GET")
// 		token := ParseJWTCookieToken(r)
// 		if token != nil {
// 			if token.Valid {
// 				// Possibly refresh token
// 				// if l.validator != nil {
// 				// 	if !l.validator.Valid() {
// 				// 		w.WriteHeader(http.StatusUnauthorized)
// 				// 		return
// 				// 	}
// 				// }
// 				logger.Info("Token headers: %+v", token)
// 				http.Redirect(w, r, l.redirect, http.StatusFound)
// 				return
// 			}
// 		}

// 		ctx := r.Context()
// 		ctx = context.WithValue(ctx, ContextKey("test"), map[string]interface{}{"key": "value"})
// 		ctx = context.WithValue(ctx, ContextKey("test2"), map[string]interface{}{"key2": "value2"})

// 		logger.Info("[GET] No JWT Token found...")
// 		l.handler.ServeHTTP(w, r.WithContext(ctx))
// 		return
// 	}

// 	if r.Method == http.MethodPost {
// 		ctx := r.Context()

// 		logger.Info("[POST] No JWT Token")

// 		if err := r.ParseForm(); err != nil {
// 			w.WriteHeader(http.StatusRequestEntityTooLarge)
// 			return
// 		}

// 		m := map[string]interface{}{}
// 		for key, values := range r.PostForm {
// 			m[key] = values[0]
// 		}

// 		logger.Info("MAP: %v, %v", m["email"], m["password"])
// 		ok, umap := l.authenticator.Authenticate(m)
// 		if ok {
// 			logger.Info("MAP: %v", umap)
// 			if l.handler != nil {
// 				// TODO: Add JWT token
// 				token := CreateCustomJWTToken(umap, l.timeout)
// 				// cookie := http.Cookie{
// 				// 	Name:  "token",
// 				// 	Value: token,
// 				// }
// 				// http.SetCookie(w, &cookie)

// 				http.SetCookie(w, &http.Cookie{
// 					Name:       "token",
// 					Value:      token,
// 					Path:       "/",
// 					RawExpires: "0",
// 				})

// 				// logger.Info("Redirecting...")
// 				// r.Method = http.MethodGet
// 				http.Redirect(w, r, l.redirect, http.StatusFound)
// 				// l.handler.ServeHTTP(w, r)
// 				return
// 			}
// 		} else {
// 			// logger.Info("Invalid credentials...")
// 			w.WriteHeader(http.StatusUnauthorized)
// 			ctx = context.WithValue(ctx, templates.TemplateContext("http_status"), http.StatusUnauthorized)
// 			l.handler.ServeHTTP(w, r.WithContext(ctx))
// 			return
// 		}
// 	}

// 	// logger.Info("Header: %v", r.Header)

// }

// // JWTFormLogin function
// func JWTFormLogin(a Authenticator, v Validator, next http.Handler, r string, t time.Duration) JWTFormLoginHandler {
// 	return JWTFormLoginHandler{
// 		authenticator: a,
// 		validator:     v,
// 		handler:       next,
// 		redirect:      r,
// 		timeout:       t,
// 	}
// }

// ------------------------------------------

// // JWTAuthenticator struct
// type JWTAuthenticator struct {
// 	handler http.Handler
// 	timeout time.Duration
// }

// func (a JWTAuthenticator) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// }

// // JWTAuthenticate function
// func JWTAuthenticate(next http.Handler, t time.Duration) JWTAuthenticator {
// 	return JWTAuthenticator{
// 		handler: next,
// 		timeout: t,
// 	}
// }

//-----------------------------------------

// // JWTJSONLoginHandler struct
// type JWTJSONLoginHandler struct {
// 	verifier func(map[string]string) (bool, map[string]string)
// }

// // ServeHTTP function
// func (l JWTJSONLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// 	// logger.Info("APILogin: %+v", r)

// 	// http.Redirect(w, r, r.Host+"/", http.StatusPermanentRedirect)
// 	// logger.Info("BEGIN: SecureHandler...")
// 	// logger.Info("Headers: %v", r.Header)
// 	// logger.Info("END: SecureHandler...")
// }

// // JWTJSONLogin function
// func JWTJSONLogin(v func(map[string]string) (bool, map[string]string)) JWTJSONLoginHandler {
// 	return JWTJSONLoginHandler{
// 		verifier: v,
// 	}
// }

// JWTContextValue func
func JWTContextValue(k string, r *http.Request) interface{} {
	ctx := r.Context()
	// logger.Info("[JWTContextValue] Cookie: %v", ctx.Value(ContextKey(k)))
	return ctx.Value(ContextKey(k))
}

// NewJWSToken function
func NewJWSToken(jwtid string, duration time.Duration, audience string, issuer string, notbefore time.Time, subject string, custom map[string]interface{}) ([]byte, error) {
	// bytes, _ := ioutil.ReadFile("./app.rsa")
	claims := jws.Claims{}
	if jwtid != "" {
		claims.SetJWTID(jwtid)
	}
	claims.SetExpiration(time.Now().Add(duration))
	claims.SetIssuedAt(time.Now())
	if audience != "" {
		claims.SetAudience(audience)
	}
	if issuer != "" {
		claims.SetIssuer(issuer)
	}
	if notbefore.After(time.Now()) {
		claims.SetNotBefore(notbefore)
	}
	if subject != "" {
		claims.SetSubject(subject)
	}
	if custom != nil {
		for k, v := range custom {
			claims.Set(k, v)
		}
	}

	rsaPrivate, err := crypto.ParseRSAPrivateKeyFromPEM([]byte(private[0]))
	if err != nil {
		logger.Info("Error: %v", err)
		return nil, err
	}

	jwt := jws.NewJWT(claims, crypto.SigningMethodRS256)
	b, _ := jwt.Serialize(rsaPrivate)

	// // output claims
	// c, _ := claims.MarshalJSON()
	// var out bytes.Buffer
	// err2 := json.Indent(&out, c, "", "    ")
	// if err2 == nil {
	// 	logger.Info("Claims: %s", out.Bytes())
	// }
	// logger.Info("Token: %s", b)

	return b, nil
}
