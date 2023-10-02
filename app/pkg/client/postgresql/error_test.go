package psql

import (
	"errors"
	"testing"
)

func TestErrCommit(t *testing.T) {
	err := errors.New("original error")
	result := ErrCommit(err)

	if result.Error() != "failed to commit Tx: original error" {
		t.Errorf("Expected error: %s, but got: %s", "failed to commit Tx: original error", result.Error())
	}
}

func TestErrRollback(t *testing.T) {
	err := errors.New("original error")
	result := ErrRollback(err)

	if result.Error() != "failed to rollback Tx: original error" {
		t.Errorf("Expected error: %s, but got: %s", "failed to rollback Tx: original error", result.Error())
	}
}

func TestErrCreateTx(t *testing.T) {
	err := errors.New("original error")
	result := ErrCreateTx(err)

	if result.Error() != "failed to create Tx: original error" {
		t.Errorf("Expected error: %s, but got: %s", "failed to create Tx: original error", result.Error())
	}
}

func TestErrCreateQuery(t *testing.T) {
	err := errors.New("original error")
	result := ErrCreateQuery(err)

	if result.Error() != "failed to create SQL Query: original error" {
		t.Errorf("Expected error: %s, but got: %s", "failed to create SQL Query: original error", result.Error())
	}
}

func TestErrScan(t *testing.T) {
	err := errors.New("original error")
	result := ErrScan(err)

	if result.Error() != "failed to scan: original error" {
		t.Errorf("Expected error: %s, but got: %s", "failed to scan: original error", result.Error())
	}
}

func TestErrExec(t *testing.T) {
	err := errors.New("original error")
	result := ErrExec(err)

	if result.Error() != "failed to execute: original error" {
		t.Errorf("Expected error: %s, but got: %s", "failed to execute: original error", result.Error())
	}
}

func TestErrDoQuery(t *testing.T) {
	err := errors.New("original error")
	result := ErrDoQuery(err)

	if result.Error() != "failed to query: original error" {
		t.Errorf("Expected error: %s, but got: %s", "failed to query: original error", result.Error())
	}
}
