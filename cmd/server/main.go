package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/brunoofgod/goexpert-lesson-4/docs"
	"github.com/brunoofgod/goexpert-lesson-4/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Clima API
// @version 1.0
// @description API que recebe um CEP e retorna a temperatura atual.
// @host localhost:8080
// @BasePath /
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Rotas Swagger
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
	})

	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))

	// Rotas aplicacao
	r.Post("/weather", handlers.GetWeather)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor rodando na porta %s...", port)
	http.ListenAndServe(":"+port, r)
}
