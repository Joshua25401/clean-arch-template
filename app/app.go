package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	// Depends to HTTP we're using gin
	router http.Handler
	// Holy closer function
	close func() error
}

func New() *App {
	router := gin.Default()

	// Setup middleware
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	router.Use(cors.New(corsConfig))
	router.Use()

	// Setup gin mode
	gin.SetMode(gin.DebugMode) // Change to relase

	// Enable golang to log filename
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Setup router / handler below
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]any{
			"message": "Welcome to API",
		})
	})
	// Setup any dependency for other layer below!

	// Setup anonymous function to clean up dependecy
	cleanup := func() error {
		// Close ur http dependency here
		return nil
	}
	// Return the app object
	return &App{
		router: router,
		close:  cleanup,
	}
}

func (app App) Run() {
	// Setup your handler server here
	httpServer := &http.Server{
		Addr:                         ":8080",
		Handler:                      app.router,
		DisableGeneralOptionsHandler: false,
	}

	// Run your http server in separate goroutine
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server Error: %v\n", err)
		}
	}()

	// Make a quit channel
	quit := make(chan os.Signal, 1)
	/**
	List of signal :
		1. Hang Up Signal
		2. Interrupt Signal
		3. Terminate Signal
		4. Quit Signal
	*/
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Wait the signal come to our quit channel
	<-quit

	fmt.Println("Try to gracefully shutdown the server.....")

	// Create shutdown signal with n sec grace period
	shutdownContext, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer func() {
		cancelFunc()

		if err := app.close(); err != nil {
			log.Printf("resource close with err %v", err)
		}
	}()

	// Shutdown the server
	if err := httpServer.Shutdown(shutdownContext); err != nil {
		log.Printf("Shutdown server with error %v", err)
		return
	}

	log.Println("Gracefully shutdown the server.....")
}
