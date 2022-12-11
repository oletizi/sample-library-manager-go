package samplelib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilesystemDataSource(t *testing.T) {
	const libdir = "../../test/data/library/multi-level"
	dataSource := NewFilesystemDataSource(libdir)
	rootNode, err := dataSource.RootNode()
	assert.Nil(t, err)
	assert.NotNil(t, rootNode)
	assert.Equal(t, "multi-level", rootNode.Name)

	// check Children
	children, err := dataSource.ChildrenOf(rootNode)
	assert.Nil(t, err)
	assert.NotNil(t, children)
	assert.Equal(t, 2, len(children))
	assert.Equal(t, "level-2a", children[0].Name)

	// check the Children of a subdirectory
	child := children[0]
	subChildren, err := dataSource.ChildrenOf(child)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(subChildren))

	// check Samples
	samples, err := dataSource.SamplesOf(rootNode)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(samples))
	assert.Equal(t, "cabasa.wav", samples[0].Name)

	// check the Samples of a subdirectory
	samples, err = dataSource.SamplesOf(child)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(samples))
	assert.Equal(t, "kick.wav", samples[0].Name)
}

func TestFileSystemDataSourceErrors(t *testing.T) {
	source := NewFilesystemDataSource("a path that points to nothing")
	node, err := source.RootNode()
	assert.Nil(t, node)
	assert.NotNil(t, err)

	children, err := source.ChildrenOf(NullNode())
	assert.Nil(t, children)
	assert.NotNil(t, err)

	samples, err := source.SamplesOf(NullNode())
	assert.Nil(t, samples)
	assert.NotNil(t, err)
}
