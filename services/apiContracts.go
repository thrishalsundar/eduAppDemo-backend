package services

import (
	models "eduapp-backend/models"
)

type ApiContracts interface {
	GetCourses() ([]*models.Course, error)
	AddCourse(models.Course) ([]*models.Course, error)
	// UpdateCourseStatus() (*models.Course, error)
}

type AuthContracts interface {
	Login(string, string) (bool, *models.User, error)
	SignUp(string, string, string) (bool, error)
}
