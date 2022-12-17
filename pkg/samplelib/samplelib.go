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

type Entity interface {
	Nullable
	Name() string
	Path() string
}

//go:generate mockgen -destination=../../mocks/samplelib/sample.go . Sample
type Sample interface {
	Entity
}

type sample struct {
	name string
	path string
	null bool
}

func (s *sample) Null() bool {
	return s.null
}

func (s *sample) Name() string {
	return s.name
}

func (s *sample) Path() string {
	return s.path
}

func NewSample(name string, path string) Sample {
	return &sample{name: name, path: path}
}

//go:generate mockgen -destination=../../mocks/samplelib/node.go . Node
type Node interface {
	Entity
	Parent() Node
}

type node struct {
	name   string
	path   string
	parent Node
	null   bool
}

func (n *node) Name() string {
	return n.name
}

func (n *node) Path() string {
	return n.path
}

func (n *node) Parent() Node {
	return n.parent
}

func (n *node) Null() bool {
	return n.null
}

func NullNode() Node {
	rv := &node{name: "", path: "", null: true}
	rv.parent = rv
	return rv
}

//go:generate mockgen -destination=../../mocks/samplelib/datasource.go . DataSource
type DataSource interface {
	RootNode() (Node, error)
	ChildrenOf(node Node) ([]Node, error)
	SamplesOf(node Node) ([]Sample, error)
}
