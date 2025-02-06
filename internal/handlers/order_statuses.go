package handlers

import (
	"encoding/json"
	"mock/internal/models"
	"mock/internal/repository"
	"net/http"
)

type OrderStatusHandler struct {
	repo *repository.OrderStatusRepository
}

func NewOrderStatusHandler(repo *repository.OrderStatusRepository) *OrderStatusHandler {
	return &OrderStatusHandler{repo: repo}
}

func (h *OrderStatusHandler) GetStatuses(w http.ResponseWriter, r *http.Request) {
	statuses, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.OrderStatusResponse{
		Code:    200,
		Message: "ok",
		Data:    statuses,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
