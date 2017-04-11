package auth

// var Auth = auth{}

type auth struct{}

// UserIdentity
type userIdentity struct {
	FirstName string
	LastName  string
}

func (a *auth) IsMemberOf(userid int, group string) bool {
	if group == "r" {
		return true
	}
	return false
}

// Auth function
func Auth() *auth {
	return &auth{}
}

// User function
func User() *userIdentity {
	user := userIdentity{"Leandro", "Lopez"}
	return &user
}
