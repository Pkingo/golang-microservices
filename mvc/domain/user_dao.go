package domain

import (
	"fmt"
	"net/http"

	"github.com/pkingo/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {ID: 1, FirstName: "Philip", LastName: "Kingo", Email: "Test@test.dk"},
	}
)

// GetUser gets an user from the database
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
