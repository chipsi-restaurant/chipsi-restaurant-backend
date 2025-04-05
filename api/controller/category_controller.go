package controller

import (
	"chipsiBackend/domain"
	"chipsiBackend/pkg/httpErrors"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"strings"
)

type CategoryController struct {
	CategoryUsecase domain.CategoryUsecase
}

func (cc *CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	var category domain.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	_, err := cc.CategoryUsecase.Create(r.Context(), &category)

	if err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(category)

	if err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (cc *CategoryController) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, `{"error": "id must be not null"}`, http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, `{"error": "id must be number"}`, http.StatusBadRequest)
		return
	}

	category, err := cc.CategoryUsecase.GetByID(r.Context(), id)

	if err != nil {
		http.Error(w, `{"error": "not found"}`, http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(*category); err != nil {
		http.Error(w, `{"error": "can't encode category to json'"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (cc *CategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
	include := r.URL.Query().Get("includeMenuItems")
	if include == "" || (include != "true" && include != "false") {
		http.Error(w, `{"error": "invalid query param"}`, http.StatusBadRequest)
		return
	}

	categories, err := cc.CategoryUsecase.GetAll(r.Context(), include == "true")

	if err != nil || len(categories) == 0 {
		http.Error(w, `{"error": "not found"}`, http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(categories); err != nil {
		http.Error(w, `{"error": "can't encode category to json'"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (cc *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, `{"error": "id must be not null"}`, http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, `{"error": "id must be number"}`, http.StatusBadRequest)
		return
	}

	err = cc.CategoryUsecase.Delete(r.Context(), id)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "SQLSTATE 23503"):
			httpErrors.JSONError(w, "can't delete: item has dependencies", http.StatusConflict)
			return
		case strings.Contains(err.Error(), "record not found"):
			httpErrors.JSONError(w, "category not found", http.StatusNotFound)
			return
		default:
			httpErrors.JSONError(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
