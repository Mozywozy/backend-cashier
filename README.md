# Cashier App Backend

<div style="text-align: center;">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" width="100">
  <img src="https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL" width="200">
</div>

This is the backend service for the Cashier App, built using Go, Gin, and Gorm. The backend provides authentication, role-based access control (Admin, Warehouse, Cashier, Customer), and allows for product management and sales transactions. The database used is PostgreSQL.

## Features

- **Authentication** (Login, Register, Logout)
- **Role-based access control**:
  - **Admin**: Manage users and products
  - **Warehouse Staff**: Manage products
  - **Cashier**: Manage sales transactions
  - **Customer**: View transaction history
- **Product management** (CRUD operations)
- **Sales transaction management**
- **JSON Web Token (JWT) authentication**

## Getting Started

### Prerequisites

Make sure you have the following installed:

- [Go](https://golang.org/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/)

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/cashier-app.git
    cd cashier-app
    ```

2. Create a `.env` file in the project root and configure your environment variables:

    ```bash
    DB_USER=postgres
    DB_PASSWORD=yourpassword
    DB_NAME=cashier_app
    DB_PORT=5432
    JWT_SECRET=your_jwt_secret
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

4. Run PostgreSQL and create the necessary database:

    ```bash
    psql -U postgres -c "CREATE DATABASE cashier_app;"
    ```

5. Start the server:

    ```bash
    go run main.go
    ```

    The server will start on `http://localhost:8080`.

## API Endpoints

### Authentication

- **POST** `/login` - Login a user and get a JWT token
- **POST** `/register` - Register a new user
- **POST** `/logout` - Logout the current user

### Admin

- **GET** `/admin/products` - Get all products
- **GET** `/admin/products/:id` - Get a single product by ID
- **POST** `/admin/products` - Create a new product
- **PUT** `/admin/products/:id` - Update a product
- **DELETE** `/admin/products/:id` - Delete a product
- **GET** `/admin/users` - Get all users
- **GET** `/admin/users/:id` - Get a user by ID
- **PUT** `/admin/users/:id` - Update user details
- **DELETE** `/admin/users/:id` - Delete a user

### Warehouse Staff

- **GET** `/petugas_gudang/products` - Get all products
- **GET** `/petugas_gudang/products/:id` - Get a single product by ID
- **PUT** `/petugas_gudang/products/:id` - Update a product
- **DELETE** `/petugas_gudang/products/:id` - Delete a product

### Cashier

- **GET** `/petugas_kasir/sales` - Get all sales
- **GET** `/petugas_kasir/sales/:id` - Get a sale by ID
- **POST** `/petugas_kasir/sales` - Create a new sale
- **PUT** `/petugas_kasir/sales/:id` - Update a sale
- **DELETE** `/petugas_kasir/sales/:id` - Delete a sale

### Customer

- **GET** `/pelanggan/transactions` - Get all transactions for the logged-in customer

## Environment Variables

You will need to configure the following environment variables in your `.env` file:

- `DB_USER`: Database username
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name
- `DB_PORT`: Database port (default: 5432)
- `JWT_SECRET`: Secret key for signing JWT tokens
