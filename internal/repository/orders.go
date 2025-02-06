package repository

import (
	"database/sql"
	"math"
	"mock/internal/models"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) GetAll(page, limit int, isArchive bool) (models.OrderData, error) {
	offset := (page - 1) * limit
	tableName := "orders"
	if isArchive {
		tableName = "orders_archive"
	}

	query := `
		SELECT
			id, number_of_products, sum, order_number,
			order_status_id, payment_status_id, is_verified, created_at
		FROM ` + tableName + `
		LIMIT ? OFFSET ?`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return models.OrderData{}, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		err := rows.Scan(
			&order.ID, &order.NumberOfProduct, &order.Sum, &order.OrderNumber,
			&order.OrderStatusId, &order.PaymentStatusId, &order.IsVerified, &order.CreatedAt,
		)
		if err != nil {
			return models.OrderData{}, err
		}
		orders = append(orders, order)
	}

	countQuery := "SELECT COUNT(*) FROM " + tableName
	var total int
	err = r.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return models.OrderData{}, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return models.OrderData{
		Items:      orders,
		Page:       page,
		PerPage:    limit,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}
