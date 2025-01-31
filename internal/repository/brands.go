package repository

import (
	"database/sql"
	"math"
	"mock/internal/models"
)

type BrandRepository struct {
	db *sql.DB
}

func NewBrandRepository(db *sql.DB) *BrandRepository {
	return &BrandRepository{db: db}
}

func (r *BrandRepository) GetAll(page, limit int) (models.Data, error) {
	offset := (page - 1) * limit

	rows, err := r.db.Query("SELECT id, name, image, url FROM brands LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return models.Data{}, err
	}
	defer rows.Close()

	var brands []models.Brand
	for rows.Next() {
		var brand models.Brand
		if err := rows.Scan(&brand.ID, &brand.Name, &brand.Image, &brand.URL); err != nil {
			return models.Data{}, err
		}
		brands = append(brands, brand)
	}

	var total int
	err = r.db.QueryRow("SELECT COUNT(*) FROM brands").Scan(&total)
	if err != nil {
		return models.Data{}, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return models.Data{
		Total:      total,
		PerPage:    limit,
		TotalPages: totalPages,
		Page:       page,
		Items:      brands,
	}, nil
}
