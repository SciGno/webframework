package db

import "github.com/gocql/gocql"

// There are the supported databases
const (
	// Cassandra integer Value
	Cassandra = 1
	// MySQL integer Value
	MySQL = 2
	// Oracle integer Value
	Oracle = 3
)

// Rosource interface implemented by connectord
type Rosource interface {
	Connect()
	GetSession()
}

// DAL is the Database Abstraction Layer
type DAL interface {
	Connect(u string, p string) *gocql.Session
	Disconnect() bool
	Select(q string) (string, error)
	CreateTable(t string, d string) (bool, error)
	DropTable(t string) (bool, error)
}

// SetResource makes a connection to the specified resource with provided
// username and password
func SetResource(resource int) (*gocql.Session, error) {
	return nil, nil
}

// Connect makes a connection to the specified resource with provided
// username and password
func Connect() (*gocql.Session, error) {
	return nil, nil
}
