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

type Sample struct {
	Name string
	Path string
	null bool
}

func NewSample(name string, path string) *Sample {
	return &Sample{Name: name, Path: path}
}

type Node struct {
	Name   string
	Path   string
	Parent *Node
	null   bool
}

func (n *Node) IsNull() bool {
	return n.null
}

func NullNode() *Node {
	rv := &Node{Name: "", Path: "", null: true}
	rv.Parent = rv
	return rv
}

type DataSource interface {
	RootNode() (*Node, error)
	ChildrenOf(node *Node) ([]*Node, error)
	SamplesOf(node *Node) ([]*Sample, error)
}
