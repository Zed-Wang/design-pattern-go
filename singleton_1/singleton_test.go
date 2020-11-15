package singleton_test

import (
	singleton "design-pattern-go/singleton_1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInstance(t *testing.T) {
	assert.Equal(t, singleton.GetInstance(), singleton.GetInstance())
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			assert.Equal(b, singleton.GetInstance(), singleton.GetInstance())
		}
	})
}
