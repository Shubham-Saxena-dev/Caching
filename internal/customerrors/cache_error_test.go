package customerrors

import (
	"errors"
	"testing"
)

func TestErrorInvalidKey(t *testing.T) {
	key := "testKey"
	errMsg := "not found"

	err := ErrorInvalidKey(key, errMsg)
	var cacheErr *CacheError
	ok := errors.As(err, &cacheErr)
	if !ok {
		t.Fatalf("Expected error of type *CacheError, got %T", err)
	}
	if cacheErr.key != key {
		t.Errorf("Expected key %v, got %v", key, cacheErr.key)
	}
	if cacheErr.cause != errMsg {
		t.Errorf("Expected cause %v, got %v", errMsg, cacheErr.cause)
	}
}
