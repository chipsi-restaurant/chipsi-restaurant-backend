package route

import (
	custommiddleware "chipsiBackend/api/middleware"
	"chipsiBackend/bootstrap"
	"chipsiBackend/setup"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

func Setup(app bootstrap.Application) chi.Router {
	r := chi.NewRouter()

	graph := setup.BuildGraph(app)

	adminMiddleware := custommiddleware.IsAdmin(graph.UCs.User)

	r.Use(custommiddleware.SetJSONContentType)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	apiRouter := func(r chi.Router) {
		r.Get("/health", healthCheck)

		r.Mount("/auth/signup", NewSignupRouter(graph.UCs.Signup, app.Log, app.Cfg))
		r.Mount("/auth/login", NewLoginRouter(graph.UCs.Login, app.Log, app.Cfg))
		r.Mount("/auth/refreshToken", NewRefreshTokenRouter(graph.UCs.RefreshToken, app.Log, app.Cfg))

		// Требуют токен
		r.Group(func(r chi.Router) {
			r.Use(custommiddleware.JwtAuth(app.Cfg.App.JwtSecretKey))
			r.Mount("/bonuses", NewBonusRouter(graph.UCs.Bonus))
			r.Mount("/categories", NewCategoryRouter(graph.UCs.Category, adminMiddleware))
			r.Mount("/menuItems", NewMenuItemRouter(graph.UCs.MenuItem, adminMiddleware))
			r.Mount("/giftCertificates", NewGiftCertificateRouter(graph.UCs.GiftCertificate))
			r.Mount("/users", NewUserRouter(graph.UCs.User, app.Log, app.Cfg))
		})
	}

	r.Route("/api/v1", apiRouter)

	return r
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := map[string]interface{}{
		"status":    "ok",
		"service":   "chipsiBackend",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	}

	_ = json.NewEncoder(w).Encode(resp)
}
