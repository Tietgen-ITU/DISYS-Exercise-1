/*
 * Mandatory exercise 1
 *
 * Mandatory exercse 1
 *
 * API version: 1.0.0
 */

package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(*gin.Context)
	DeleteUser(*gin.Context)
	UpdateUser(*gin.Context)
	GetUser(*gin.Context)
}

type userController struct {
}

func NewUserController() UserController {
	return userController{}
}

func (u userController) CreateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "Ok"})
}

func (u userController) DeleteUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "Ok"})
}

func (u userController) UpdateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "Ok"})
}

func (u userController) GetUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "Ok"})
}
