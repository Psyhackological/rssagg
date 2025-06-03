package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Psyhackological/rssagg/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT is not found in the environment")
	}
	fmt.Println("Port:", PORT)

	DB_URL := os.Getenv("DB_URL")
	if DB_URL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", DB_URL)
	if err != nil {
		log.Fatal("Can't connect to database")
	}

	queries := database.New(conn)

	apiCfg := apiConfig{
		DB: queries,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	router.Mount("/v1", v1Router)

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	log.Printf("Server starting on port %v", PORT)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
