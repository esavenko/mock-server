package repository

import (
	"database/sql"
	"mock/internal/models"
)

type OrderStatusRepository struct {
	db *sql.DB
}

func NewOrderStatusRepository(db *sql.DB) *OrderStatusRepository {
	return &OrderStatusRepository{db: db}
}

func (r *OrderStatusRepository) GetAll() ([]models.OrderStatus, error) {
	rows, err := r.db.Query("SELECT id, name, description FROM orders_statuses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []models.OrderStatus
	for rows.Next() {
		var status models.OrderStatus
		if err := rows.Scan(&status.ID, &status.Name, &status.Description); err != nil {
			return nil, err
		}
		statuses = append(statuses, status)
	}

	return statuses, nil
}
