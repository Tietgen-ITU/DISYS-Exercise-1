/*
 * Mandatory exercise 1
 *
 * Mandatory exercse 1
 *
 * API version: 1.0.0
 */

package main

import (
	"fmt"
	"github.com/ArneProductions/DISYS-exercise-1/endpoints"
	"github.com/gin-gonic/gin"
	"log"
)

func SetupRouter() {
	router := gin.Default()

	setupRoutes(router)

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func setupRoutes(router *gin.Engine) {
	userController := endpoints.NewUserController()
	courseController := endpoints.NewCourseController()
	satisfactionController := endpoints.NewSatisfactionController()
	workloadController := endpoints.NewWorkloadController()

	v1 := router.Group("/v1")
	{
		v1.GET("/", Index)

		users := v1.Group("users")
		{
			users.POST("/", userController.CreateUser)
			users.DELETE("/", userController.DeleteUser)
			users.PUT("/", userController.UpdateUser)
			users.GET("/:user", userController.GetUser)
		}

		courses := v1.Group("courses")
		{
			courses.POST("/", courseController.AddCourse)
			courses.PUT("/:courseId/addStudent", courseController.AddStudentsToCourse)
			courses.DELETE("/:courseId", courseController.DeleteCourse)
			courses.GET("/", courseController.GetCourses)
			courses.POST("/:courseId/student/:studentId", courseController.RemoveStudentFromCourse)
		}

		satisfactions := v1.Group("satisfaction")
		{
			satisfactions.GET("/course/:courseId", satisfactionController.GetCourseSatisfaction)
			satisfactions.POST("/", satisfactionController.AddSatisfaction)
			satisfactions.GET("/student/:studentId", satisfactionController.GetStudentSatisfaction)
		}

		workloads := v1.Group("workload")
		{
			workloads.GET("/:courseId/:studentId", workloadController.GetStudentWorkloadFromCourse)
			workloads.POST("/", workloadController.AddWorkload)
		}
	}
}

func Index(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Hello World!")
}
