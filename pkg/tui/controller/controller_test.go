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

package controller

import (
	"github.com/golang/mock/gomock"
	mock_samplelib "github.com/oletizi/samplemgr/mocks/samplelib"
	mock_tui "github.com/oletizi/samplemgr/mocks/tui"
	mock_view "github.com/oletizi/samplemgr/mocks/tui/view"
	"github.com/oletizi/samplemgr/pkg/tui"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	ctl := gomock.NewController(t)
	ds := mock_samplelib.NewMockDataSource(ctl)
	errorHandler := mock_tui.NewMockErrorHandler(ctl)
	nodeView := mock_view.NewMockNodeView(ctl)
	infoView := mock_view.NewMockInfoView(ctl)
	logView := mock_view.NewMockLogView(ctl)

	c := New(ds, errorHandler, nodeView, infoView, logView)
	assert.NotNil(t, c)
}

func TestController_UpdateNode(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	ds := mock_samplelib.NewMockDataSource(ctl)
	eh := mock_tui.NewMockErrorHandler(ctl)
	nodeView := mock_view.NewMockNodeView(ctl)
	infoView := mock_view.NewMockInfoView(ctl)
	logView := mock_view.NewMockLogView(ctl)
	node := mock_samplelib.NewMockNode(ctl)

	// make a new controller
	c := &controller{
		ds:     ds,
		nv:     nodeView,
		iv:     infoView,
		lv:     logView,
		eh:     eh,
		logger: log.Default(),
	}
	// XXX: can't figure out how to get function arguments to match
	nodeView.EXPECT().UpdateNode(ds, node, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1)
	infoView.EXPECT().UpdateNode(node)
	node.EXPECT().Name()
	assert.NotNil(t, c)
	c.UpdateNode(node)
}

func TestController_selectNodeAndSample(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	infoView := mock_view.NewMockInfoView(ctl)
	node := mock_samplelib.NewMockNode(ctl)
	sample := mock_samplelib.NewMockSample(ctl)

	c := &controller{
		iv: infoView,
	}

	infoView.EXPECT().UpdateNode(node)
	c.nodeSelected(node)

	infoView.EXPECT().UpdateSample(sample)
	c.sampleSelected(sample)
}

func TestController_chooseNodeAndSample(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ds := mock_samplelib.NewMockDataSource(ctl)
	nodeView := mock_view.NewMockNodeView(ctl)
	infoView := mock_view.NewMockInfoView(ctl)
	node := mock_samplelib.NewMockNode(ctl)
	sample := mock_samplelib.NewMockSample(ctl)

	c := &controller{
		ds:     ds,
		nv:     nodeView,
		iv:     infoView,
		logger: tui.NewLogger(log.Default()),
	}

	nodeView.EXPECT().UpdateNode(ds, node, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())
	infoView.EXPECT().UpdateNode(node)
	node.EXPECT().Name().Return("node name").AnyTimes()
	c.nodeChosen(node)

	infoView.EXPECT().UpdateSample(sample)
	c.sampleChosen(sample)
}
