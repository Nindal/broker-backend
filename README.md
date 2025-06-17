
# 🏦 Multi-Stock Broker Platform Backend

This is the backend API service for a multi-stock broker platform built with Go. It supports user authentication (with JWT), mock holdings, orderbook, and position data, and follows a modular, scalable architecture.

## 🔧 Features

- JWT-based Authentication with Refresh Tokens
- PostgreSQL for user data
- REST APIs for Holdings, Orderbook, and Positions
- Circuit Breaker Pattern (using GoBreaker)
- Dockerized setup with `docker-compose`
- Secure and modular Go project structure

## 📦 API Endpoints

| Method | Endpoint       | Description                            | Auth Required |
|--------|----------------|----------------------------------------|---------------|
| GET    | `/health`      | Health check                           | ❌            |
| POST   | `/signup`      | Register new user                      | ❌            |
| POST   | `/login`       | Authenticate user and return tokens    | ❌            |
| POST   | `/refresh`     | Refresh access token using refresh     | ❌            |
| GET    | `/holdings`    | Get mock holdings                      | ✅            |
| GET    | `/orderbook`   | Get mock orderbook and PnL data        | ✅            |
| GET    | `/positions`   | Get mock positions and PnL summary     | ✅            |

## 🗂 Project Structure

broker-backend/
├── cmd/server/ # Entry point
├── db/ # Database init and migrations
├── handlers/ # HTTP handlers
├── middleware/ # JWT Middleware
├── models/ # Struct models
├── utils/ # JWT, mock data, circuit breaker
├── Dockerfile
├── docker-compose.yml
├── go.mod / go.sum
└── .env.example


## 🔐 Authentication

- Access tokens are valid for 10 minutes.
- Refresh tokens are valid for 24 hours.
- Use `Authorization: Bearer <access_token>` in headers for authenticated routes.

## 🚀 Getting Started

### 1. Clone and Set Up
git clone https://github.com/yourusername/broker-backend.git
cd broker-backend
cp .env.example .env

Update .env with your secrets if needed.

2. Run with Docker
bash
Copy
Edit
docker-compose up --build
The backend will run on http://localhost:8080

PostgreSQL is exposed on port 5432.

3. Migrate the DB
Once PostgreSQL is up:
docker exec -i <db_container_id> psql -U broker -d broker < db/migrations/init.sql
Or automate it with a migration tool of your choice.
