package route

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
)

func Setup(db *gorm.DB) chi.Router {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", func(writer http.ResponseWriter, request *http.Request) {
			_, err := writer.Write([]byte("OK"))
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			writer.WriteHeader(http.StatusOK)
		})
		r.Mount("/", NewBonusRouter(db))
	})

	return r
}
