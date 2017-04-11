package config

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"

	"bitbucket.com/scigno/webframework/db"
	"bitbucket.com/scigno/webframework/logger"
	"bitbucket.com/scigno/webframework/relations"
	"bitbucket.com/scigno/webframework/router"
	"bitbucket.com/scigno/webframework/session"
	"bitbucket.com/scigno/webframework/uuid4"
)

// Routes this is the routes hash
var Routes = make(map[string]func(w http.ResponseWriter, r *http.Request))

// Init function
func Init() {

	logger.Info("[init] Initializing...")

	// Routes["/"] = relations.Index
	Routes["/login"] = relations.Login
	Routes["/api"] = relations.API

	// router := router.New()
	router.AddHandler("/", relations.Index)
	router.AddHandlers(Routes)
	logger.Info("Paths set in router:")
	logger.Info(router.String())

	u, _ := uuid4.New()
	h := md5.New()
	io.WriteString(h, u+",hello")
	logger.Info("MD5: %x", h.Sum(nil))
	logger.Info("UUID4: " + u)

	db.SetResource(db.Cassandra)
	db.CassandraConnect("username", "username", "username")

	s := session.NewSession()
	s.SetTimeout(5.0)

	fmt.Println("ID: " + s.GetID())
	fmt.Println("Hash: " + s.GetHash())
	s.Hash("asfasdfsadfasdf")
	fmt.Println("Hash2: " + s.GetHash())
	fmt.Println("Created on: " + s.SessionCreated().String())
	fmt.Println("Created on: " + s.SessionCreated().Local().String())
	fmt.Println("Expires on: " + s.SessionExpiration().String())

}
