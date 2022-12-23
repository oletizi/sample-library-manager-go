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
	"testing"
)

func TestNullNode(t *testing.T) {
	null := NullNode()
	assert.True(t, null.Null())
}

func TestNode(t *testing.T) {
	node := &node{
		entity: entity{name: "name"},
		parent: NullNode(),
	}
	assert.False(t, node.Null())
	assert.NotNil(t, node.Parent())
	assert.Equal(t, node.name, node.Name())
	assert.Equal(t, node.path, node.Path())
}

func TestNewSample(t *testing.T) {
	name := "the name"
	path := "the path"
	sample := &sample{
		entity: entity{name: name, path: path, nullable: nullable{isNull: false}},
	}
	assert.NotNil(t, sample)
	assert.False(t, sample.Null())
	assert.Equal(t, sample.Name(), name)
	assert.Equal(t, sample.Path(), path)
}

func TestEqual(t *testing.T) {
	name1 := "name 1"
	name2 := "name 2"
	path1 := ""
	path2 := "path 2"

	sN1 := nullSample()
	sNull1 := &sN1
	assert.True(t, sNull1.Equal(sNull1))
	sN2 := nullSample()
	sNull2 := &sN2
	assert.True(t, sNull1.Equal(sNull2))
	assert.False(t, sNull1.Equal(nil))

	s1 := newSample(name1, path1)
	sample1 := &s1
	assert.True(t, sample1.Equal(sample1))
	assert.False(t, sample1.Equal(sNull1))

	s2 := newSample(name1, path1)
	sample2 := &s2
	assert.True(t, sample2.Equal(sample2))

	s3 := newSample(name2, path2)
	sample3 := &s3
	assert.False(t, sample1.Equal(sample3))

	n1 := newNode(name1, path1, NullNode())
	node1 := &n1
	assert.True(t, node1.Equal(node1))

	n2 := newNode(name2, path2, NullNode())
	node2 := &n2
	assert.False(t, node1.Equal(node2))

}

func TestNullConstructors(t *testing.T) {
	assert.True(t, NullNullable().Null())
	assert.True(t, NullEntity().Null())
	assert.True(t, NullNode().Null())
	assert.True(t, NullSample().Null())
	assert.True(t, NullMeta().Null())
	assert.True(t, NullSampleMeta().Null())
	assert.True(t, NullAudioStream().Null())
}
