package controller

import (
	"chipsiBackend/domain"
	"chipsiBackend/pkg/httpErrors"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"mime/multipart"
	"net/http"
	"strconv"
)

type MenuItemController struct {
	MenuItemUsecase domain.MenuItemUsecase
}

func (mc *MenuItemController) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // до 10MB
	if err != nil {
		httpErrors.JSONError(w, "invalid multipart data", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		httpErrors.JSONError(w, "file not found", http.StatusBadRequest)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			httpErrors.JSONError(w, "upload failed:"+err.Error(), http.StatusInternalServerError)
			return
		}
	}(file)

	var menuItem domain.MenuItem
	menuItemJson := r.Form.Get("metadata")
	if menuItemJson == "" {
		httpErrors.JSONError(w, "empty metadata", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal([]byte(menuItemJson), &menuItem); err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	newMenuItem, err := mc.MenuItemUsecase.Create(r.Context(), &menuItem, file, handler)
	if err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(newMenuItem)

	if err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (mc *MenuItemController) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		httpErrors.JSONError(w, "id must be not null", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		httpErrors.JSONError(w, "id must be number", http.StatusBadRequest)
		return
	}

	menuItem, err := mc.MenuItemUsecase.GetByID(r.Context(), id)

	if err != nil {
		httpErrors.JSONError(w, "not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(*menuItem); err != nil {
		httpErrors.JSONError(w, "can't encode menuItem to json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (mc *MenuItemController) GetAll(w http.ResponseWriter, r *http.Request) {

	categoryId := r.URL.Query().Get("categoryId")

	if categoryId == "" {
		httpErrors.JSONError(w, "query-param categoryId must be not null", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(categoryId, 10, 64)
	if err != nil {
		httpErrors.JSONError(w, "id must be number", http.StatusBadRequest)
		return
	}

	menuItems, err := mc.MenuItemUsecase.GetAllByCategory(r.Context(), id)

	if err != nil {
		httpErrors.JSONError(w, "category not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(menuItems); err != nil {
		httpErrors.JSONError(w, "can't encode menuItems to json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (mc *MenuItemController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		httpErrors.JSONError(w, "id must be not null", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		httpErrors.JSONError(w, "id must be number", http.StatusBadRequest)
		return
	}

	err = mc.MenuItemUsecase.Delete(r.Context(), id)
	if err != nil {
		httpErrors.JSONError(w, "Menu item not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
