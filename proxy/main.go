package main

import (
	"context"
	"fmt"
	"log"
	"metricsProm/proxy/internal/repository"
	"net/http"
	"os"
	"os/signal"

	"time"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

// swagger:route POST /api/address/search  addr RequestAddressSearch
// getting address
// responses:
// 200:

var repo repository.GeoRepo

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env load error" + err.Error())
	}
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	))

	r := repository.NewPostgreGeoRepo(db)
	repo = &r

	router := chi.NewRouter()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	proxy := NewReverseProxy("hugo", "1313")

	//go TimeUpdate()
	//go BinTreeBuilt()
	//go graphRandomBuilt()

	os.Setenv("HOST", proxy.host)

	router.Use(proxy.ReverseProxy)

	router.Mount("/", newApiRouter())

	go func() {
		log.Println("Starting server...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")

	//http.ListenAndServe(":8080", router)

}
