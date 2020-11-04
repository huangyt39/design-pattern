package singleton

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetInstance (t *testing.T) {
	s := GetInstance()
	assert.Equal(t, s.GetNum(), 1)
}

func TestGetInstanceWithOnce (t *testing.T) {
	s := GetInstanceWithOnce()
	assert.Equal(t, s.GetNum(), 1)
}