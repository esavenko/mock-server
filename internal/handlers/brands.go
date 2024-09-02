package handlers

import (
	"encoding/json"
	"mock/internal/models"
	"net/http"
	"strconv"
)

func BrandsHandlers(w http.ResponseWriter, r *http.Request) {
	pageParam := r.URL.Query().Get("page")
	perPageParam := r.URL.Query().Get("pageSize")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	perPage, err := strconv.Atoi(perPageParam)
	if err != nil || perPage < 1 {
		perPage = 25
	}

	brands := []models.Brands{
		{Name: "1", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "2", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "3", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "4", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "5", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "6", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "7", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "8", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "9", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "10", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "11", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "12", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "13", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "14", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "15", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "16", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "17", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "18", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "19", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "20", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "21", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "22", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "23", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "24", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "25", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "26", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "27", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "28", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "29", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "30", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "31", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "32", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "33", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "34", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "35", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "36", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "37", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "38", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "39", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "40", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "41", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "42", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "43", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "44", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "45", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "46", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "47", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "48", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "49", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "50", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "51", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "52", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "53", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "54", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "55", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "56", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "57", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "58", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "59", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "60", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "61", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "62", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "63", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "64", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "65", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "66", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "67", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "68", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "69", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "70", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "71", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "72", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "73", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "74", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "75", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "76", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "77", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "78", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "79", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "80", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "81", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "82", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "83", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "84", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "85", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "86", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "87", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "88", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "89", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "90", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "91", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "92", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "93", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "94", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "95", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "96", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "97", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
		{Name: "98", Image: "https://placehold.jp/150x150.png", URL: "https://placehold.jp/150x150.png"},
	}

	startIndex := (page - 1) * perPage
	endIndex := startIndex + perPage

	if endIndex > len(brands) {
		startIndex = len(brands)
	}

	if endIndex > len(brands) {
		endIndex = len(brands)
	}

	pagedBrands := brands[startIndex:endIndex]

	response := map[string]interface{}{
		"totalCount":     len(brands),
		"pageSize":       perPage,
		"page":           page,
		"totalPageCount": (len(brands) + perPage - 1) / perPage,
		"data":           pagedBrands,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
