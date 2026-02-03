package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/raybman/gomaterials-slt-sandbox/internal/models"
	"github.com/raybman/gomaterials-slt-sandbox/internal/repository"
	"github.com/raybman/gomaterials-slt-sandbox/internal/service"
)

// Handler provides HTTP handlers for the inventory API
type Handler struct {
	service *service.InventoryService
}

// NewHandler creates a new handler
func NewHandler(service *service.InventoryService) *Handler {
	return &Handler{service: service}
}

// Utility functions

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}

// Seller handlers

func (h *Handler) CreateSeller(w http.ResponseWriter, r *http.Request) {
	var seller models.Seller
	if err := json.NewDecoder(r.Body).Decode(&seller); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.service.CreateSeller(&seller); err != nil {
		if err == repository.ErrAlreadyExists {
			respondError(w, http.StatusConflict, "Seller already exists")
		} else {
			respondError(w, http.StatusInternalServerError, "Failed to create seller")
		}
		return
	}

	respondJSON(w, http.StatusCreated, seller)
}

func (h *Handler) ListSellers(w http.ResponseWriter, r *http.Request) {
	sellers, err := h.service.ListSellers()
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to list sellers")
		return
	}
	respondJSON(w, http.StatusOK, sellers)
}

// Buyer handlers

func (h *Handler) CreateBuyer(w http.ResponseWriter, r *http.Request) {
	var buyer models.Buyer
	if err := json.NewDecoder(r.Body).Decode(&buyer); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.service.CreateBuyer(&buyer); err != nil {
		if err == repository.ErrAlreadyExists {
			respondError(w, http.StatusConflict, "Buyer already exists")
		} else {
			respondError(w, http.StatusInternalServerError, "Failed to create buyer")
		}
		return
	}

	respondJSON(w, http.StatusCreated, buyer)
}

func (h *Handler) ListBuyers(w http.ResponseWriter, r *http.Request) {
	buyers, err := h.service.ListBuyers()
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to list buyers")
		return
	}
	respondJSON(w, http.StatusOK, buyers)
}

// Vendor handlers

func (h *Handler) CreateVendor(w http.ResponseWriter, r *http.Request) {
	var vendor models.Vendor
	if err := json.NewDecoder(r.Body).Decode(&vendor); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.service.CreateVendor(&vendor); err != nil {
		if err == repository.ErrAlreadyExists {
			respondError(w, http.StatusConflict, "Vendor already exists")
		} else {
			respondError(w, http.StatusInternalServerError, "Failed to create vendor")
		}
		return
	}

	respondJSON(w, http.StatusCreated, vendor)
}

func (h *Handler) ListVendors(w http.ResponseWriter, r *http.Request) {
	vendors, err := h.service.ListVendors()
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to list vendors")
		return
	}
	respondJSON(w, http.StatusOK, vendors)
}

// Product handlers

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.service.CreateProduct(&product); err != nil {
		if err == repository.ErrAlreadyExists {
			respondError(w, http.StatusConflict, "Product already exists")
		} else if err == repository.ErrNotFound {
			respondError(w, http.StatusBadRequest, "Vendor not found")
		} else {
			respondError(w, http.StatusInternalServerError, "Failed to create product")
		}
		return
	}

	respondJSON(w, http.StatusCreated, product)
}

func (h *Handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.ListProducts()
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to list products")
		return
	}
	respondJSON(w, http.StatusOK, products)
}

// Inventory handlers

func (h *Handler) CreateInventoryItem(w http.ResponseWriter, r *http.Request) {
	var item models.InventoryItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.service.CreateInventoryItem(&item); err != nil {
		if err == repository.ErrAlreadyExists {
			respondError(w, http.StatusConflict, "Inventory item already exists")
		} else if err == repository.ErrNotFound {
			respondError(w, http.StatusBadRequest, "Product not found")
		} else {
			respondError(w, http.StatusInternalServerError, "Failed to create inventory item")
		}
		return
	}

	respondJSON(w, http.StatusCreated, item)
}

func (h *Handler) ListInventoryItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.ListInventoryItems()
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to list inventory items")
		return
	}
	respondJSON(w, http.StatusOK, items)
}

func (h *Handler) UpdateInventoryQuantity(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID       string `json:"id"`
		Quantity int    `json:"quantity"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.service.UpdateInventoryQuantity(req.ID, req.Quantity); err != nil {
		if err == repository.ErrNotFound {
			respondError(w, http.StatusNotFound, "Inventory item not found")
		} else {
			respondError(w, http.StatusInternalServerError, "Failed to update inventory quantity")
		}
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Quantity updated successfully"})
}
