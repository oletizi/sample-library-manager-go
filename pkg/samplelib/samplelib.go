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

import "github.com/h2non/filetype/types"

type Nullable interface {
	Null() bool
}

type nullable struct {
	isNull bool
}

func (n *nullable) Null() bool {
	return n.isNull
}

func newNullable() nullable {
	return nullable{false}
}

func nullNullable() nullable {
	return nullable{true}
}

func NullNullable() Nullable {
	n := nullNullable()
	return &n
}

type Entity interface {
	Nullable
	Name() string
	Path() string
	Equal(e Entity) bool
}

type entity struct {
	nullable
	name string
	path string
}

func (e *entity) Equal(cmp Entity) bool {
	if cmp == nil {
		return false
	}
	if e.Null() {
		return cmp.Null()
	}
	if cmp.Null() {
		return e.Null()
	}
	return e.Path() == cmp.Path()
}

func newEntity(name string, path string) entity {
	return entity{
		name:     name,
		path:     path,
		nullable: nullable{false},
	}
}

func nullEntity() entity {
	return entity{name: "", path: "", nullable: nullable{true}}
}

func NullEntity() Entity {
	n := nullEntity()
	return &n
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

func newSample(name string, path string) sample {
	s := sample{
		entity: newEntity(name, path),
	}
	return s
}

func nullSample() sample {
	return sample{entity: nullEntity()}
}

func NullSample() Sample {
	s := nullSample()
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

func nullNode() node {
	n := node{entity: nullEntity()}
	n.parent = &n
	return n
}

func NullNode() Node {
	n := nullNode()
	return &n
}

func (n *node) Parent() Node {
	return n.parent
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

func newMeta(name string, path string, description string, keywords []string) meta {
	return meta{entity: newEntity(name, path), description: description, keywords: keywords}
}

func nullMeta() meta {
	return meta{entity: nullEntity(), description: "", keywords: []string{}}
}

func NullMeta() Meta {
	m := nullMeta()
	return &m
}

func (m *meta) Description() string {
	return m.description
}

func (m *meta) Keywords() []string {
	return m.keywords
}

type FileType struct {
	types.Type
}

//go:generate mockgen -destination=../../mocks/samplelib/samplemeta.go . SampleMeta
type SampleMeta interface {
	Meta
	FileType() FileType
	AudioStream() AudioStream
}

type sampleMeta struct {
	meta
	sample      Sample
	audioStream AudioStream
	filetype    FileType
}

func (s *sampleMeta) FileType() FileType {
	return s.filetype
}
func (s *sampleMeta) AudioStream() AudioStream {
	return s.audioStream
}

func newSampleMeta(sample Sample, description string, keywords []string, filetype types.Type, stream AudioStream) sampleMeta {
	return sampleMeta{
		meta:        newMeta(sample.Name(), sample.Path(), description, keywords),
		sample:      sample,
		filetype:    FileType{filetype},
		audioStream: stream,
	}
}

func nullSampleMeta() sampleMeta {
	return sampleMeta{meta: nullMeta(), sample: NullSample(), audioStream: NullAudioStream()}
}

func NullSampleMeta() SampleMeta {
	s := nullSampleMeta()
	return &s
}

//go:generate mockgen -destination=../../mocks/samplelib/audiostream.go . AudioStream
type AudioStream interface {
	Nullable
	SampleRate() string
	BitDepth() int
	ChannelCount() int
	CodecName() string
	CodecType() string
	Duration() string
}

func (s *audioStream) SampleRate() string {
	return s.sampleRate
}

func (s *audioStream) BitDepth() int {
	return s.bitDepth
}

type audioStream struct {
	nullable
	sample       Sample
	sampleRate   string
	bitDepth     int
	channelCount int
	codecName    string
	codecType    string
	duration     string
}

func (s *audioStream) ChannelCount() int {
	return s.channelCount
}

func (s *audioStream) CodecName() string {
	return s.codecName
}

func (s *audioStream) CodecType() string {
	return s.codecType
}

func (s *audioStream) Duration() string {
	return s.duration
}

// newAudioStream internal constructor for audioStream struct
func newAudioStream(sample Sample) audioStream {
	return audioStream{
		nullable: newNullable(),
		sample:   sample,
	}
}
func nullAudioStream() audioStream {
	return audioStream{sample: NullSample(), sampleRate: "", nullable: nullNullable()}
}

func NullAudioStream() AudioStream {
	n := nullAudioStream()
	return &n
}

//go:generate mockgen -destination=../../mocks/samplelib/datasource.go . DataSource
type DataSource interface {
	RootNode() (Node, error)
	ChildrenOf(node Node) ([]Node, error)
	SamplesOf(node Node) ([]Sample, error)
	MetaOf(sample Sample) (SampleMeta, error)
}
