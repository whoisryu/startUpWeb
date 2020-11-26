package handler

import (
	"BWAPROJECT/helper"
	"BWAPROJECT/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//UserHandler ...
type UserHandler struct {
	userService user.ServiceInterface
}

//NewUserHandler ...
func NewUserHandler(service user.ServiceInterface) *UserHandler {
	return &UserHandler{service}
}

//RegisterUser ...
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	fmt.Print(err)

	if err != nil {
		errors := helper.FormatError(err)
		errMessage := gin.H{"errors": errors}

		responseBad := helper.APIResponse(http.StatusUnprocessableEntity, "Faild register user", "Bad Request", errMessage)
		c.JSON(400, responseBad)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		responseBad := helper.APIResponse(http.StatusUnprocessableEntity, "Faild register user", "Bad Request", err.Error())
		c.JSON(400, responseBad)
		return
	}

	formatter := user.FormatterUser(newUser, "")

	responseSuccess := helper.APIResponse(200, "Success register user", "success", formatter)

	c.JSON(200, responseSuccess)
}
