package psql

import (
	"errors"
	"testing"

	"github.com/jackc/pgconn"
)

func TestParsePgError(t *testing.T) {
	// Test case 1: Error is not of type *pgconn.PgError
	err := errors.New("some error")
	result := ParsePgError(err)
	if !errors.Is(result, err) {
		t.Errorf("Expected error: %v, but got: %v", err, result)
	}

	// Test case 2: Error is of type *pgconn.PgError
	pgErr := &pgconn.PgError{
		Message: "Test error message",
		Detail:  "Test error detail",
		Where:   "Test error location",
	}

	err = pgErr
	result = ParsePgError(err)
	expected := "database error. message:Test error message, detail:Test error detail, where:Test error location, sqlstate:"
	if result.Error() != expected {
		t.Errorf("Expected error: %s, but got: %s", expected, result.Error())
	}
}

func BenchmarkParsePgError(b *testing.B) {
	err := &pgconn.PgError{
		Message: "Test error message",
		Detail:  "Test error detail",
		Where:   "Test error location",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParsePgError(err)
	}
}

/*func BenchmarkParsePgErrorOld(b *testing.B) {
	err := &pgconn.PgError{
		Message: "Test error message",
		Detail:  "Test error detail",
		Where:   "Test error location",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParsePgErrorOld(err)
	}
}*/
