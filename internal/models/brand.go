package models

type Brand struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	URL   string `json:"url"`
}

type Data struct {
	Total      int     `json:"total"`
	PerPage    int     `json:"perPage"`
	TotalPages int     `json:"totalPages"`
	Page       int     `json:"page"`
	Items      []Brand `json:"items"`
}

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
