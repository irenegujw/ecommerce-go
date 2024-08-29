# E-commerce-Go

This project is an e-commerce application built using Go (Golang). It provides functionalities for managing products, users, orders, and payments. The application is designed with a modular architecture, leveraging Go's concurrency and efficiency to deliver a robust and scalable e-commerce platform.

## Features

- **User Management**: Registration, login, authentication, and profile management.
- **Product Management**: Add, update, delete, and search products.
- **Order Management**: Create, update, and manage orders.
- **Payment Integration**: Secure payment processing with third-party payment gateways.
- **RESTful API**: Exposes a set of RESTful endpoints for front-end integration.
- **Concurrency Handling**: Efficient handling of concurrent requests and operations.
- **Modular Architecture**: Designed for easy expansion and maintainability.

## Tech Stack

- **Backend**: Go (Golang)
- **Database**: PostgreSQL (or any preferred SQL database)
- **Authentication**: JWT-based authentication
- **API**: RESTful API using Go’s standard libraries
- **Dependency Management**: Go Modules

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or above)
- [PostgreSQL](https://www.postgresql.org/download/)
- Git

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/irenegujw/ecommerce-go.git
   cd ecommerce-go
   ```

2. **Install dependencies:**
   
   ```bash
   go mod tidy
   ```
   
3. **Set up environment variables:**

   Create a `.env` file in the root directory and add the following:

  ```env
  DB_HOST=localhost
  DB_PORT=5432
  DB_USER=yourusername
  DB_PASSWORD=yourpassword
  DB_NAME=ecommerce
  JWT_SECRET=your_jwt_secret
   ```

4. **Run the application:**

  ```bash
  go run main.go
  ```

5. **Access the API:**

        The server will be running on http://localhost:8000. You can use tools like Postman or curl to interact with the API endpoints.
   
