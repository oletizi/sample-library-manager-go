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
	"encoding/json"
	"golang.org/x/exp/slices"
	"log"
	"os"
	"path"
	"strings"
)

type fsDataSource struct {
	root string
}

func (f *fsDataSource) MetaOf(sample Sample) (SampleMeta, error) {
	metaPath := path.Join(path.Dir(sample.Path()), ".meta", sample.Name()+".json")
	log.Println("meta path: " + metaPath)
	meta := &sampleMeta{}
	err := loadMeta(metaPath, meta)

	if err != nil {
		return nil, err
	} else {
		return meta, err
	}
}

func newNode(name string, path string, parent Node) *node {
	return &node{
		entity: entity{name: name, path: path, nullable: nullable{isNull: false}}, parent: parent}
}

func (f *fsDataSource) RootNode() (Node, error) {
	_, err := os.ReadDir(f.root)
	if err != nil {
		return nil, err
	}

	//rv := &node{
	//	entity: entity{name: path.Base(f.root), path: f.root, nullable: nullable{isNull: false}},
	//	parent: NullNode(),
	//}
	return newNode(path.Base(f.root), f.root, NullNode()),
		nil
}

func (f *fsDataSource) ChildrenOf(parent Node) ([]Node, error) {
	dir, err := os.ReadDir(parent.Path())
	if err != nil {
		return nil, err
	}
	children := make([]Node, 0)
	for _, item := range dir {
		if item.IsDir() && !strings.HasPrefix(item.Name(), ".") {
			child := newNode(item.Name(), path.Join(parent.Path(), item.Name()), parent)
			children = append(children, child)
		}
	}
	return children, nil
}

func (f *fsDataSource) SamplesOf(node Node) ([]Sample, error) {
	dir, err := os.ReadDir(node.Path())
	if err != nil {
		return nil, err
	}
	samples := make([]Sample, 0)
	for _, item := range dir {
		if !item.IsDir() {
			// XXX: this set of supported file types should:
			// o be more robust (actually check the file)
			// o be defined publicly somewhere
			if slices.Contains([]string{".wav", ".aif", ".aiff", ".mp3", ".m4a", ".flac"}, path.Ext(item.Name())) {
				sample := newSample(item.Name(), path.Join(node.Path(), item.Name()))
				samples = append(samples, sample)
			}
		}
	}
	return samples, nil
}

func loadMeta(path string, data any) error {
	file, err := os.Open(path)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Panic(err)
		}
	}(file)

	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	return decoder.Decode(data)
}

func NewFilesystemDataSource(root string) DataSource {
	return &fsDataSource{root: root}
}
