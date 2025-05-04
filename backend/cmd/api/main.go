package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/svenskhalsovard/api/internal/config"
	"github.com/svenskhalsovard/api/internal/handlers"
	"github.com/svenskhalsovard/api/internal/middleware"
	"github.com/svenskhalsovard/api/internal/repository"
	"github.com/svenskhalsovard/api/internal/service"
	"github.com/svenskhalsovard/api/internal/svea"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Warning: .env file not found or cannot be read: %v\n", err)
	}

	// Setup logger
	setupLogger()

	// Load application configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	log.Info().Msg("Starting Svensk Hälsovård API...")
	log.Info().Str("environment", cfg.Env).Msg("Environment")

	// Initialize database connection
	db, err := repository.NewDatabase(cfg.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer db.Close()

	// Run database migrations if enabled
	if cfg.Database.RunMigrations {
		if err := repository.RunMigrations(cfg.Database); err != nil {
			log.Fatal().Err(err).Msg("Failed to run database migrations")
		}
	}

	// Initialize repositories
	serviceRepo := repository.NewServiceRepository(db)
	bookingRepo := repository.NewBookingRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)

	// Initialize Svea Ekonomi client
	sveaClient := svea.NewClient(cfg.Svea)

	// Initialize services
	serviceService := service.NewServiceService(serviceRepo)
	paymentService := service.NewPaymentService(paymentRepo, sveaClient)
	bookingService := service.NewBookingService(bookingRepo, paymentRepo)
	checkoutService := service.NewCheckoutService(paymentService, bookingService, serviceService)

	// Initialize router
	router := handlers.NewRouter(cfg)

	// Register health check and version endpoints
	router.Get("/health", handlers.HealthCheck)
	router.Get("/version", handlers.Version)

	// Register API routes
	apiRouter := router.Group("/api")
	apiRouter.Use(middleware.RequestID)
	apiRouter.Use(middleware.Logging)
	apiRouter.Use(middleware.Recovery)
	apiRouter.Use(middleware.CORS(cfg.CORS))
	apiRouter.Use(middleware.RateLimiter(cfg.RateLimit))

	// Register service handlers
	serviceHandler := handlers.NewServiceHandler(serviceService)
	apiRouter.Get("/services", serviceHandler.GetServices)
	apiRouter.Get("/services/{id}", serviceHandler.GetServiceByID)

	// Register checkout handlers
	checkoutHandler := handlers.NewCheckoutHandler(checkoutService)
	apiRouter.Post("/checkout/initiate", checkoutHandler.InitiateCheckout)
	apiRouter.Post("/checkout/payment", checkoutHandler.ProcessPayment)
	apiRouter.Get("/checkout/verify/{paymentId}", checkoutHandler.VerifyPayment)

	// Register booking handlers
	bookingHandler := handlers.NewBookingHandler(bookingService)
	apiRouter.Post("/bookings", bookingHandler.CreateBooking)

	// Start server with graceful shutdown
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeoutSeconds) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeoutSeconds) * time.Second,
		IdleTimeout:  time.Duration(cfg.Server.IdleTimeoutSeconds) * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Info().Int("port", cfg.Server.Port).Msg("HTTP server started")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	// Setup graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Shutting down server...")

	// Create context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Server.ShutdownTimeoutSeconds)*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exited properly")
}

func setupLogger() {
	// Set up structured logging with zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Enable pretty console logging in development
	if os.Getenv("APP_ENV") != "production" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}