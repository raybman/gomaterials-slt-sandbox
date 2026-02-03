package service

import (
	"github.com/raybman/gomaterials-slt-sandbox/internal/models"
	"github.com/raybman/gomaterials-slt-sandbox/internal/repository"
)

// InventoryService provides business logic for inventory management
type InventoryService struct {
	repo *repository.InMemoryRepository
}

// NewInventoryService creates a new inventory service
func NewInventoryService(repo *repository.InMemoryRepository) *InventoryService {
	return &InventoryService{repo: repo}
}

// Seller operations

func (s *InventoryService) CreateSeller(seller *models.Seller) error {
	return s.repo.CreateSeller(seller)
}

func (s *InventoryService) GetSeller(id string) (*models.Seller, error) {
	return s.repo.GetSeller(id)
}

func (s *InventoryService) ListSellers() ([]*models.Seller, error) {
	return s.repo.ListSellers()
}

// Buyer operations

func (s *InventoryService) CreateBuyer(buyer *models.Buyer) error {
	return s.repo.CreateBuyer(buyer)
}

func (s *InventoryService) GetBuyer(id string) (*models.Buyer, error) {
	return s.repo.GetBuyer(id)
}

func (s *InventoryService) ListBuyers() ([]*models.Buyer, error) {
	return s.repo.ListBuyers()
}

// Vendor operations

func (s *InventoryService) CreateVendor(vendor *models.Vendor) error {
	return s.repo.CreateVendor(vendor)
}

func (s *InventoryService) GetVendor(id string) (*models.Vendor, error) {
	return s.repo.GetVendor(id)
}

func (s *InventoryService) ListVendors() ([]*models.Vendor, error) {
	return s.repo.ListVendors()
}

// Product operations

func (s *InventoryService) CreateProduct(product *models.Product) error {
	// Verify vendor exists before creating product
	_, err := s.repo.GetVendor(product.VendorID)
	if err != nil {
		return err
	}
	return s.repo.CreateProduct(product)
}

func (s *InventoryService) GetProduct(id string) (*models.Product, error) {
	return s.repo.GetProduct(id)
}

func (s *InventoryService) ListProducts() ([]*models.Product, error) {
	return s.repo.ListProducts()
}

// Inventory operations

func (s *InventoryService) CreateInventoryItem(item *models.InventoryItem) error {
	// Verify product exists before creating inventory item
	_, err := s.repo.GetProduct(item.ProductID)
	if err != nil {
		return err
	}
	return s.repo.CreateInventoryItem(item)
}

func (s *InventoryService) GetInventoryItem(id string) (*models.InventoryItem, error) {
	return s.repo.GetInventoryItem(id)
}

func (s *InventoryService) UpdateInventoryQuantity(id string, quantity int) error {
	return s.repo.UpdateInventoryQuantity(id, quantity)
}

func (s *InventoryService) ListInventoryItems() ([]*models.InventoryItem, error) {
	return s.repo.ListInventoryItems()
}
