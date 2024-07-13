package customerrors

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidKey = errors.New("key does not exist")
)

type CacheError struct {
	key     interface{}
	ErrType error
	cause   string
}

func newCacheError(errType error, key interface{}, cause string) *CacheError {
	return &CacheError{
		ErrType: errType,
		key:     key,
		cause:   cause,
	}
}

func (c *CacheError) Error() string {
	return fmt.Sprintf("%s: %v: %v", c.ErrType, c.key, c.cause)
}

func (c *CacheError) Unwrap() error {
	return c.ErrType
}

func ErrorInvalidKey(key interface{}, err string) *CacheError {
	return newCacheError(ErrInvalidKey, key, err)
}
