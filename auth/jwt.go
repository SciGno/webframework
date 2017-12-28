package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"stagezero.com/leandro/marketbin/api"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/scigno/webframework/auth/policy"
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
	// JWTPolicyValid is used for Request context
	JWTPolicyValid = "JWTPolicyValid"
	// JWTPolicyKey is the JWT key for the policy
	JWTPolicyKey = "policy"
)

//-----------------------------------------

// JWTSecuredHandler struct
type JWTSecuredHandler struct {
	handler             http.Handler
	processPolicy       bool
	policyValidator     policy.Validator
	rejectInvalidPolicy bool
}

// SetPolicyValidator function
func (s JWTSecuredHandler) SetPolicyValidator(func(p policy.Policy, r *http.Request)) JWTSecuredHandler {
	return s
}

// ServeHTTP function
func (s JWTSecuredHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// logger.Info("[JWTSecuredHandler] ServeHTTP")
	// logger.Info("[JWTSecuredHandler] Headers: %v", r.Header)

	ctx := r.Context()
	// userid := httprouter.Vars("user_id", r).(string)
	headerToken, err := jws.ParseJWTFromRequest(r)
	// logger.Info("[JWTSecuredHandler] Token: %v", headerToken)

	ctx = context.WithValue(ctx, ContextKey(JWTTokenFound), false)
	ctx = context.WithValue(ctx, ContextKey(JWTTokenValid), false)
	ctx = context.WithValue(ctx, ContextKey(JWTPolicyValid), false)

	if err == nil {
		ctx = context.WithValue(ctx, ContextKey(JWTTokenFound), true)
		// Validate token
		var validationErr error
		validated := false
		for _, v := range publicKeys {
			rsaPublic, _ := crypto.ParseRSAPublicKeyFromPEM([]byte(v))
			// logger.Info("[JWTSecuredHandler] publicKeys: %v", rsaPublic)
			validationErr = headerToken.Validate(rsaPublic, crypto.SigningMethodRS256)
			if validationErr == nil {
				// logger.Info("[JWTSecuredHandler] validationErr: %v", validationErr)
				ctx = context.WithValue(ctx, ContextKey(JWTTokenValid), true)
				if s.processPolicy {
					claims := headerToken.Claims()
					accessPolicy := claims.Get(JWTPolicyKey)

					logger.Info("[JWTSecuredHandler] headerPolicy: %v", accessPolicy)
					p, _ := json.Marshal(accessPolicy)
					logger.Info("[JWTSecuredHandler] JSON Policy: %s", p)
					policyOBJ := []policy.Policy{}
					json.Unmarshal([]byte(p), &policyOBJ)
					logger.Info("[JWTSecuredHandler] Policy: %v", policyOBJ)
					var valid bool
					for _, access := range policyOBJ {
						// logger.Info("[JWTSecuredHandler] Loop")
						// logger.Info("[JWTSecuredHandler] Acess Policy: %v", access)
						valid = s.policyValidator.Validate(access, r)
						if valid {
							// logger.Info("[JWTSecuredHandler] Policy Valid: %v", valid)
							break
						}
					}
					if s.rejectInvalidPolicy && !valid {
						api.ResponseError(w, api.Unauthorized)
						return
					}
					ctx = context.WithValue(ctx, ContextKey(JWTPolicyValid), valid)
				}
				ctx = context.WithValue(ctx, ContextKey(JWTClaim), headerToken.Claims())
				validated = true
				break
			}
		}
		if !validated {
			logger.Error("%+v", validationErr)
			if s.rejectInvalidPolicy {
				api.ResponseError(w, api.Unauthorized)
				return
			}
		}
	} else {
		if s.rejectInvalidPolicy {
			logger.Error("%+v", err.Error())
			api.ResponseError(w, api.Unauthorized)
			return
		}
	}

	s.handler.ServeHTTP(w, r.WithContext(ctx))
	return
}

// JWTProtectedHandler function
func JWTProtectedHandler(h http.Handler) JWTSecuredHandler {
	return JWTSecuredHandler{
		handler: h,
	}
}

// JWTPolicyHandler function
func JWTPolicyHandler(h http.Handler) JWTSecuredHandler {
	return JWTSecuredHandler{
		handler:         h,
		processPolicy:   true,
		policyValidator: policy.DefaultValidator{},
	}
}

// JWTPolicyProtectedHandler function
func JWTPolicyProtectedHandler(h http.Handler, v policy.Validator) JWTSecuredHandler {
	if v == nil {
		v = policy.DefaultValidator{}
	}
	return JWTSecuredHandler{
		handler:             h,
		processPolicy:       true,
		policyValidator:     v,
		rejectInvalidPolicy: true,
	}
}

//-----------------------------------------

// JWTContextValue func
func JWTContextValue(k string, r *http.Request) interface{} {
	ctx := r.Context()
	// logger.Info("[JWTContextValue] Cookie: %v", ctx.Value(ContextKey(k)))
	return ctx.Value(ContextKey(k))
}

//-----------------------------------------

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

	rsaPrivate, err := crypto.ParseRSAPrivateKeyFromPEM([]byte(privateKeys[0]))
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
