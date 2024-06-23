package handlers

/*
This is a handler file where all routes are directed to.
*/

import (
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"qonto/internal/customerrors"
	"qonto/pkg/models"
	"qonto/pkg/service"
)

const Error = "error"

type AccountController interface {
	TransferHandler(*gin.Context)
}

type controller struct {
	service.TransferService
}

func NewController(service service.TransferService) AccountController {
	return &controller{
		service,
	}
}

// TransferHandler
// @Summary transfer money between accounts
// @Description transfer money between accounts
// @Produce json
// @Accept json
// @Success 200 {array} string
// @Failure 400 {array} string
// @Failure 422 {array} string
// @Router /transfer [post]
func (c *controller) TransferHandler(ctx *gin.Context) {
	var request models.TransferRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{Error: err.Error()})
		log.Error(err)
		return
	}
	err := c.TransferService.ProcessTransfer(request)
	if err != nil {
		if errors.Is(err, customerrors.ErrInsufficientBalance) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{Error: err.Error()})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{Error: err.Error()})
		}
		log.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "transfer done successfully"})
}
