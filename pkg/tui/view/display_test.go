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

package view

import (
	"github.com/golang/mock/gomock"
	mocksamplelib "github.com/oletizi/samplemgr/mocks/samplelib"
	mocktui "github.com/oletizi/samplemgr/mocks/tui"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"strings"
	"testing"
	"text/template"
)

func TestRender(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	errorHandler := mocktui.NewMockErrorHandler(ctl)
	tmpl, err := template.New("test").Parse("{{ .YourNotGoingToFindIt }}")
	assert.Nil(t, err)

	errorHandler.EXPECT().Handle(gomock.Any())
	render(tmpl, struct{}{}, errorHandler)
}

func TestSampleDisplay(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ds := mocksamplelib.NewMockDataSource(ctl)
	sampleMeta := mocksamplelib.NewMockSampleMeta(ctl)
	audioStream := mocksamplelib.NewMockAudioStream(ctl)
	samplePath := "the sample path"
	sampleName := "the sample name"
	sample := mocksamplelib.NewMockSample(ctl)
	errorHandler := mocktui.NewMockErrorHandler(ctl)

	sample.EXPECT().Path().Return(samplePath)
	sample.EXPECT().Name().Return(sampleName)
	errorHandler.EXPECT().Handle(gomock.Any()).MinTimes(1)

	ds.EXPECT().MetaOf(sample).Return(sampleMeta, nil)

	keywords := []string{"keyword1", "keyword2"}
	sampleMeta.EXPECT().Keywords().Return(keywords)
	sampleMeta.EXPECT().AudioStream().MinTimes(1).Return(audioStream)
	sampleMeta.EXPECT().FileType()
	sampleRate := "100"
	bitDepth := 8
	audioStream.EXPECT().SampleRate().Return(sampleRate)
	audioStream.EXPECT().BitDepth().Return(bitDepth)
	audioStream.EXPECT().ChannelCount()
	audioStream.EXPECT().Duration()
	audioStream.EXPECT().CodecName()
	// do the thing!
	display, err := NewDisplay(log.Default(), errorHandler)

	assert.Nil(t, err)
	assert.NotNil(t, display)
	text := display.DisplaySampleAsText(ds, sample)
	assert.True(t, strings.Contains(text, samplePath))
	assert.True(t, strings.Contains(text, sampleName))
	assert.True(t, strings.Contains(text, sampleRate))
	assert.True(t, strings.Contains(text, strconv.Itoa(bitDepth)))

	sample.EXPECT().Name().Return(sampleName)
	text = display.DisplaySampleAsListing(sample)
	assert.True(t, strings.Contains(text, sampleName))
}

func TestNodeDisplay(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	errorHandler := mocktui.NewMockErrorHandler(ctl)

	nodePath := "the node path"
	nodeName := "the node name"
	node := mocksamplelib.NewMockNode(ctl)
	node.EXPECT().Path().Return(nodePath)
	node.EXPECT().Name().Return(nodeName)

	errorHandler.EXPECT().Handle(gomock.Any()).MinTimes(1)
	d, err := NewDisplay(log.Default(), errorHandler)
	assert.Nil(t, err)

	text := d.DisplayNodeAsText(nil, node)
	assert.NotNil(t, text)
	assert.True(t, strings.Contains(text, nodePath))
	assert.True(t, strings.Contains(text, nodeName))

	node.EXPECT().Name().Return(nodeName)
	text = d.DisplayNodeAsListing(node, false)
	assert.NotNil(t, text)
	assert.True(t, strings.Contains(text, nodeName))

	text = d.DisplayNodeAsListing(node, true)
	assert.NotNil(t, text)
	assert.Equal(t, "..", text)
}
