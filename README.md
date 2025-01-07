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
- **Scheduling:** Cron package for scheduled tasks
- **Monitoring:** Prometheus

## Project Structure
```bash
orderly/
├── cmd/ # Application entry points
│   └── api/ # Main application
├── internal/ # Internal application code
│   ├── api/ # API layer
│   │   ├── controls/ # API controls: Price management
│   │   ├── handler/ # HTTP request handlers
│   │   ├── middleware/ # HTTP middlewares
│   │   └── routes/ # Route definitions
│   ├── domain/ # Business entities
│   ├── infrastructure/ # External implementations
│   │   ├── config/ # Configuration
│   │   ├── db/ # Database setup
│   │   └── di/ # Dependency injection
│   ├── repository/ # Data access layer
│   │   ├── interface/ # Repository interfaces
│   └── usecase/ # Business logic
│       ├── interface/ # Use case interfaces
└── pkg/ # Public libraries
    ├── jwt-token/ # JWT token handling
    ├── twilio/ # Twilio integration
    ├── utils/ # Utility functions
    └── validation/ # Validation utilities
```

## Prerequisites

- Go 1.23 or higher
- PostgreSQL 14 or higher
- Make (optional, for using Makefile commands)

## Getting Started

1. Clone the repository
```bash
git clone https://github.com/AbdulRahimOM/Orderly.git
cd Orderly
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

Enter postgres CLI:
```bash
psql -h localhost -p 5432 -U postgres -d orderly
```
In postgres, enter this to create database
```sql
CREATE DATABASE orderly;
```

5. Run migrations
```bash
go run cmd/migrate/main.go
```

6. Start the server
```bash
go run cmd/main.go
```

## Environment Variables

Key environment variables needed:
```env
PORT="6000"
JWT_SECRET_KEY="your_secret_key"
DB_HOST="localhost"
DB_USER="postgres"
DB_PASSWORD="your_db_password"
DB_NAME="orderly"
INITIAL_SUPER_ADMIN_USERNAME="initialSuperAdminUsername"
INITIAL_SUPER_ADMIN_PASSWORD="initialSuperAdminPassword"
```

Export initial super admin credentials if running in locally
```bash
export INITIAL_SUPER_ADMIN_USERNAME=superAdmin
export INITIAL_SUPER_ADMIN_PASSWORD=initialSuperAdminPassword
```

## API Documentation

### Browse Endpoints (public)
- `GET /browse/category` - List categories
- `GET /browse/category/:id` - Get category details
- `GET /browse/product` - List products
- `GET /browse/product/:id` - Get product details

#### Authentication Endpoints
- `POST /login/superAdmin` - Super admin login
- `POST /login/admin` - Admin login
- `POST /login/user` - User login
- `POST /user-signup-get-otp` - Get OTP for user signup
- `POST /user-signup-verify-otp` - Verify OTP and complete signup

### Super Admin Endpoints
#### Admin Management
- `POST /superAdmin/admin` - Create admin
- `GET /superAdmin/admin` - List admins
- `GET /superAdmin/admin/:id` - Get admin details
- `PUT /superAdmin/admin/:id` - Update admin
- `DELETE /superAdmin/admin/:id` - Soft delete admin
- `PATCH /superAdmin/admin/undo-delete/:id` - Undo soft delete admin
- `PATCH /superAdmin/admin/activate/:id` - Activate admin
- `PATCH /superAdmin/admin/deactivate/:id` - Deactivate admin

#### Access-role management
- `POST /superAdmin/access-privileges` - Create access privilege
- `GET /superAdmin/access-privileges` - List access privileges
- `GET /superAdmin/access-privileges/:admin_id` - Get access privilege by admin ID
- `DELETE /superAdmin/access-privileges/:admin_id/:privilege` - Delete access privilege


### Admin Endpoints
#### User account management: *(Available only for access-role `user_manager`)*
- `GET /admin/users` - List users
- `GET /admin/users/:id` - Get user details
- `PATCH /admin/users/activate/:id` - Activate user
- `PATCH /admin/users/deactivate/:id` - Deactivate user
- `DELETE /admin/users/:id` - Soft delete user
- `PATCH /admin/users/undo-delete/:id` - Undo soft delete user

#### Inventory Management *(Available only for access-role `inventory_manager`)*
#### ➤ Category Endpoints
- ➕ `POST /admin/category` - Create category
- 📄 `GET /admin/category` - List categories
- 🔍 `GET /admin/category/:id` - Get category details
- ✏️ `PUT /admin/category/:id` - Update category
- 🗑️ `DELETE /admin/category/:id` - Soft delete category
- 🔄 `PATCH /admin/category/undo-delete/:id` - Undo soft delete category

#### -> Product Endpoints
- `POST /admin/product` - Create product
- `GET /admin/product` - List products
- `GET /admin/product/:id` - Get product details
- `PUT /admin/product/:id` - Update product
- `DELETE /admin/product/:id` - Soft delete product
- `PATCH /admin/product/undo-delete/:id` - Undo soft delete product
- `GET /admin/product/stock/:id` - Get product stock details
- `PUT /admin/product/stock/add/:id` - Add product stock

#### Order Endpoints *(Available only for access-role `sales_manager`)*
- `GET /admin/order` - List all orders
- `GET /admin/order/:id` - Get order details
- `PATCH /admin/order/cancel/:id` - Cancel order
- `PATCH /admin/order/mark-as-delivered/:id` - Mark order as delivered

### User Endpoints 
#### Account Management
- `GET /user/account/profile` - Get user profile
- `GET /user/account/address` - List user addresses
- `GET /user/account/address/:id` - Get user address by ID
- `POST /user/account/address` - Create user address
- `PUT /user/account/address/:id` - Update user address by ID
- `DELETE /user/account/address/:id` - Delete user address by ID

#### Cart Management
- `GET /user/cart` - Get user cart
- `PUT /user/cart` - Add to cart
- `DELETE /user/cart/product/:id` - Remove product from cart
- `PATCH /user/cart/update-quantity` - Update cart item quantity
- `DELETE /user/cart/clear` - Clear cart

#### Order Management
- `GET /user/order` - List user orders
- `GET /user/order/:id` - Get user order details
- `POST /user/order` - Create order
- `PATCH /user/order/cancel/:id` - Cancel order


## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request


