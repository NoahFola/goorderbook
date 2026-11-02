package orderbook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimitAddOrder(t *testing.T) {
	l := NewLimit(10000)
	n := 1000
	size := 50.0
	for i := 0; i < n; i++ {
		o := NewAskOrder(size)
		l.addOrder(o)
		assert.Equal(t, i, o.limitIndex)
	}
	assert.Equal(t, n, len(l.orders))
	assert.Equal(t, float64(n)*size, l.totalVolume)
}

func TestLimitDeleteOrder(t *testing.T) {
	l := NewLimit(20000)
	n := 1000
	size := 0.5
	for i := 0; i < n; i++ {
		o := NewAskOrder(size)
		assert.Equal(t, o.size, l.totalVolume)
		assert.Equal(t, 1, len(l.orders))
		l.deleteOrder(o)
		assert.Equal(t, 0.0, l.totalVolume)
		assert.Equal(t, 0, len(l.orders))
	}
}
