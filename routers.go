/*
 * Mandatory exercise 1
 *
 * Mandatory exercse 1
 *
 * API version: 1.0.0
 */

package main

import (
	"github.com/ArneProductions/DISYS-exercise-1/endpoints"
	"github.com/ArneProductions/DISYS-exercise-1/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strconv"
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
	userRepository := repository.NewSqliteUserRepository(db)

	userController := endpoints.NewUserController(userRepository)

	workloadRepository := repository.NewSqliteWorkloadRepository(db)

	studentWorkloadRepository := repository.NewSqliteStudentWorkloadRepository(db)

	courseController := endpoints.NewCourseController()
	satisfactionController := endpoints.NewSatisfactionController()
	workloadController := endpoints.NewWorkloadController(workloadRepository, studentWorkloadRepository)

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
