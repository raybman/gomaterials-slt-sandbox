package models

import "time"

// Seller represents a seller entity in the system
type Seller struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}

// Buyer represents a buyer entity in the system
type Buyer struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

// Vendor represents a vendor entity in the system
type Vendor struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

// Product represents a product in the inventory
type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Price       float64   `json:"price"`
	VendorID    string    `json:"vendor_id"`
	CreatedAt   time.Time `json:"created_at"`
}

// InventoryItem represents an inventory item with quantity tracking
type InventoryItem struct {
	ID        string    `json:"id"`
	ProductID string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Location  string    `json:"location"`
	UpdatedAt time.Time `json:"updated_at"`
}
