package services

import (
	"github.com/pkingo/golang-microservices/mvc/domain"
	"github.com/pkingo/golang-microservices/mvc/utils"
)

// GetUser returns the user
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
