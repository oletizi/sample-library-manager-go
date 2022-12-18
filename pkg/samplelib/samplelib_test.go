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

func TestNullConstructors(t *testing.T) {
	assert.True(t, NullNullable().Null())
	assert.True(t, NullEntity().Null())
	assert.True(t, NullNode().Null())
	assert.True(t, NullSample().Null())
	assert.True(t, NullMeta().Null())
	assert.True(t, NullSampleMeta().Null())
	assert.True(t, NullAudioStream().Null())
}
