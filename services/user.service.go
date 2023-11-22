package services

import "example/web-service-api-gin/models"

type UserService interface {
	FindUserById(string) (*models.DBResponse, error)
	FindUserByEmail(string) (*models.DBResponse, error)
}
