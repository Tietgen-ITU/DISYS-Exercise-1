/*
 * Mandatory exercise 1
 *
 * Mandatory exercse 1
 *
 * API version: 1.0.0
 */

package endpoints

import (
	"github.com/ArneProductions/DISYS-exercise-1/models"
	"github.com/ArneProductions/DISYS-exercise-1/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUser(*gin.Context)
	GetUsers(*gin.Context)
	CreateUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
}

type userController struct {
	userRepository repository.UserRepository
}

func NewUserController(r repository.UserRepository) UserController {
	return userController{
		userRepository: r,
	}
}

func (u userController) GetUser(ctx *gin.Context) {
	log.Println("{USER CONTROLLER} GetUser")

	// Relies on the middleware
	userId := ctx.MustGet("userId_int").(uint64)

	user, err := u.userRepository.GetById(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Failed retrieving user",
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "User retrieved",
		"data": user,
	})
}

func (u userController) GetUsers(ctx *gin.Context) {
	log.Println("{USER CONTROLLER} GetUsers")

	users, err := u.userRepository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Failed retrieving users",
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Users retrieved",
		"data": users,
	})
}

func (u userController) CreateUser(ctx *gin.Context) {
	log.Println("{USER CONTROLLER} Create")

	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Bad request data",
			"error": err.Error(),
		})

		return
	}

	user, err := u.userRepository.Create(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Failed saving user",
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "User saved",
		"data": user,
	})
}

func (u userController) UpdateUser(ctx *gin.Context) {
	log.Println("{USER CONTROLLER} Update")

	var user models.User

	// Relies on the middleware
	userId := ctx.MustGet("userId_int").(uint64)

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Bad request data",
			"error": err.Error(),
		})

		return
	}

	// Set the user id to the given id
	user.ID = userId

	user, err := u.userRepository.Update(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Failed updating user",
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "User updated",
		"data": user,
	})
}

func (u userController) DeleteUser(ctx *gin.Context) {
	log.Println("{USER CONTROLLER} Create")

	// Relies on the middleware
	userId := ctx.MustGet("userId_int").(uint64)

	if err := u.userRepository.Delete(userId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Could not delete user",
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "User deleted"})
}
