package psql

import (
	"context"
	"github.com/pkg/errors"
	"testing"
	"time"
)

func TestNewClientNew(t *testing.T) {
	cfg := PgConfig{
		username: "myuser",
		password: "mypassword",
		host:     "0.0.0.0",
		port:     "5432",
		database: "prod_service",
	}

	ctx := context.Background()
	maxAttempts := 3
	maxDelay := time.Second * 5
	dsn := cfg.ConnStringFromCfg()
	binary := false

	pool, err := NewClient(ctx, maxAttempts, maxDelay, dsn, binary)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if pool == nil {
		t.Error("Expected a non-nil pool, but got nil")
	}

	// Test parsing config error
	_, err = NewClient(ctx, maxAttempts, maxDelay, "invalid_dsn", binary)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

}

func TestDoWithAttempts(t *testing.T) {
	// Mock the necessary variables
	maxAttempts := 3
	delay := time.Second
	var attempts int

	// Create a test function that always returns an error
	testFunc := func() error {
		attempts++
		return errors.New("test error")
	}

	// Call the DoWithAttempts function
	err := DoWithAttempts(testFunc, maxAttempts, delay)

	// Check if the expected number of attempts were made
	if attempts != maxAttempts {
		t.Errorf("Expected %d attempts, but got %d", maxAttempts, attempts)
	}

	// Check for the expected error
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}
