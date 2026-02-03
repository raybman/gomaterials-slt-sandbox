# Go Materials Inventory Management System

A basic inventory management application for Go Materials, a garden supply company. This system provides functionality for managing sellers, buyers, vendors, products, and inventory items.

## Features

- Manage sellers, buyers, and vendors
- Track products from various vendors
- Maintain inventory with quantity tracking
- RESTful API for all operations
- In-memory data storage

## Project Structure

```
.
├── cmd/
│   └── server/          # Main application entry point
├── internal/
│   ├── models/          # Data models
│   ├── repository/      # Data storage layer
│   ├── service/         # Business logic layer
│   └── handlers/        # HTTP handlers
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.24 or higher

### Installation

1. Clone the repository:
```bash
git clone https://github.com/raybman/gomaterials-slt-sandbox.git
cd gomaterials-slt-sandbox
```

2. Install dependencies:
```bash
go mod download
```

### Running the Application

Start the server:
```bash
go run cmd/server/main.go
```

The server will start on port 8080. Visit http://localhost:8080 for API documentation.

### Running Tests

Run all tests:
```bash
go test ./...
```

Run tests with verbose output:
```bash
go test -v ./...
```

## API Endpoints

### Sellers
- `POST /api/sellers` - Create a new seller
- `GET /api/sellers` - List all sellers

### Buyers
- `POST /api/buyers` - Create a new buyer
- `GET /api/buyers` - List all buyers

### Vendors
- `POST /api/vendors` - Create a new vendor
- `GET /api/vendors` - List all vendors

### Products
- `POST /api/products` - Create a new product
- `GET /api/products` - List all products

### Inventory
- `POST /api/inventory` - Create a new inventory item
- `GET /api/inventory` - List all inventory items
- `POST /api/inventory/update` - Update inventory quantity

### Health Check
- `GET /health` - Check server health

## Example Usage

### Create a Vendor
```bash
curl -X POST http://localhost:8080/api/vendors \
  -H "Content-Type: application/json" \
  -d '{
    "id": "v1",
    "name": "Garden Supplies Co",
    "email": "info@gardensupplies.com",
    "phone": "555-0200",
    "address": "123 Garden St"
  }'
```

### Create a Product
```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{
    "id": "p1",
    "name": "Organic Fertilizer",
    "description": "High-quality organic fertilizer",
    "category": "Soil Amendments",
    "price": 29.99,
    "vendor_id": "v1"
  }'
```

### Create an Inventory Item
```bash
curl -X POST http://localhost:8080/api/inventory \
  -H "Content-Type: application/json" \
  -d '{
    "id": "i1",
    "product_id": "p1",
    "quantity": 100,
    "location": "Warehouse A"
  }'
```

### Update Inventory Quantity
```bash
curl -X POST http://localhost:8080/api/inventory/update \
  -H "Content-Type: application/json" \
  -d '{
    "id": "i1",
    "quantity": 75
  }'
```

### List Products
```bash
curl http://localhost:8080/api/products
```

## License

MIT