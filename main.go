package main

import (
	"find-pairs/config"
	"find-pairs/controllers"
	"find-pairs/middleware"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/find-pairs", middleware.JwtAuthMiddleware(controllers.FindPairs))
	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
