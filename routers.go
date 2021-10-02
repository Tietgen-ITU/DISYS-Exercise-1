/*
 * Mandatory exercise 1
 *
 * Mandatory exercse 1
 *
 * API version: 1.0.0
 */

package main

import (
	"log"
	"strconv"

	"github.com/ArneProductions/DISYS-exercise-1/endpoints"
	"github.com/ArneProductions/DISYS-exercise-1/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) {
	router := gin.Default()

	setupRoutes(router, db)

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func convertToUInt(name string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		val := ctx.Params.ByName(name)

		if val == "" {
			return
		}

		conv, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			ctx.Error(err)
		}

		ctx.Set(name+"_int", conv)
		ctx.Next()
	}
}

func setupRoutes(router *gin.Engine, db *gorm.DB) {
	// Define repositories
	userRepository := repository.NewSqliteUserRepository(db)
	satisfactionRepository := repository.NewSqliteSatisfactionRepository(db)
	courseRepository := repository.NewSqliteCourseRepository(db)

	// Create controllers
	userController := endpoints.NewUserController(userRepository)

	satisfactionController := endpoints.NewSatisfactionController(satisfactionRepository)
	courseController := endpoints.NewCourseController(courseRepository)
	workloadController := endpoints.NewWorkloadController()

	v1 := router.Group("/v1")
	{
		users := v1.Group("users")
		{
			users.POST("/", userController.CreateUser)
			users.GET("/", userController.GetUsers)

			usersWithId := users.Group(":userId")
			{
				usersWithId.Use(convertToUInt("userId"))
				usersWithId.PUT("/", userController.UpdateUser)
				usersWithId.GET("/", userController.GetUser)
				usersWithId.DELETE("/", userController.DeleteUser)
			}
		}

		courses := v1.Group("course")
		{
			courses.POST("/", courseController.AddCourse)
			courses.GET("/", courseController.GetCourses)

			coursesWithId := courses.Group(":courseId")
			{
				coursesWithId.Use(convertToUInt("courseId"))
				coursesWithId.PUT("/addStudent", courseController.AddStudentsToCourse)
				coursesWithId.DELETE("/", courseController.DeleteCourse)

				coursesWithIdAndStudentId := coursesWithId.Group("student/:studentId")
				{
					coursesWithIdAndStudentId.Use(convertToUInt("studentId"))
					coursesWithIdAndStudentId.DELETE("/", courseController.RemoveStudentFromCourse)
				}
			}
		}

		satisfactions := v1.Group("satisfaction")
		{
			satisfactions.Use(convertToUInt("courseId"))
			satisfactions.Use(convertToUInt("studentId"))
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
