package handlers

import (
	"net/http"

	"github.com/difmaj/password-validator-challenge/internal/pkg/password"
	"github.com/difmaj/password-validator-challenge/internal/pkg/password/entity"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Handler struct {
	PasswordService password.UseCase
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Verify(ctx *gin.Context) {

	request := new(entity.Request)
	if err := ctx.ShouldBindBodyWith(request, binding.JSON); err != nil {
		ctx.Error(err)
		return
	}

	response, err := h.PasswordService.Verify(request)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
