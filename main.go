package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"mock/config"
	"mock/internal/handlers"
	"mock/internal/middleware"
	"mock/internal/repository"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	db, err := sql.Open("sqlite3", cfg.DBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	brandRepo := repository.NewBrandRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	orderStatusRepo := repository.NewOrderStatusRepository(db)

	brandHandler := handlers.NewBrandHandler(brandRepo)
	orderHandler := handlers.NewOrderHandler(orderRepo)
	orderStatushandler := handlers.NewOrderStatusHandler(orderStatusRepo)

	mux := http.NewServeMux()
	mux.Handle("/api/v1/brands", http.HandlerFunc(brandHandler.GetBrands))
	mux.Handle("/api/v1/orders", http.HandlerFunc(orderHandler.GetOrders))
	mux.Handle("/api/v1/orders/archives", http.HandlerFunc(orderHandler.GetOrders))
	mux.Handle("/api/v1/orders/statuses", http.HandlerFunc(orderStatushandler.GetStatuses))

	fmt.Println("Mock server at :4200")

	wrapperMux := middleware.CORSMiddleware(mux)

	if err := http.ListenAndServe(":4200", wrapperMux); err != nil {
		log.Fatal("Mock server failed:", err)
	}
}
