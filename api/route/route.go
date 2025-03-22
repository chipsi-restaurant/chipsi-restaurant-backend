package route

import (
	custommiddleware "chipsiBackend/api/middleware"
	"chipsiBackend/bootstrap"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Setup(app bootstrap.Application) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	apiRouter := func(r chi.Router) {
		r.Get("/health", healthCheck)

		r.Mount("/auth/signup", NewSignupRouter(app.Db, app.Log, app.Cfg))
		r.Mount("/auth/login", NewLoginRouter(app.Db, app.Log, app.Cfg))
		r.Mount("/auth/refreshToken", NewRefreshTokenRouter(app.Db, app.Log, app.Cfg))

		// Требуют токен
		r.Group(func(r chi.Router) {
			r.Use(custommiddleware.JwtAuth(app.Cfg.App.JwtSecretKey))
			r.Mount("/bonuses", NewBonusRouter(app.Db))
			r.Mount("/users", NewUserRouter(app.Db, app.Log, app.Cfg))
		})
	}

	r.Route("/api/v1", apiRouter)

	return r
}

func healthCheck(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("OK"))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
