package userHandler

import (
	"net/http"
	"serverhealthcarepanel/dto"
	userService "serverhealthcarepanel/services/user"
	"serverhealthcarepanel/utils/code"
	"serverhealthcarepanel/utils/response"

	"github.com/labstack/echo/v4"
)

// CreateUser
// @Summary create user
// @Description Create new user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Tags User
// @Param payload body dto.AddUser true "create new user"
// @Success 200 {object} response.Struct
// @Failure 400 {object} response.Struct "wrong request parameter"
// @Failure 500 {object} response.Struct
// @Router /user [post]
func CreateUser(ctx echo.Context) error {
	newUser := new(dto.AddUser)

	if err := ctx.Bind(&newUser); err != nil {
		return response.Error(ctx, http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
	}

	if err := ctx.Validate(*newUser); err != nil {
		return response.Error(ctx, http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
	}

	if err := userService.CreateUser(newUser); err != nil {
		return response.Error(ctx, http.StatusOK, code.ErrorFailedAddNewUser, code.GetMsg(code.ErrorFailedAddNewUser), err.Error())
	}

	return response.Success(ctx, nil)
}
