package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/raybman/gomaterials-slt-sandbox/internal/models"
)

var (
	ErrNotFound      = errors.New("entity not found")
	ErrAlreadyExists = errors.New("entity already exists")
)

// InMemoryRepository provides in-memory storage for all entities
type InMemoryRepository struct {
	sellers   map[string]*models.Seller
	buyers    map[string]*models.Buyer
	vendors   map[string]*models.Vendor
	products  map[string]*models.Product
	inventory map[string]*models.InventoryItem
	mu        sync.RWMutex
}

// NewInMemoryRepository creates a new in-memory repository
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		sellers:   make(map[string]*models.Seller),
		buyers:    make(map[string]*models.Buyer),
		vendors:   make(map[string]*models.Vendor),
		products:  make(map[string]*models.Product),
		inventory: make(map[string]*models.InventoryItem),
	}
}

// Seller methods

func (r *InMemoryRepository) CreateSeller(seller *models.Seller) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.sellers[seller.ID]; exists {
		return ErrAlreadyExists
	}
	seller.CreatedAt = time.Now()
	r.sellers[seller.ID] = seller
	return nil
}

func (r *InMemoryRepository) GetSeller(id string) (*models.Seller, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	seller, exists := r.sellers[id]
	if !exists {
		return nil, ErrNotFound
	}
	return seller, nil
}

func (r *InMemoryRepository) ListSellers() ([]*models.Seller, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	sellers := make([]*models.Seller, 0, len(r.sellers))
	for _, seller := range r.sellers {
		sellers = append(sellers, seller)
	}
	return sellers, nil
}

// Buyer methods

func (r *InMemoryRepository) CreateBuyer(buyer *models.Buyer) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.buyers[buyer.ID]; exists {
		return ErrAlreadyExists
	}
	buyer.CreatedAt = time.Now()
	r.buyers[buyer.ID] = buyer
	return nil
}

func (r *InMemoryRepository) GetBuyer(id string) (*models.Buyer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	buyer, exists := r.buyers[id]
	if !exists {
		return nil, ErrNotFound
	}
	return buyer, nil
}

func (r *InMemoryRepository) ListBuyers() ([]*models.Buyer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	buyers := make([]*models.Buyer, 0, len(r.buyers))
	for _, buyer := range r.buyers {
		buyers = append(buyers, buyer)
	}
	return buyers, nil
}

// Vendor methods

func (r *InMemoryRepository) CreateVendor(vendor *models.Vendor) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.vendors[vendor.ID]; exists {
		return ErrAlreadyExists
	}
	vendor.CreatedAt = time.Now()
	r.vendors[vendor.ID] = vendor
	return nil
}

func (r *InMemoryRepository) GetVendor(id string) (*models.Vendor, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	vendor, exists := r.vendors[id]
	if !exists {
		return nil, ErrNotFound
	}
	return vendor, nil
}

func (r *InMemoryRepository) ListVendors() ([]*models.Vendor, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	vendors := make([]*models.Vendor, 0, len(r.vendors))
	for _, vendor := range r.vendors {
		vendors = append(vendors, vendor)
	}
	return vendors, nil
}

// Product methods

func (r *InMemoryRepository) CreateProduct(product *models.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.products[product.ID]; exists {
		return ErrAlreadyExists
	}
	product.CreatedAt = time.Now()
	r.products[product.ID] = product
	return nil
}

func (r *InMemoryRepository) GetProduct(id string) (*models.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	product, exists := r.products[id]
	if !exists {
		return nil, ErrNotFound
	}
	return product, nil
}

func (r *InMemoryRepository) ListProducts() ([]*models.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	products := make([]*models.Product, 0, len(r.products))
	for _, product := range r.products {
		products = append(products, product)
	}
	return products, nil
}

// Inventory methods

func (r *InMemoryRepository) CreateInventoryItem(item *models.InventoryItem) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.inventory[item.ID]; exists {
		return ErrAlreadyExists
	}
	item.UpdatedAt = time.Now()
	r.inventory[item.ID] = item
	return nil
}

func (r *InMemoryRepository) GetInventoryItem(id string) (*models.InventoryItem, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	item, exists := r.inventory[id]
	if !exists {
		return nil, ErrNotFound
	}
	return item, nil
}

func (r *InMemoryRepository) UpdateInventoryQuantity(id string, quantity int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	item, exists := r.inventory[id]
	if !exists {
		return ErrNotFound
	}
	item.Quantity = quantity
	item.UpdatedAt = time.Now()
	return nil
}

func (r *InMemoryRepository) ListInventoryItems() ([]*models.InventoryItem, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	items := make([]*models.InventoryItem, 0, len(r.inventory))
	for _, item := range r.inventory {
		items = append(items, item)
	}
	return items, nil
}
