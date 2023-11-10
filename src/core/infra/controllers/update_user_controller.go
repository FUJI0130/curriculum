package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/FUJI0130/curriculum/src/core/app/userapp"
	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
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

	txObj, ok := c.Get("Conn")
	if !ok || txObj == nil {
		c.Error(errors.New("transaction not found"))
		return
	}

	log.Printf("CreateUserController Create: txObj: %v", txObj) //確認用
	ctxWithTx := context.WithValue(c.Request.Context(), "Conn", txObj)

	if err := ctrl.UpdateUserService.ExecUpdate(ctxWithTx, &req); err != nil {
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
