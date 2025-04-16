package main

import (
	"log"
	"net/http"

	"github.com/Prayag2003/bloom-filter/bloom"
	"github.com/Prayag2003/bloom-filter/handlers"
	"github.com/Prayag2003/bloom-filter/middleware"
	"github.com/Prayag2003/bloom-filter/storage"
)

func main() {
	filter := bloom.New(1_000_000, 3)
	storage := storage.NewFileStorage("data/users.txt")
	handler := handlers.NewUserHandler(filter, storage)

	handler.LoadUsernames()

	mux := http.NewServeMux()
	mux.HandleFunc("/check-username", handler.CheckUsername)
	mux.HandleFunc("/register", handler.RegisterUsername)

	log.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", middleware.EnableCORS(middleware.LogRequests(mux))))
}
