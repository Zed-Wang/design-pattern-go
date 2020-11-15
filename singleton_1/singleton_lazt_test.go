package singleton_test

import (
	singleton "design-pattern-go/singleton_1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLazyInstance(t *testing.T) {
	assert.Equal(t, singleton.GetLazyInstance(), singleton.GetLazyInstance())
}

func BenchmarkGetLazyInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			assert.Equal(b, singleton.GetLazyInstance(), singleton.GetLazyInstance())
		}
	})
}
