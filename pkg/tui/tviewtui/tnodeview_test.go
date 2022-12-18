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

package tviewtui

import (
	"github.com/golang/mock/gomock"
	mock_samplelib "github.com/oletizi/samplemgr/mocks/samplelib"
	mock_tui "github.com/oletizi/samplemgr/mocks/tui"
	mock_view "github.com/oletizi/samplemgr/mocks/tui/view"
	"github.com/oletizi/samplemgr/pkg/samplelib"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTNodeView_Methods(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	list := tview.NewList()
	display := mock_view.NewMockDisplay(ctl)
	logger := mock_tui.NewMockLogger(ctl)
	eh := mock_tui.NewMockErrorHandler(ctl)
	nv := newTNodeView(list, display, logger, eh)

	assert.NotNil(t, nv)
	assert.Equal(t, list, nv.list)
	assert.Equal(t, display, nv.display)
	assert.Equal(t, logger, nv.logger)
	assert.Equal(t, eh, nv.eh)
	nodeSelected := func(node samplelib.Node) {}
	sampleSelected := func(sample samplelib.Sample) {}
	nodeChosen := func(node samplelib.Node) {}
	sampleChosen := func(sample samplelib.Sample) {}

	ds := mock_samplelib.NewMockDataSource(ctl)
	parent := mock_samplelib.NewMockNode(ctl)
	node := mock_samplelib.NewMockNode(ctl)
	child := mock_samplelib.NewMockNode(ctl)
	var children []samplelib.Node
	children = append(children, child)

	sample := mock_samplelib.NewMockSample(ctl)
	var samples []samplelib.Sample
	samples = append(samples, sample)

	ds.EXPECT().ChildrenOf(node).Return(children, nil)
	ds.EXPECT().SamplesOf(node).Return(samples, nil)
	display.EXPECT().DisplayNodeAsListing(child, false)
	display.EXPECT().DisplaySampleAsListing(sample)
	eh.EXPECT().Handle(nil).Times(2)
	node.EXPECT().Name().Return("the name")
	node.EXPECT().Parent().AnyTimes().Return(parent)
	parent.EXPECT().Null().Return(false)
	nv.UpdateNode(ds, node, nodeSelected, sampleSelected, nodeChosen, sampleChosen)
}
