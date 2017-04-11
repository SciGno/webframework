package session

import "fmt"

// CreateSession creates a new session from and maps it to a username
func CreateSession() (Session, error) {
	s := NewSession()
	return s, nil
}

// GetSessionByID retreives the session by ID
func GetSessionByID(id string) (string, error) {
	return "nil", fmt.Errorf("Session not found")
}

// GetSessionByHash retreives the session by Hash
func GetSessionByHash(hash string) (string, error) {
	return "nil", fmt.Errorf("Session not found")
}
