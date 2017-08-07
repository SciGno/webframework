package auth

// import "github.com/dgrijalva/jwt-go"

type auth struct{}

// Auth variable
var Auth = auth{}

// Provider interface
type Provider interface {
	ValidateCredentials(user string, password string) bool
}

// JWTProvider interface
type JWTProvider interface {
	ValidateJWT(jwt interface{})
	CreateJWTToken() interface{}
}

// RBAC interface
type RBAC interface {
	AddGroup(group string) bool
	AddToGroup(users []string)
}

func (a *auth) IsMemberOf(userid int, group string) bool {
	if group == "r" {
		return true
	}
	return false
}
