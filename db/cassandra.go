package db

import (
	"fmt"
	"log"

	"github.com/scigno/webframework/logger"

	"github.com/gocql/gocql"
)

const (
	// ReplicationFactor is the default replication factor in Cassandra
	ReplicationFactor = 1
)

// CassandraConnect to the ip using the username and password provided
func CassandraConnect(ip string, u string, p string) (DAL, error) {
	// connect to the cluster
	logger.Info("[cassandra] Initializing...")
	cluster := gocql.NewCluster("192.168.60.100")
	// cluster.Authenticator := gocql.PasswordAuthenticator {
	//
	// }
	cluster.SslOpts = &gocql.SslOptions{
		CertPath: "/Users/c10002a/Development/GoCode/src/github.com/scigno/webframework/certs/rootCa.crt",
		KeyPath:  "/Users/c10002a/Development/GoCode/src/github.com/scigno/webframework/certs/rootCa.key",
	}
	cluster.Keyspace = "hiera"
	session, err := cluster.CreateSession()
	if err != nil {
		logger.Info(err.Error())
		return nil, nil
	}
	defer session.Close()

	// var id gocql.UUID
	var source, key, name string
	iter := session.Query(`select * from hiera.value_by_source where source=? and key=?`, "global", "name").Iter()
	for iter.Scan(&source, &key, nil, nil, nil, nil, &name) {
		fmt.Println(source, key, name)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	return nil, nil
}
