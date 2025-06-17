package main

import (
	"log"
	"net/http"
	"broker-backend/db"
	"broker-backend/handlers"
	"broker-backend/middleware"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db.InitDB()

	r := mux.NewRouter()

	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
	r.HandleFunc("/signup", handlers.Signup).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/refresh", handlers.RefreshToken).Methods("POST")

	auth := r.PathPrefix("/").Subrouter()
	auth.Use(middleware.JWTAuthMiddleware)

	auth.HandleFunc("/holdings", handlers.GetHoldings).Methods("GET")
	auth.HandleFunc("/orderbook", handlers.GetOrderbook).Methods("GET")
	auth.HandleFunc("/positions", handlers.GetPositions).Methods("GET")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
