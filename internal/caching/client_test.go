package cache

import (
	"errors"
	"semrush/internal/customerrors"
	"testing"
	"time"
)

func TestTimeBasedCache(t *testing.T) {

	c := GetCache(3, 2*time.Second, NewTimeEviction())

	c.Set("key1", "value1")
	c.Set("key2", "value2")
	c.Set("key3", "value3")

	if size := c.CacheSize(); size != 3 {
		t.Errorf("Expected cache size 3, got %d", size)
	}

	testCases := []struct {
		key         string
		expVal      interface{}
		expErr      error
		waitingTime time.Duration
	}{
		{"key1", "value1", nil, 0},
		{"key2", "value2", nil, 1},
		{"key3", "value3", nil, 2},
		{"key4", nil, customerrors.ErrorInvalidKey("key4", "not found"), 1},
	}

	for _, tc := range testCases {
		time.Sleep(tc.waitingTime)
		val, err := c.Get(tc.key)
		if val != tc.expVal {
			t.Errorf("Expected value %v for key %s, got %v", tc.expVal, tc.key, val)
		}
		if tc.key == "key4" && errors.Is(err, tc.expErr) {
			t.Errorf("Expected error %v for key %s, got %v", tc.expErr, tc.key, err)
		}
	}

	c.Set("key2", "value2")
	time.Sleep(2 * time.Second)
	if _, err := c.Get("key1"); err == nil {
		t.Errorf("Expected key1 to be expired, but got a value")
	}

	c.Set("key4", "value4")
	c.Set("key5", "value5")
	c.Set("key6", "value6")

	if _, err := c.Get("key2"); err == nil {
		t.Errorf("Expected key2 to be evicted, but got a value")
	}

	if size := c.CacheSize(); size != 3 {
		t.Errorf("Expected final cache size 3, got %d", size)
	}
}

func TestLRUBasedCache(t *testing.T) {

	c := GetCache(3, 2*time.Second, NewLruEviction())
	c.Set("key1", "value1")
	c.Set("key2", "value2")
	c.Set("key3", "value3")

	if size := c.CacheSize(); size != 3 {
		t.Errorf("Expected cache size 3, got %d", size)
	}
	c.Evict() // nothing will happen
}
