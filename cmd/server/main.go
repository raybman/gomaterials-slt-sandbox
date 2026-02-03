package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/raybman/gomaterials-slt-sandbox/internal/handlers"
	"github.com/raybman/gomaterials-slt-sandbox/internal/repository"
	"github.com/raybman/gomaterials-slt-sandbox/internal/service"
)

func main() {
	// Initialize components
	repo := repository.NewInMemoryRepository()
	svc := service.NewInventoryService(repo)
	handler := handlers.NewHandler(svc)

	// Setup routes
	mux := http.NewServeMux()

	// Sellers
	mux.HandleFunc("/api/sellers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateSeller(w, r)
		case http.MethodGet:
			handler.ListSellers(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Buyers
	mux.HandleFunc("/api/buyers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateBuyer(w, r)
		case http.MethodGet:
			handler.ListBuyers(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Vendors
	mux.HandleFunc("/api/vendors", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateVendor(w, r)
		case http.MethodGet:
			handler.ListVendors(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Products
	mux.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateProduct(w, r)
		case http.MethodGet:
			handler.ListProducts(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Inventory
	mux.HandleFunc("/api/inventory", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateInventoryItem(w, r)
		case http.MethodGet:
			handler.ListInventoryItems(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/inventory/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler.UpdateInventoryQuantity(w, r)
	})

	// Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Root handler
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("Go Materials Inventory Management System\n\nAPI Endpoints:\n" +
			"  POST   /api/sellers     - Create a seller\n" +
			"  GET    /api/sellers     - List all sellers\n" +
			"  POST   /api/buyers      - Create a buyer\n" +
			"  GET    /api/buyers      - List all buyers\n" +
			"  POST   /api/vendors     - Create a vendor\n" +
			"  GET    /api/vendors     - List all vendors\n" +
			"  POST   /api/products    - Create a product\n" +
			"  GET    /api/products    - List all products\n" +
			"  POST   /api/inventory   - Create an inventory item\n" +
			"  GET    /api/inventory   - List all inventory items\n" +
			"  POST   /api/inventory/update - Update inventory quantity\n" +
			"  GET    /health          - Health check\n"))
	})

	// Start server
	port := "8080"
	fmt.Printf("Starting server on port %s...\n", port)
	fmt.Printf("Visit http://localhost:%s for API documentation\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
