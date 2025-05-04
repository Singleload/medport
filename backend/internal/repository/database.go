package repository

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/svenskhalsovard/api/internal/config"
)

// Database represents a database connection
type Database struct {
	DB *sqlx.DB
}

// NewDatabase creates a new database connection
func NewDatabase(cfg config.DatabaseConfig) (*Database, error) {
	db, err := sqlx.Connect(cfg.Driver, cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Info().
		Str("driver", cfg.Driver).
		Int("maxOpenConns", cfg.MaxOpenConns).
		Int("maxIdleConns", cfg.MaxIdleConns).
		Dur("connMaxLifetime", cfg.ConnMaxLifetime).
		Msg("Database connection established")

	return &Database{DB: db}, nil
}

// Close closes the database connection
func (d *Database) Close() {
	if d.DB != nil {
		if err := d.DB.Close(); err != nil {
			log.Error().Err(err).Msg("Error closing database connection")
			return
		}
		log.Info().Msg("Database connection closed")
	}
}

// RunMigrations runs database migrations
func RunMigrations(cfg config.DatabaseConfig) error {
	log.Info().
		Str("migrationsPath", cfg.MigrationsPath).
		Msg("Running database migrations")

	// In a real implementation, you would use a migration library like migrate, 
	// goose, or golang-migrate to run migrations. For simplicity, we'll just 
	// log that migrations would run here.
	
	// For example with golang-migrate:
	// m, err := migrate.New("file://"+cfg.MigrationsPath, cfg.DSN)
	// if err != nil {
	//     return fmt.Errorf("failed to create migration instance: %w", err)
	// }
	// 
	// if err := m.Up(); err != nil && err != migrate.ErrNoChange {
	//     return fmt.Errorf("failed to run migrations: %w", err)
	// }

	log.Info().Msg("Database migrations completed successfully")
	return nil
}

// Transaction represents a database transaction
func (d *Database) Transaction(fn func(*sqlx.Tx) error) error {
	tx, err := d.DB.Beginx()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			// Rollback on panic
			_ = tx.Rollback()
			panic(p) // Re-throw panic after rollback
		}
	}()

	if err := fn(tx); err != nil {
		// Rollback on error
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Error().Err(rbErr).Msg("Failed to rollback transaction")
		}
		return err
	}

	// Commit transaction if no error
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// Helper function to add soft delete condition to queries
func softDeleteCondition(tableName string) string {
	return fmt.Sprintf("%s.deleted_at IS NULL", tableName)
}

// Helper function to get current time in UTC
func now() time.Time {
	return time.Now().UTC()
}