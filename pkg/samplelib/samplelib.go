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

type Nullable interface {
	Null() bool
}

type nullable struct {
	isNull bool
}

func (n *nullable) Null() bool {
	return n.isNull
}

type Entity interface {
	Nullable
	Name() string
	Path() string
}

type entity struct {
	nullable
	name string
	path string
}

func newEntity(name string, path string) entity {
	return entity{
		name:     name,
		path:     path,
		nullable: nullable{false},
	}
}

func (e *entity) Name() string {
	return e.name
}
func (e *entity) Path() string {
	return e.path
}

//go:generate mockgen -destination=../../mocks/samplelib/sample.go . Sample
type Sample interface {
	Entity
}

type sample struct {
	entity
}

func newSample(name string, path string) Sample {
	s := sample{
		entity: newEntity(name, path),
	}
	return &s
}

//go:generate mockgen -destination=../../mocks/samplelib/node.go . Node
type Node interface {
	Entity
	Parent() Node
}

type node struct {
	entity
	parent Node
}

func newNode(name string, path string, parent Node) node {
	return node{entity: newEntity(name, path), parent: parent}
}

func (n *node) Parent() Node {
	return n.parent
}

func NullNode() Node {
	rv := &node{entity: entity{name: "", path: "", nullable: nullable{isNull: true}}}
	rv.parent = rv
	return rv
}

//go:generate mockgen -destination=../../mocks/samplelib/meta.go . Meta
type Meta interface {
	Entity
	Description() string
	Keywords() []string
}

type meta struct {
	entity
	description string
	keywords    []string
}

func (m *meta) Description() string {
	return m.description
}

func (m *meta) Keywords() []string {
	return m.keywords
}

//go:generate mockgen -destination=../../mocks/samplelib/samplemeta.go . SampleMeta
type SampleMeta interface {
	Meta
}

type sampleMeta struct {
	meta
	sample Sample
}

func newSampleMeta(sample Sample, description string, keywords []string) sampleMeta {
	return sampleMeta{
		meta: meta{
			description: description,
			keywords:    keywords,
			entity:      entity{name: sample.Name(), path: sample.Path(), nullable: nullable{isNull: false}},
		},
		sample: sample,
	}
}

//go:generate mockgen -destination=../../mocks/samplelib/datasource.go . DataSource
type DataSource interface {
	RootNode() (Node, error)
	ChildrenOf(node Node) ([]Node, error)
	SamplesOf(node Node) ([]Sample, error)
	MetaOf(sample Sample) (SampleMeta, error)
}
