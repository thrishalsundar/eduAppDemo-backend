package controllers

import (
	"eduapp-backend/models"
	"eduapp-backend/services"

	"github.com/gin-gonic/gin"
)

type Remote struct {
	Calls services.ApiContracts
}

func RemoteMaker(s services.ApiContracts) *Remote {
	return &Remote{
		Calls: s,
	}
}

func (apiApp *Remote) CoursesRoutes(incomingRoutes *gin.RouterGroup) {
	routes := incomingRoutes.Group("/courses")
	routes.GET("/allcourses", apiApp.GetCourses)
	routes.POST("/addcourse", apiApp.AddCourse)
	// routes.PUT("/updatecourse", apiApp.UpdateCourseStatus)

}

func (apiApp *Remote) GetCourses(c *gin.Context) {
	data, err := apiApp.Calls.GetCourses()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		panic(err)
	}

	c.JSON(200, gin.H{"data": data})

}

func (apiApp *Remote) AddCourse(c *gin.Context) {
	var CourseFromApi models.Course

	if err := c.BindJSON(&CourseFromApi); err != nil {
		c.JSON(403, gin.H{"message": "Invalid Creds", "error": err})
		panic(err)
	}

	data, err := apiApp.Calls.AddCourse(CourseFromApi)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Insert Error or geterr",
		})
	}

	c.JSON(200, gin.H{
		"data": data,
	})
}
