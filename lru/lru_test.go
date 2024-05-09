package lru

import "testing"

func TestLRU(t *testing.T) {
	cache := New(3)

	cache.Add(1, 1)
	cache.Add(2, 2)
	cache.Add(3, 3)

	value := cache.Get(1)
	if value != 1 {
		t.Errorf("Expected 1, got %d", value)
	}
}

func TestLRU_Evict(t *testing.T) {
	cache := New(3)

	cache.Add(1, 1)
	cache.Add(2, 2)
	cache.Add(3, 3)
	cache.Add(4, 4) // This should evict key 1

	value := cache.Get(1)
	if value != nil {
		t.Errorf("Expected 0, got %d", value)
	}
}
