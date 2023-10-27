package main

import (
	"context"
	"fmt"
	"generate-auth-workshop/controller/jwtController"
	"generate-auth-workshop/controller/sessionController"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pterm/pterm"
)

func main() {
	// Init Variables:
	port := 3000

	// Load environment variables:
	pterm.Println("Loading environment variables")
	if err := godotenv.Load(".env"); err != nil {
		pterm.Fatal.WithFatal(false).Println("Failed to load environment variables!")
		os.Exit(1)
	}

	// Initialize an echo instance:
	e := echo.New()
	e.HideBanner = true
	logger := e.Logger
	logger.SetLevel(log.INFO)
	logger.SetHeader("")

	pterm.Printfln("Spinning up server on port %d", port)
	// Go routine to spin up the backend on port :3000
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", port)); err != nil && err != http.ErrServerClosed {
			pterm.Fatal.WithFatal(false).Println(err)
			os.Exit(1)
		}
	}()

	pterm.Println("Registering Backend Routes")
	// Register the backend routes:
	api := e.Group("/api")
	sessioncontroller.RegisterSessionAuthRoutes(api, logger)
	jwtcontroller.RegisterJWTAuthRoutes(api, logger)

	// Graceful shutdown:
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	pterm.Println("\nShutting down server...")
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
