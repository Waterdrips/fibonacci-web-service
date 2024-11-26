package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func fibonacciHandler(c *gin.Context) {
	var input struct {
		Number *int `json:"number,omitempty"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if input.Number == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Number field is required"})
		return
	}

	if *input.Number < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Number must be non-negative"})
		return
	}
	result := fibonacciSequence(*input.Number)

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/", fibonacciHandler)

	r.GET("/healthz", func(c *gin.Context) {
		// Our service has 0 dependencies, so it's always healthy as long as it loads!
		// This can be expanded upon should we add dependencies in the future.
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a context with timeout to allow for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
