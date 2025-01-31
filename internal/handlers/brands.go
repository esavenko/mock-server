package handlers

import (
	"encoding/json"
	"mock/internal/models"
	"mock/internal/repository"
	"net/http"
	"strconv"
)

type BrandHandler struct {
	repo *repository.BrandRepository
}

func NewBrandHandler(repo *repository.BrandRepository) *BrandHandler {
	return &BrandHandler{repo: repo}
}

func (h *BrandHandler) GetBrands(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 {
		limit = 25
	}

	data, err := h.repo.GetAll(page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.APIResponse{
		Code:    200,
		Message: "ok",
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
