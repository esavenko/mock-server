package handlers

import (
	"encoding/json"
	"mock/internal/models"
	"net/http"
)

func BrandsHandlers(w http.ResponseWriter, r *http.Request) {
	brands := []models.Brands{
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "Electrolux", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
	}

	data := models.Data{
		Total:      50,
		PerPage:    25,
		TotalPages: 3,
		Items:      brands,
	}

	response := models.ApiResponse{
		Code:    200,
		Message: "ok",
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
