package controller

import (
	"chipsiBackend/domain"
	"encoding/json"
	"net/http"
)

type BonusController struct {
	BonusUsecase domain.BonusUsecase
}

func (bc *BonusController) Create(w http.ResponseWriter, r *http.Request) {
	var bonus domain.Bonus
	if err := json.NewDecoder(r.Body).Decode(&bonus); err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}
	_, err := bc.BonusUsecase.Create(r.Context(), &bonus)
	if err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(bonus)
	if err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
