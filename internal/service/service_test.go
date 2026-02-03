package service

import (
	"testing"

	"github.com/raybman/gomaterials-slt-sandbox/internal/models"
	"github.com/raybman/gomaterials-slt-sandbox/internal/repository"
)

func TestCreateProductWithValidVendor(t *testing.T) {
	repo := repository.NewInMemoryRepository()
	svc := NewInventoryService(repo)

	// Create vendor first
	vendor := &models.Vendor{
		ID:    "v1",
		Name:  "Garden Supplies Co",
		Email: "info@gardensupplies.com",
		Phone: "555-0200",
	}
	if err := svc.CreateVendor(vendor); err != nil {
		t.Fatalf("Failed to create vendor: %v", err)
	}

	// Create product with valid vendor
	product := &models.Product{
		ID:          "p1",
		Name:        "Fertilizer",
		Description: "Organic fertilizer",
		Category:    "Soil Amendments",
		Price:       29.99,
		VendorID:    "v1",
	}
	err := svc.CreateProduct(product)
	if err != nil {
		t.Fatalf("Failed to create product: %v", err)
	}
}

func TestCreateProductWithInvalidVendor(t *testing.T) {
	repo := repository.NewInMemoryRepository()
	svc := NewInventoryService(repo)

	// Try to create product without vendor
	product := &models.Product{
		ID:          "p1",
		Name:        "Fertilizer",
		Description: "Organic fertilizer",
		Category:    "Soil Amendments",
		Price:       29.99,
		VendorID:    "nonexistent",
	}
	err := svc.CreateProduct(product)
	if err != repository.ErrNotFound {
		t.Errorf("Expected ErrNotFound, got %v", err)
	}
}

func TestCreateInventoryItemWithValidProduct(t *testing.T) {
	repo := repository.NewInMemoryRepository()
	svc := NewInventoryService(repo)

	// Create vendor and product first
	vendor := &models.Vendor{
		ID:    "v1",
		Name:  "Garden Supplies Co",
		Email: "info@gardensupplies.com",
	}
	if err := svc.CreateVendor(vendor); err != nil {
		t.Fatalf("Failed to create vendor: %v", err)
	}

	product := &models.Product{
		ID:       "p1",
		Name:     "Fertilizer",
		VendorID: "v1",
	}
	if err := svc.CreateProduct(product); err != nil {
		t.Fatalf("Failed to create product: %v", err)
	}

	// Create inventory item
	item := &models.InventoryItem{
		ID:        "i1",
		ProductID: "p1",
		Quantity:  100,
		Location:  "Warehouse A",
	}
	err := svc.CreateInventoryItem(item)
	if err != nil {
		t.Fatalf("Failed to create inventory item: %v", err)
	}
}

func TestCreateInventoryItemWithInvalidProduct(t *testing.T) {
	repo := repository.NewInMemoryRepository()
	svc := NewInventoryService(repo)

	// Try to create inventory item without product
	item := &models.InventoryItem{
		ID:        "i1",
		ProductID: "nonexistent",
		Quantity:  100,
		Location:  "Warehouse A",
	}
	err := svc.CreateInventoryItem(item)
	if err != repository.ErrNotFound {
		t.Errorf("Expected ErrNotFound, got %v", err)
	}
}
