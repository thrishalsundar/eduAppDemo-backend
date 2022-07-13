package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	CourseID    string             `json:"course_id"`
	CourseName  string             `json:"course_name"`
	Price       int64              `json:"price"`
	Duration    int64              `json:"duration"`
	Description string             `json:"description"`
	Status      bool               `json:"status"`
}

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	UserID   string             `json:"user_id"`
	UserName string             `json:"user_name"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Courses  []Course           `bson:"courses" json:"courses"`
}
