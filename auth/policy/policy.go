package policy

import (
	"net/http"
	"strings"

	"github.com/scigno/webframework/logger"
)

const (
	// PolicyCREATE to POST mapping
	PolicyCREATE = "POST"
	// PolicyREAD to GET mapping
	PolicyREAD = "GET"
	// PolicyUPDATE to PUT mapping
	PolicyUPDATE = "PUT"
	// PolicyDELETE to DELETE mapping
	PolicyDELETE = "DELETE"
	// PolicyALL to ALL mapping
	PolicyALL = "*"
)

// Statement struct
type Statement struct {
	Resource []string `json:"resource"`
	Action   []string `json:"action"`
}

// Policy struct
type Policy struct {
	PolicyID  string      `json:"policy_id"`
	Version   string      `json:"version"`
	Name      string      `json:"name"`
	Statement []Statement `json:"statement"`
}

// Validator interface
type Validator interface {
	Validate(p Policy, r *http.Request) bool
}

// DefaultValidator struct
type DefaultValidator struct{}

// Validate func
func (v DefaultValidator) Validate(p Policy, r *http.Request) bool {
	logger.Info("[DefaultValidator] Validate")
	// logger.Info("[DefaultValidator] Statement: %v", p.Statement)
	for _, resources := range p.Statement {
		resourceMatched := false
		actionMatched := false
		for _, resource := range resources.Resource {
			path := r.URL.Path
			if path[len(path)-1:] == "/" {
				path = path[:len(path)-1]
			}
			if len(path) > 0 {
				path = path[1:]
			}
			logger.Info("Policy Path: %v", resource)
			logger.Info("Requested Path: %v", path)
			logger.Info("---------------------------")
			if strings.HasPrefix(path, resource) || resource == "*" {
				logger.Info("Policy Path Matched: %v", resource)
				resourceMatched = true
				break
			}
		}
		if resourceMatched {
			for _, action := range resources.Action {
				logger.Info("Policy Action: %v", action)
				if action == r.Method || action == "*" {
					actionMatched = true
					break
				}
			}
		}
		if actionMatched && resourceMatched {
			return true
		}
	}
	return false
}
