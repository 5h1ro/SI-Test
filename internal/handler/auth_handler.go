package handler

import (
	"customer/internal/response"
	"customer/internal/service"
	"net/http"

	"github.com/go-chi/render"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return AuthHandler{
		authService: authService,
	}
}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	token, err := h.authService.Login()
	if err != nil {
		render.Render(w, r, response.Response[any]{
			Status:     "failed",
			StatusCode: 500,
			Message:    err.Error(),
		})
		return
	}
	res := response.Response[any]{
		Status:     "success",
		StatusCode: 200,
		Message:    "Berhasil",
		Data: map[string]interface{}{
			"token": token,
		}}
	render.Render(w, r, res)
}
