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
	mock_tui "github.com/oletizi/samplemgr/mocks/tui"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
	"text/template"
)

func TestRender(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	errorHandler := mock_tui.NewMockErrorHandler(ctl)
	tmpl, err := template.New("test").Parse("{{ .YourNotGoingToFindIt }}")
	assert.Nil(t, err)

	errorHandler.EXPECT().Handle(gomock.Any())
	render(tmpl, struct{}{}, errorHandler)
}

func TestNodeDisplay(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	errorHandler := mock_tui.NewMockErrorHandler(ctl)

	nodePath := "the node path"
	nodeName := "the node name"
	node := mocksamplelib.NewMockNode(ctl)
	node.EXPECT().Path().Return(nodePath)
	node.EXPECT().Name().Return(nodeName)

	errorHandler.EXPECT().Handle(gomock.Any()).MinTimes(1)
	d, err := NewDisplay(log.Default(), errorHandler)
	assert.Nil(t, err)

	text := d.DisplayNodeAsText(node)
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
