package samplelib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNode(t *testing.T) {
	name := "the name"
	path := "the path"
	parent := NullNode()

	children := make([]*Node, 0)
	samples := make([]*Sample, 0)
	node := NewNode(name, path, parent, children, samples)
	assert.NotNil(t, node)
	assert.Equal(t, node.name, name)
	assert.Equal(t, node.path, path)
}

func TestNewSample(t *testing.T) {
	name := "the name"
	path := "the path"
	sample := NewSample(name, path)
	assert.NotNil(t, sample)
	assert.Equal(t, sample.name, name)
	assert.Equal(t, sample.path, path)
}
