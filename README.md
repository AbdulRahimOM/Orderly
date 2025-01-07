# Orderly - Order and Inventory Management System

Orderly is a robust backend system built with Go that handles order processing and inventory management with dynamic pricing capabilities. The system is designed to be scalable, maintainable, and follows clean architecture principles.

## Features

### User Management
- User authentication (signup/signin)
- Role-based access control (Super Admin, Admin, User)
- OTP verification for user signup
- User profile and address management

### Product Management
- Category management
- Product CRUD operations
- Dynamic pricing based on demand and availability
- Stock management

### Order Management
- Shopping cart functionality
- Order processing
- Order status tracking
- Order cancellation and delivery management

### Admin Features
- Admin user management
- Access privilege control
- Inventory management
- Order management with filtering and sorting
- Sales and inventory statistics

## Tech Stack

- **Language:** Go 1.21+
- **Framework:** Fiber
- **Database:** PostgreSQL
- **ORM:** GORM
- **Authentication:** JWT
- **SMS:** Twilio
- **Email:** SMTP

## Project Structure
```bash
orderly/
├── cmd/
│ └── api/ # Application entry points
├── internal/
│ ├── api/
│ │ ├── handler/ # HTTP request handlers
│ │ ├── middleware/ # HTTP middlewares
│ │ └── routes/ # Route definitions
│ ├── core/
│ │ ├── domain/ # Business entities
│ │ └── ports/ # Interfaces
│ ├── infrastructure/ # External implementations
│ │ ├── config/ # Configuration
│ │ └── db/ # Database setup
│ ├── repository/ # Data access layer
│ └── usecase/ # Business logic
└── pkg/ # Public libraries
```

## Prerequisites

- Go 1.23 or higher
- PostgreSQL 14 or higher
- Make (optional, for using Makefile commands)

## Getting Started

1. Clone the repository
```bash
git clone https://github.com/yourusername/orderly.git
cd orderly
```

2. Set up environment variables

```bash
cp .env.example .env
```

 Edit .env with your configuration

3. Install dependencies
```bash
go mod download
```

4. Set up the database
```bash
createdb orderly
```

5. Run migrations
```bash
go run cmd/migrate/main.go
```

6. Start the server
```bash
go run cmd/api/main.go
```

## Environment Variables

Key environment variables needed:
```env
PORT=6000
JWT_SECRET_KEY=your_secret_key
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=orderly
```

## API Documentation

### Authentication Endpoints
- `POST /login/superAdmin` - Super admin login
- `POST /login/admin` - Admin login
- `POST /login/user` - User login
- `POST /user-signup-get-otp` - Get OTP for user signup
- `POST /user-signup-verify-otp` - Verify OTP and complete signup

### Admin Endpoints
- `POST /admin/users` - Create user
- `GET /admin/users` - List users
- `GET /admin/users/:id` - Get user details
- `PUT /admin/users/:id` - Update user
- `DELETE /admin/users/:id` - Delete user

### Product Endpoints
- `POST /admin/product` - Create product
- `GET /admin/product` - List products
- `GET /admin/product/:id` - Get product details
- `PUT /admin/product/:id` - Update product
- `DELETE /admin/product/:id` - Delete product

### Order Endpoints
- `POST /user/order` - Create order
- `GET /user/order` - List user orders
- `GET /admin/order` - List all orders (admin)
- `PATCH /admin/order/mark-as-delivered/:id` - Mark order as delivered

## Testing

Run the test suite:

```bash
go test ./...
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request


