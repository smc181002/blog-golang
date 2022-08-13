package main

import (
	"fmt"
	// "log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
	"github.com/smc181002/blog-golang/config"
	"github.com/smc181002/blog-golang/models"
	// "gorm.io/gorm"
)

func main() {
    app := chi.NewMux()
    const port int = 8080

    go http.ListenAndServe(fmt.Sprintf(":%v", port), app)
    fmt.Printf("Server started on http://localhost:%v\n", port)
    fmt.Println("Press Ctrl+C to stop the server")

    models.DBInitialize(config.ServerConfig)
    fmt.Println("Database connected successfully")

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt)
    <- stop
}
