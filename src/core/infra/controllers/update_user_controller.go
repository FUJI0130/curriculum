package controllers

import (
	"log"
	"net/http"

	"github.com/FUJI0130/curriculum/src/core/app/userapp"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type UpdateUserController struct {
	UpdateUserService *userapp.UpdateUserAppService
}

func NewUpdateUserController(s *userapp.UpdateUserAppService) *UpdateUserController {
	return &UpdateUserController{UpdateUserService: s}
}

func (ctrl UpdateUserController) Update(c *gin.Context) {
	var req userapp.UpdateUserRequestData
	if err := c.BindJSON(&req); err != nil {
		log.Printf("Error while binding JSON: %v", err)
		c.Error(customerrors.WrapUnprocessableEntityError(err, "JSON binding error"))
		return
	}

	ctx := c.Request.Context()
	txObj, ok := ctx.Value("Conn").(*sqlx.Tx)
	if !ok || txObj == nil {
		c.Error(errors.New("transaction not found"))
		return
	}

	if err := ctrl.UpdateUserService.ExecUpdate(c.Request.Context(), &req); err != nil {
		if customErr, ok := err.(customerrors.BaseError); ok {
			c.Status(customErr.StatusCode())
		} else {
			c.Status(http.StatusInternalServerError)
		}

		c.Error(customerrors.WrapUnprocessableEntityError(err, "UpdateUserService ExecUpdate error"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Updated successfully"})
}
