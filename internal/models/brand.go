package models

type Brands struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	URL   string `json:"url"`
}

type Data struct {
	Total      int      `json:"total"`
	PerPage    int      `json:"perPage"`
	TotalPages int      `json:"totalPages"`
	Items      []Brands `json:"items"`
}
