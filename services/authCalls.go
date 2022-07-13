package services

import (
	"context"
	"eduapp-backend/models"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthApp struct {
	Users *mongo.Collection
	ctx   context.Context
}

func AuthAppContruct(userColl *mongo.Collection, ct context.Context) AuthContracts {
	return &AuthApp{
		Users: userColl,
		ctx:   ct,
	}
}

func (app *AuthApp) Login(email string, password string) (bool, *models.User, error) {
	var foundUser *models.User
	err := app.Users.FindOne(app.ctx, bson.M{"email": email}).Decode(&foundUser)
	if err != nil {
		fmt.Println("error")
		return false, nil, err
	}

	if foundUser.Password != password {
		passWrongErr := errors.New("Wrong Pass")
		return false, nil, passWrongErr
	}

	return true, foundUser, nil
}

func (app *AuthApp) SignUp(username string, email string, password string) (bool, error) {
	count, err := app.Users.CountDocuments(app.ctx, bson.M{"email": email})
	if err != nil {
		return false, errors.New("ethoerror")
	}
	if count > 0 {
		userExistsErr := errors.New("User already exists")
		return false, userExistsErr
	}
	var givenUser models.User
	var newId primitive.ObjectID
	newId = primitive.NewObjectID()
	givenUser.ID = newId
	givenUser.UserID = newId.Hex()
	givenUser.UserName = username
	givenUser.Email = email
	givenUser.Password = password
	givenUser.Courses = make([]models.Course, 0)
	fmt.Println(givenUser)

	_, err = app.Users.InsertOne(app.ctx, givenUser)
	if err != nil {
		panic(err)
	}
	return true, nil
}
