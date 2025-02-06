package models

import "time"

type Order struct {
	ID              string    `json:"id"`
	NumberOfProduct int       `json:"numberOfProduct"`
	Sum             float64   `json:"sum"`
	OrderNumber     string    `json:"orderNumber"`
	OrderStatusId   string    `json:"orderStatusId"`
	PaymentStatusId string    `json:"paymentStatusId"`
	IsVerified      bool      `json:"isVerified"`
	CreatedAt       time.Time `json:"createdAt"`
}

type OrderStatus struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type OrderData struct {
	Items      []Order `json:"items"`
	Page       int     `json:"page"`
	PerPage    int     `json:"perPage"`
	Total      int     `json:"total"`
	TotalPages int     `json:"totalPages"`
}

type OrderResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    OrderData `json:"data"`
}

type OrderStatusResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    []OrderStatus `json:"data"`
}
