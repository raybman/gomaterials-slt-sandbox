package repository

import (
	"testing"

	"github.com/raybman/gomaterials-slt-sandbox/internal/models"
)

func TestCreateAndGetSeller(t *testing.T) {
	repo := NewInMemoryRepository()

	seller := &models.Seller{
		ID:    "s1",
		Name:  "John Doe",
		Email: "john@example.com",
		Phone: "555-0100",
	}

	err := repo.CreateSeller(seller)
	if err != nil {
		t.Fatalf("Failed to create seller: %v", err)
	}

	retrieved, err := repo.GetSeller("s1")
	if err != nil {
		t.Fatalf("Failed to get seller: %v", err)
	}

	if retrieved.Name != seller.Name {
		t.Errorf("Expected name %s, got %s", seller.Name, retrieved.Name)
	}
}

func TestCreateDuplicateSeller(t *testing.T) {
	repo := NewInMemoryRepository()

	seller := &models.Seller{
		ID:    "s1",
		Name:  "John Doe",
		Email: "john@example.com",
		Phone: "555-0100",
	}

	err := repo.CreateSeller(seller)
	if err != nil {
		t.Fatalf("Failed to create seller: %v", err)
	}

	err = repo.CreateSeller(seller)
	if err != ErrAlreadyExists {
		t.Errorf("Expected ErrAlreadyExists, got %v", err)
	}
}

func TestListSellers(t *testing.T) {
	repo := NewInMemoryRepository()

	sellers := []*models.Seller{
		{ID: "s1", Name: "Seller 1", Email: "s1@example.com", Phone: "555-0101"},
		{ID: "s2", Name: "Seller 2", Email: "s2@example.com", Phone: "555-0102"},
	}

	for _, seller := range sellers {
		if err := repo.CreateSeller(seller); err != nil {
			t.Fatalf("Failed to create seller: %v", err)
		}
	}

	list, err := repo.ListSellers()
	if err != nil {
		t.Fatalf("Failed to list sellers: %v", err)
	}

	if len(list) != 2 {
		t.Errorf("Expected 2 sellers, got %d", len(list))
	}
}

func TestCreateAndGetVendor(t *testing.T) {
	repo := NewInMemoryRepository()

	vendor := &models.Vendor{
		ID:      "v1",
		Name:    "Garden Supplies Co",
		Email:   "info@gardensupplies.com",
		Phone:   "555-0200",
		Address: "123 Garden St",
	}

	err := repo.CreateVendor(vendor)
	if err != nil {
		t.Fatalf("Failed to create vendor: %v", err)
	}

	retrieved, err := repo.GetVendor("v1")
	if err != nil {
		t.Fatalf("Failed to get vendor: %v", err)
	}

	if retrieved.Name != vendor.Name {
		t.Errorf("Expected name %s, got %s", vendor.Name, retrieved.Name)
	}
}

func TestCreateProductAndInventory(t *testing.T) {
	repo := NewInMemoryRepository()

	// Create vendor first
	vendor := &models.Vendor{
		ID:    "v1",
		Name:  "Garden Supplies Co",
		Email: "info@gardensupplies.com",
		Phone: "555-0200",
	}
	if err := repo.CreateVendor(vendor); err != nil {
		t.Fatalf("Failed to create vendor: %v", err)
	}

	// Create product
	product := &models.Product{
		ID:          "p1",
		Name:        "Fertilizer",
		Description: "Organic fertilizer",
		Category:    "Soil Amendments",
		Price:       29.99,
		VendorID:    "v1",
	}
	if err := repo.CreateProduct(product); err != nil {
		t.Fatalf("Failed to create product: %v", err)
	}

	// Create inventory item
	item := &models.InventoryItem{
		ID:        "i1",
		ProductID: "p1",
		Quantity:  100,
		Location:  "Warehouse A",
	}
	if err := repo.CreateInventoryItem(item); err != nil {
		t.Fatalf("Failed to create inventory item: %v", err)
	}

	// Retrieve and verify
	retrieved, err := repo.GetInventoryItem("i1")
	if err != nil {
		t.Fatalf("Failed to get inventory item: %v", err)
	}

	if retrieved.Quantity != 100 {
		t.Errorf("Expected quantity 100, got %d", retrieved.Quantity)
	}
}

func TestUpdateInventoryQuantity(t *testing.T) {
	repo := NewInMemoryRepository()

	// Create vendor first
	vendor := &models.Vendor{
		ID:    "v1",
		Name:  "Garden Supplies Co",
		Email: "info@gardensupplies.com",
	}
	if err := repo.CreateVendor(vendor); err != nil {
		t.Fatalf("Failed to create vendor: %v", err)
	}

	// Create product
	product := &models.Product{
		ID:       "p1",
		Name:     "Fertilizer",
		VendorID: "v1",
	}
	if err := repo.CreateProduct(product); err != nil {
		t.Fatalf("Failed to create product: %v", err)
	}

	// Create inventory item
	item := &models.InventoryItem{
		ID:        "i1",
		ProductID: "p1",
		Quantity:  100,
		Location:  "Warehouse A",
	}
	if err := repo.CreateInventoryItem(item); err != nil {
		t.Fatalf("Failed to create inventory item: %v", err)
	}

	// Update quantity
	err := repo.UpdateInventoryQuantity("i1", 50)
	if err != nil {
		t.Fatalf("Failed to update quantity: %v", err)
	}

	// Verify update
	updated, err := repo.GetInventoryItem("i1")
	if err != nil {
		t.Fatalf("Failed to get inventory item: %v", err)
	}

	if updated.Quantity != 50 {
		t.Errorf("Expected quantity 50, got %d", updated.Quantity)
	}
}

func TestGetNonExistentEntity(t *testing.T) {
	repo := NewInMemoryRepository()

	_, err := repo.GetSeller("nonexistent")
	if err != ErrNotFound {
		t.Errorf("Expected ErrNotFound, got %v", err)
	}

	_, err = repo.GetBuyer("nonexistent")
	if err != ErrNotFound {
		t.Errorf("Expected ErrNotFound, got %v", err)
	}

	_, err = repo.GetVendor("nonexistent")
	if err != ErrNotFound {
		t.Errorf("Expected ErrNotFound, got %v", err)
	}

	_, err = repo.GetProduct("nonexistent")
	if err != ErrNotFound {
		t.Errorf("Expected ErrNotFound, got %v", err)
	}

	_, err = repo.GetInventoryItem("nonexistent")
	if err != ErrNotFound {
		t.Errorf("Expected ErrNotFound, got %v", err)
	}
}
