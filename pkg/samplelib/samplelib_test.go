package samplelib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNullNode(t *testing.T) {
	null := NullNode()
	assert.True(t, null.null)
}

func TestNewSample(t *testing.T) {
	name := "the name"
	path := "the path"
	sample := NewSample(name, path)
	assert.NotNil(t, sample)
	assert.Equal(t, sample.Name, name)
	assert.Equal(t, sample.Path, path)
}
