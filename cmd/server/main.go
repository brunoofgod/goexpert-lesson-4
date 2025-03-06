package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/brunoofgod/goexpert-lesson-4/docs"
	"github.com/brunoofgod/goexpert-lesson-4/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Clima API
// @version 1.0
// @description API que recebe um CEP e retorna a temperatura atual.
// @BasePath /
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Middleware CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Rotas do Swagger
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
	})

	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("https://"+os.Getenv("HOSTNAME")+"/swagger/doc.json")))

	// Rotas da aplicacao
	r.Post("/weather", handlers.GetWeather)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor rodando na porta %s...", port)
	http.ListenAndServe(":"+port, r)
}
