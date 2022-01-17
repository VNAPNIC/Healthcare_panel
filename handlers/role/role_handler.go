package roleHandler

import (
	"net/http"
	"serverhealthcarepanel/services/role"
	"serverhealthcarepanel/utils/code"
	"serverhealthcarepanel/utils/response"

	"github.com/labstack/echo/v4"
)

// @Summary Create role
// @Description Create role
// @Accept json
// @Produce json
// @Tags Role
// @Param role_id path int true "role_id"
// @Param payload body roleService.CreateRoleStruct true "YES"、
// @Success 200 {object} response.Struct
// @Failure 400 {object} response.Struct "wrong request parameter"
// @Router /role [post]
func CreateRole(ctx echo.Context) error {
	newRole := new(role.CreateRoleStruct)

	if err := ctx.Bind(&newRole); err != nil {
		return response.Error(ctx, http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())
	}

	if err := ctx.Validate(*newRole); err != nil {
		return response.Error(ctx, http.StatusBadRequest, code.InvalidParams, code.GetMsg(code.InvalidParams), err.Error())

	}

	if err := role.CreateRole(newRole); err != nil {
		return response.Error(ctx, http.StatusInternalServerError, code.ErrorFailedAddNewRole, code.GetMsg(code.ErrorFailedAddNewRole), err.Error())
	}

	return response.Success(ctx, nil)
}