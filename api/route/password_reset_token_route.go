package route

import (
	"chipsiBackend/api/controller"
	"chipsiBackend/domain"
	"github.com/go-chi/chi/v5"
)

func NewPasswordRouter(passwordUsecase domain.PasswordResetTokenUsecase) chi.Router {
	pc := controller.PasswordController{
		PasswordUsecase: passwordUsecase,
	}
	r := chi.NewRouter()

	r.Post("/forgot", pc.ForgotPassword)
	r.Post("/reset", pc.ResetPassword)

	return r
}
