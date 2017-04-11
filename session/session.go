package session

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/scigno/webframework/uuid4"
)

var (
	timeout = 5.0
)

// Session struct
type Session struct {
	id      string
	hash    string
	created time.Time
	expire  time.Time
}

// NewSession creates a new session object
// where d is the session duration
func NewSession() Session {
	u, _ := uuid4.New()
	// h := md5.New()
	// io.WriteString(h, u)

	ses := Session{
		id:      u,
		hash:    fmt.Sprintf("%x", md5.Sum([]byte(u))),
		created: time.Now().UTC(),
		expire:  time.Now().Add(time.Minute * time.Duration(timeout)).UTC(),
	}

	return ses
}

// SessionCreated returns the session's expiration time
func (s *Session) SessionCreated() *time.Time {
	return &s.created
}

// SessionExpiration returns the session's expiration time
func (s *Session) SessionExpiration() *time.Time {
	return &s.expire
}

// GetID returns the session's ID (UUID4)
func (s *Session) GetID() string {
	return s.id
}

// GetHash returns the session's hash in MD5
func (s *Session) GetHash() string {
	return s.hash
}

// IsValid resturn true if the session has not expired
func (s *Session) IsValid() bool {
	return true
}

// Hash return a new hash from created by UUID4+data
func (s *Session) Hash(data string) {
	s.hash = fmt.Sprintf("%x", md5.Sum([]byte(s.GetID()+data)))
}

// SetTimeout return a new hash from created by UUID4+data
func (s *Session) SetTimeout(timeout float64) {
	s.expire = s.created.Add(time.Minute * time.Duration(timeout)).UTC()
}
