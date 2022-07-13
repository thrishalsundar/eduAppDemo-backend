package main

import (
	"context"
	"eduapp-backend/controllers"
	"eduapp-backend/database"
	"eduapp-backend/services"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	port         string
	server       *gin.Engine
	forRoutes    *controllers.Remote
	userRoutes   *controllers.UserRemote
	ctx          context.Context
	Courses      *mongo.Collection
	Users        *mongo.Collection
	client       *mongo.Client
	servicer     services.ApiContracts
	authServicer services.AuthContracts
)

func init() {
	client, ctx = database.DBSetup()
	Courses = client.Database("EduApp").Collection("Courses")
	Users = client.Database("EduApp").Collection("Users")
	servicer = services.AppMaker(Courses, ctx)
	authServicer = services.AuthAppContruct(Users, ctx)
	forRoutes = controllers.RemoteMaker(servicer)
	userRoutes = controllers.UserRemConstruct(authServicer)
	port = os.Getenv("PORT")
	server = gin.Default()
	server.Use(cors.Default())
}

func main() {
	defer client.Disconnect(ctx)
	basePath := server.Group("/apis")
	forRoutes.CoursesRoutes(basePath)
	userRoutes.UserRoutes(basePath)
	server.Run(":" + port)
}
