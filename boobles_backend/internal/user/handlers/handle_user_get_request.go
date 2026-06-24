package handlers

import "net/http"

// TODO: implement a new middleware for this!!
// These are funcs, that only our frontend can access!!

// Handles the getting the user by the accesstoken val
func HandleGettingUserByTokenVal(w http.ResponseWriter, r *http.Request)

// Handles getting the user by the user id
func HandleGettingUserById(w http.ResponseWriter, r *http.Request)

// Handles getting the user by tenant and user name
func HandleGettingUserByTenantIdAndUserName(w http.ResponseWriter, r *http.Request)

// Handles the request for checking if a user has a tenant
func HandleHasUserATenant(w http.ResponseWriter, r *http.Request)
