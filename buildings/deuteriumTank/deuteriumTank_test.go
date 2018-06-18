package deuteriumTank

import (
	"testing"

	"github.com/alaingilbert/ogame"
	"github.com/stretchr/testify/assert"
)

func TestCost(t *testing.T) {
	dt := New()
	assert.Equal(t, ogame.Resources{Metal: 1000, Crystal: 1000}, dt.GetPrice(1))
	assert.Equal(t, ogame.Resources{Metal: 2000, Crystal: 2000}, dt.GetPrice(2))
	assert.Equal(t, ogame.Resources{Metal: 4000, Crystal: 4000}, dt.GetPrice(3))
	assert.Equal(t, ogame.Resources{Metal: 8000, Crystal: 8000}, dt.GetPrice(4))
	assert.Equal(t, ogame.Resources{Metal: 16000, Crystal: 16000}, dt.GetPrice(5))
	assert.Equal(t, ogame.Resources{Metal: 32000, Crystal: 32000}, dt.GetPrice(6))
}

func TestCapacity(t *testing.T) {
	dt := New()
	assert.Equal(t, 10000, dt.Capacity(0))
	assert.Equal(t, 20000, dt.Capacity(1))
	assert.Equal(t, 40000, dt.Capacity(2))
	assert.Equal(t, 75000, dt.Capacity(3))
	assert.Equal(t, 140000, dt.Capacity(4))
	assert.Equal(t, 255000, dt.Capacity(5))
}
