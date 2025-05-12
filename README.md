# 🏥 Pharmacy API

A RESTful API for managing pharmacy items and invoices, built with Go (Fiber), GORM, and PostgreSQL. Supports live reloading with Air and containerized using Docker.

## 📦 Tech Stack

* **Golang (Fiber)**
* **GORM**
* **PostgreSQL**
* **Docker + Docker Compose**
* **Air** (for live reload in dev)

## ⚙️ Prerequisites

- [Docker](https://www.docker.com/) installed

## 🚀 Setup Instructions

1. **Clone the repository**

   ```bash
   git clone https://github.com/DinethDilhara/pharmacy-api.git
   cd pharmacy-api
   ```

2. **Create environment file**

   ```bash
   touch .env
   ```

   Add the following variables to `.env`:

   ```env
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   ```

3. **Start the project using Docker Compose**

   ```bash
   docker-compose up --build
   ```

## 🧪 API Endpoints

Refer to the included [Postman collection](#https://github.com/DinethDilhara/pharmacy-api.git/Pharmacy-API.postman_collection.json) and [Swagger Doc](#https://github.com/DinethDilhara/pharmacy-api.git/Pharmacy-api-swagger.yaml)for full endpoint examples including:


* CRUD for Items and Invoices
* Adding, updating, deleting Invoice Items
* Viewing items in a specific Invoice

