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
	"golang.org/x/exp/slices"
	"os"
	"path"
)

type fsDataSource struct {
	root string
}

func (f fsDataSource) RootNode() (*Node, error) {
	_, err := os.ReadDir(f.root)
	if err != nil {
		return nil, err
	}

	rv := &Node{Name: path.Base(f.root), Path: f.root, Parent: NullNode()}
	return rv, nil
}

func (f fsDataSource) ChildrenOf(node *Node) ([]*Node, error) {
	dir, err := os.ReadDir(node.Path)
	if err != nil {
		return nil, err
	}
	children := make([]*Node, 0)
	for _, item := range dir {
		if item.IsDir() {
			// make a new node

			child := &Node{}
			child.Name = item.Name()
			child.Path = path.Join(node.Path, item.Name())
			child.Parent = node
			children = append(children, child)
		}
	}
	return children, nil
}

func (f fsDataSource) SamplesOf(node *Node) ([]*Sample, error) {
	dir, err := os.ReadDir(node.Path)
	if err != nil {
		return nil, err
	}
	samples := make([]*Sample, 0)
	for _, item := range dir {
		if !item.IsDir() {
			// XXX: this set of supported file types should:
			// o be more robust (actually check the file)
			// o be defined publicly somewhere
			if slices.Contains([]string{".wav", ".aif", ".aiff", ".mp3", ".m4a", ".flac"}, path.Ext(item.Name())) {
				sample := NewSample(item.Name(), path.Join(node.Path, item.Name()))
				samples = append(samples, sample)
			}
		}
	}
	return samples, nil
}

func NewFilesystemDataSource(root string) DataSource {
	return fsDataSource{root: root}
}
