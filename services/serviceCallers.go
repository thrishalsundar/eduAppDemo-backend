package services

import (
	"context"
	"eduapp-backend/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Courses *mongo.Collection
	ctx     context.Context
}

func AppMaker(cs *mongo.Collection, c context.Context) ApiContracts {
	return &App{
		Courses: cs,
		ctx:     c,
	}
}

func (app *App) GetCourses() ([]*models.Course, error) {
	cursor, err := app.Courses.Find(app.ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var CourseArr []*models.Course
	for cursor.Next(app.ctx) {
		var item models.Course
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		CourseArr = append(CourseArr, &item)
	}

	if err := cursor.Close(app.ctx); err != nil {
		return nil, err
	}

	return CourseArr, nil
}

func (app *App) AddCourse(keladescopeDream models.Course) ([]*models.Course, error) {
	keladescopeDream.ID = primitive.NewObjectID()
	keladescopeDream.CourseID = keladescopeDream.ID.Hex()
	_, err := app.Courses.InsertOne(app.ctx, keladescopeDream)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return app.GetCourses()
}

// func (app *App) UpdateCourseStatus() (*models.Course, error) {
// 	return nil, nil
// }
