package main

import (
	"log"
	"net/http"
	"os"
	"service-nexa/pkg/database"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	log.Printf("Debug: DATABASE_URL is '%s'", os.Getenv("DATABASE_URL"))

	db, err := database.ConnectPostgres()
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Service Nexa"))
	})
	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
