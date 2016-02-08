package simplestats

import (
	"fmt"
	"sync"
)

// Stats - struct that wraps all stats
type Stats struct {
	data  map[string]int64
	mutex sync.RWMutex
}

// String - format all stats as single string
func (s *Stats) String() (output string) {
	for key, value := range s.data {
		output += fmt.Sprintf("%s => %d, ", key, value)
	}
	return output
}

// New - simple stats
func New() (s *Stats) {
	return &Stats{
		data: make(map[string]int64),
	}
}

// Increment - increment key count by 1
func (s *Stats) Increment(key string) {
	s.IncrementBy(key, 1)
}

// IncrementBy - increment key by specified count
func (s *Stats) IncrementBy(key string, value int64) {
	s.mutex.Lock()
	s.data[key] += value
	s.mutex.Unlock()
}

// Get - returns count of key
func (s *Stats) Get(key string) int64 {
	s.mutex.Lock()
	value := s.data[key]
	s.mutex.Unlock()
	return value
}
