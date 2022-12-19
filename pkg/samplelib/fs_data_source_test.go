/*
 * Copyright (c) 2022 Orion Letizi
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package samplelib

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestFilesystemDataSource(t *testing.T) {
	const libdir = "../../test/data/library/multi-level"
	dataSource := NewFilesystemDataSource(libdir)
	rootNode, err := dataSource.RootNode()
	assert.Nil(t, err)
	assert.NotNil(t, rootNode)
	assert.Equal(t, "multi-level", rootNode.Name())

	// check Children
	children, err := dataSource.ChildrenOf(rootNode)
	assert.Nil(t, err)
	assert.NotNil(t, children)
	assert.Equal(t, 2, len(children))
	assert.Equal(t, "level-2a", children[0].Name())

	// check the Children of a subdirectory
	child := children[0]
	subChildren, err := dataSource.ChildrenOf(child)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(subChildren))

	// check Samples
	samples, err := dataSource.SamplesOf(rootNode)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(samples))
	assert.Equal(t, "cabasa.wav", samples[0].Name())

	// check the Samples of a subdirectory
	samples, err = dataSource.SamplesOf(child)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(samples))
	assert.Equal(t, "kick.wav", samples[0].Name())
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

func TestFsDataSource_MetaOf(t *testing.T) {
	source := NewFilesystemDataSource("../../test/data/library/multi-level")
	root, err := source.RootNode()
	assert.Nil(t, err)

	samples, err := source.SamplesOf(root)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(samples))

	// test a sample with metadata
	sample := samples[0]

	meta, err := source.MetaOf(sample)
	assert.Nil(t, err)
	assert.NotNil(t, meta)
	keywords := meta.Keywords()
	assert.Equal(t, 2, len(keywords))
	assert.Equal(t, "percussion", keywords[0])
	assert.Equal(t, "shaker", keywords[1])

	description := meta.Description()
	assert.Equal(t, "The description for cabasa.wav", description)

	// test a sample without metadata
	sample = samples[1]
	meta, err = source.MetaOf(sample)
	assert.Equal(t, "", meta.Description())
	assert.Equal(t, 0, len(meta.Keywords()))

	stream := meta.AudioStream()
	assert.NotNil(t, stream)
	assert.False(t, stream.Null())
	assert.Equal(t, "44100", stream.SampleRate())
	assert.Equal(t, 16, stream.BitDepth())
	log.Printf("sample rate: %s", stream.SampleRate())
}
