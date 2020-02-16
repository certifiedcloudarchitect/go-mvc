package controllers

import (
	"encoding/json"
	"github.com/certifiedcloudarchitect/go-mvc/services"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("user_id must be a number"))
		return
	}
	user, err := services.GetUser(userId)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		// Handle err and return to client
		return
	}
	// Return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
