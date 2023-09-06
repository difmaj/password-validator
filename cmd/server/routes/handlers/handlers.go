package handlers

import (
	"net/http"

	"github.com/difmaj/password-validator/internal/pkg/entities"
	"github.com/difmaj/password-validator/internal/pkg/repositories/password"
	"github.com/difmaj/password-validator/internal/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// PasswordHandler is the struct that contains the service instances.
type PasswordHandler struct {
	PasswordService *services.PasswordService
}

// NewPasswordHandler returns a new Handler instance.
func NewPasswordHandler() *PasswordHandler {
	return &PasswordHandler{
		PasswordService: services.NewPasswordService(
			password.NewRepository(),
		),
	}
}

// Verify is the handler function that validates the password.
func (h *PasswordHandler) Verify(ctx *gin.Context) {

	request := new(entities.Request)
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
