package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddOne(t *testing.T) {
	assert.Equal(t, AddOne(1), 2, "Wrong...")
}

func TestAddOne2(t *testing.T) {
	assert.Equal(t, AddOne2(1), 1, "False...")
}
