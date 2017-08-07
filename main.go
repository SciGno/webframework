package main

import (
	"encoding/json"
	"log"
	"net/http"

	// "github.com/julienschmidt/httprouter"
	"github.com/scigno/webframework/httprouter"
)

// This is the main function
func main() {

	s := httprouter.New()
	s.HandleFuncGET("/user/profile/{profile_id}", UserStatus)
	s.HandleFuncGET("/status/", APIStatus)

	log.Fatal(http.ListenAndServe(":8080", s))

}

// UserStatus This is the API function for the handler.
func UserStatus(w http.ResponseWriter, r *http.Request) {
	// logger.Info("UserStatus: Serving...")
	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"first_name": "Someone",
			"last_name":  "else",
		},
	)
	return
}

// APIStatus This is the API function for the handler.
func APIStatus(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"Status": "GOOD",
			"CODE":   200,
		},
	)
	return
}
