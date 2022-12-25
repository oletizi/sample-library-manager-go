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
	"context"
	"encoding/json"
	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"
	"golang.org/x/exp/slices"
	"gopkg.in/vansante/go-ffprobe.v2"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

type fsDataSource struct {
	root string
}

func (f *fsDataSource) RootNode() (Node, error) {
	_, err := os.ReadDir(f.root)
	if err != nil {
		return nil, err
	}
	node := newNode(path.Base(f.root), f.root, NullNode())
	return &node, nil
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
			children = append(children, &child)
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
			// - be more robust (actually check the file)
			// - be defined publicly somewhere
			if slices.Contains([]string{".wav", ".aif", ".aiff", ".mp3", ".m4a", ".flac"}, path.Ext(item.Name())) {
				sample := newSample(item.Name(), path.Join(node.Path(), item.Name()))
				samples = append(samples, &sample)
			}
		}
	}
	return samples, nil
}

func readFileType(path string) types.Type {
	unknown := types.NewType("unknown", "unknown")
	file, err := os.Open(path)
	if err != nil {
		// notest
		return unknown
	}
	// read enough of the file to get the header
	head := make([]byte, 1024)
	_, err = file.Read(head)
	if err != nil {
		// notest
		return unknown
	}
	match, err := filetype.Match(head)
	if err != nil {
		// notest
		return unknown
	}
	return match
}

func (f *fsDataSource) MetaOf(sample Sample) (SampleMeta, error) {
	fileType := readFileType(sample.Path())

	meta := newSampleMeta(sample, "", []string{}, fileType, NullAudioStream())

	// look for a metadata file for this sample
	// IMPROVE?:
	// - use a file hash instead of a name so files can be moved renamed without losing the mapping
	//   between metadata and media file?
	// - move the metadata directory to the top of the library (à la git) so files can be moved without
	//   also moving their associated metadata file?
	metaPath := path.Join(path.Dir(sample.Path()), ".meta", sample.Name()+".json")
	if _, err := os.Stat(metaPath); err == nil {
		// XXX: This so verbose. There must be a better way to do this.
		// Need to declare this temporary struct b/c the json.Unmarshal function can only write to
		// public fields; the meta struct only has private fields.
		data := struct {
			Description string
			Keywords    []string
		}{Description: "", Keywords: []string{}}

		loadMeta(metaPath, &data)
		meta.description = data.Description
		meta.keywords = data.Keywords
	}
	// Fetch metadata about audio file from the audio file
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	data, err := ffprobe.ProbeURL(ctx, sample.Path())
	if err != nil {
		// notest
		log.Printf("Error getting audio data: %v", err)
	} else {
		stream := data.FirstAudioStream()
		as := newAudioStream(sample)
		as.sampleRate = stream.SampleRate
		as.bitDepth = stream.BitsPerSample
		as.channelCount = stream.Channels
		as.codecName = stream.CodecLongName
		as.codecType = stream.CodecType
		as.duration = stream.Duration
		meta.audioStream = &as
	}
	return &meta, nil
}

func loadMeta(path string, data any) {
	b, err := os.ReadFile(path)
	if err != nil {
		// notest
		log.Panic(err)
	}
	err = json.Unmarshal(b, data)
}

func NewFilesystemDataSource(root string) DataSource {
	return &fsDataSource{root: root}
}
