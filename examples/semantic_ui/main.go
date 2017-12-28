// Package main functions
package main

import (
	"log"
	"net/http"

	"github.com/scigno/webframework/auth"

	"github.com/SciGno/webframework/httprouter"
	"github.com/scigno/webframework/examples/semantic_ui/relations"
)

// This is the main function
func main() {

	// userPassword1 := "some user-provided password"

	// // Generate "hash" to store from user password
	// hash, err := bcrypt.GenerateFromPassword([]byte(userPassword1), bcrypt.DefaultCost)
	// if err != nil {
	// 	// TODO: Properly handle error
	// 	log.Fatal(err)
	// }
	// fmt.Println("Hash to store:", string(hash))
	// // Store this "hash" somewhere, e.g. in your database

	// // After a while, the user wants to log in and you need to check the password he entered
	// userPassword2 := "some user-provided password"
	// hashFromDatabase := hash

	// // Comparing the password with the hash
	// if err := bcrypt.CompareHashAndPassword(hashFromDatabase, []byte(userPassword2)); err != nil {
	// 	// TODO: Properly handle error
	// 	log.Fatal(err)
	// }

	// fmt.Println("Password was correct!")

	s := httprouter.New()
	fs := http.FileServer(http.Dir("static"))
	s.HandleStatic("/static", http.StripPrefix("/static", fs))
	// s.HandleFuncGET("/", relations.Index)

	// Used to add RSA keys to the framework
	s.HandleFuncPOST("/v1/jwt/keys", auth.JWTKeys)

	// create a protectd handler with a redirect string
	s.Handle("/index", auth.JWTProtectedHandler(relations.Index))
	s.Handle("/login", auth.JWTProtectedHandler(relations.Login))
	s.Handle("/register", auth.JWTProtectedHandler(relations.Register))
	s.Handle("/registered", auth.JWTProtectedHandler(relations.Registered))

	// s.HandleFuncGET("/", relations.Index)
	// s.HandlePOST("/v1/jwt/form/login",
	// 	auth.JWTFormLogin(
	// 		relations.APILoginVerifier,
	// 		httprouter.Func2Handler(relations.Index),
	// 	),
	// )

	// s.Handle(
	// 	"/login",
	// 	auth.JWTFormLogin( // implements JWT token
	// 		relations.UserAuthenticator{},            // process user authentication
	// 		relations.UserValidator{},                // validates JWT claims
	// 		httprouter.Func2Handler(relations.Login), // main handler
	// 		"/settings",                              // redirect URL after successfull authentication and validation
	// 		10,                                       // JWT Timeout in seconds
	// 	),
	// )

	// s.HandlePOST("/login", auth.JWTFormLogin(relations.LoginVerifier, httprouter.Func2Handler(relations.Login)))

	log.Fatal(http.ListenAndServe("localhost:8888", s))

}
