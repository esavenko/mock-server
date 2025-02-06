package handlers

import (
	"encoding/json"
	"mock/internal/models"
	"mock/internal/repository"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	repo *repository.OrderRepository
}

func NewOrderHandler(repo *repository.OrderRepository) *OrderHandler {
	return &OrderHandler{repo: repo}
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 {
		limit = 10
	}

	isArchive := r.URL.Path == "/api/v1/orders/archives"

	data, err := h.repo.GetAll(page, limit, isArchive)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.OrderResponse{
		Code:    200,
		Message: "ok",
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
