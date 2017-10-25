// Package auth relation
// and more
package auth

import (
	"encoding/json"
	"net/http"

	"github.com/scigno/webframework/logger"

	"stagezero.com/leandro/marketbin/api"
)

var maxKeys = 2

// SignKeys is the current key being used
var SignKeys = []string{"secret"}

// VerifyKeys is the fallback key
var VerifyKeys = []string{"secret"}

// SignInRSAKey is used for siging JWT tokens
type SignInRSAKey struct {
	SignInKey string `json:"key"`
}

// VerifyRSAKey is used to validate JWT tokens
type VerifyRSAKey struct {
	VerifyKey string `json:"key"`
}

// JWTSignInKeys handler.
func JWTSignInKeys(w http.ResponseWriter, r *http.Request) {

	for k, v := range r.Header {
		logger.Info("%v : %v", k, v)
	}
	var data SignInRSAKey

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		api.ResponseError(w, api.JSONValidationError)
		return
	}

	// logger.Info("SignInKey: %v", data.SignInKey)

	// decoded, err := base64.StdEncoding.DecodeString(data.SignInKey)
	// if err != nil {
	// 	fmt.Println("decode error:", err)
	// 	return
	// }
	AddSigningKey(data.SignInKey)

}

// JWTVerifyKeys handler.
func JWTVerifyKeys(w http.ResponseWriter, r *http.Request) {

	var data VerifyRSAKey

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		api.ResponseError(w, api.JSONValidationError)
		return
	}

	// logger.Info("VerifyKey: %v", data.VerifyKey)

	// decoded, err := base64.StdEncoding.DecodeString(data.VerifyKey)
	// if err != nil {
	// 	fmt.Println("decode error:", err)
	// 	return
	// }
	// fmt.Println(string(decoded))
	AddVerifyKey(data.VerifyKey)
}

// AddKeys adds a keys
func AddKeys(signkey string, veifykey string) {
	AddSigningKey(signkey)
	AddVerifyKey(veifykey)
}

// AddSigningKey adds a key to the beginign of the array
func AddSigningKey(key string) {
	found := false
	for _, v := range SignKeys {
		if v == key {
			found = true
			break
		}
	}

	if !found {
		if len(SignKeys) == 0 {
			SignKeys = append(SignKeys, key)
		} else if len(SignKeys) > 0 && len(SignKeys) < maxKeys {
			tmp := SignKeys
			SignKeys = []string{key}
			SignKeys = append(SignKeys, tmp...)

		} else {
			tmp := SignKeys[:maxKeys-1]
			SignKeys = []string{key}
			SignKeys = append(SignKeys, tmp...)
		}
	}

	// logger.Info("SignKeys: %v", SignKeys)
}

// AddVerifyKey adds a key to the beginign of the array
func AddVerifyKey(key string) {
	found := false
	for _, v := range VerifyKeys {
		if v == key {
			found = true
			break
		}
	}

	if !found {
		if len(VerifyKeys) == 0 {
			VerifyKeys = append(VerifyKeys, key)
		} else if len(VerifyKeys) > 0 && len(VerifyKeys) < maxKeys {
			tmp := VerifyKeys
			VerifyKeys = []string{key}
			VerifyKeys = append(VerifyKeys, tmp...)

		} else {
			tmp := VerifyKeys[:maxKeys-1]
			VerifyKeys = []string{key}
			VerifyKeys = append(VerifyKeys, tmp...)
		}
	}

	// logger.Info("VerifyKeys: %v", VerifyKeys)
}
