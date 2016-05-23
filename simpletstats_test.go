package simplestats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrement(t *testing.T) {
	s := New()

	key := "test_key"

	s.Increment(key)
	s.IncrementBy(key, 10)
	s.Increment(key)

	assert.EqualValues(t, 12, s.Get(key))

	data := s.GetData()
	assert.EqualValues(t, data[key], 12)
}
