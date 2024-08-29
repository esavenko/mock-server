package main

import (
	"fmt"
	"log"
	"mock/internal/handlers"
	"mock/internal/middleware"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/brands", handlers.BrandsHandlers)

	fmt.Println("Mock server at :4200")

	wrapperMux := middleware.CORSMiddleware(mux)

	if err := http.ListenAndServe(":4200", wrapperMux); err != nil {
		log.Fatal("Mock server failde:", err)
	}
}
