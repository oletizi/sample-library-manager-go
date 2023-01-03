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
	mock_audio "github.com/oletizi/samplemgr/mocks/audio"
	mocksamplelib "github.com/oletizi/samplemgr/mocks/samplelib"
	mocktui "github.com/oletizi/samplemgr/mocks/tui"
	mockview "github.com/oletizi/samplemgr/mocks/tui/view"
	mock_util "github.com/oletizi/samplemgr/mocks/util"
	"github.com/oletizi/samplemgr/pkg/util"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	ctl := gomock.NewController(t)
	ac := mock_audio.NewMockContext(ctl)
	ds := mocksamplelib.NewMockDataSource(ctl)
	errorHandler := mocktui.NewMockErrorHandler(ctl)
	nodeView := mockview.NewMockNodeView(ctl)
	infoView := mockview.NewMockInfoView(ctl)
	logView := mockview.NewMockLogView(ctl)

	nodeView.EXPECT().Focus()
	c := New(ac, ds, errorHandler, nodeView, infoView, logView)
	assert.NotNil(t, c)
}

func TestController_UpdateNode(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	ds := mocksamplelib.NewMockDataSource(ctl)
	eh := mocktui.NewMockErrorHandler(ctl)
	nodeView := mockview.NewMockNodeView(ctl)
	infoView := mockview.NewMockInfoView(ctl)
	logView := mockview.NewMockLogView(ctl)
	node := mocksamplelib.NewMockNode(ctl)

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
	infoView.EXPECT().UpdateNode(ds, node)
	assert.NotNil(t, c)
	c.UpdateNode(node)
}

func TestController_selectNodeAndSample(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ds := mocksamplelib.NewMockDataSource(ctl)
	infoView := mockview.NewMockInfoView(ctl)
	node := mocksamplelib.NewMockNode(ctl)
	sample := mocksamplelib.NewMockSample(ctl)

	c := &controller{
		ds: ds,
		iv: infoView,
	}

	infoView.EXPECT().UpdateNode(ds, node)
	c.nodeSelected(node)

	infoView.EXPECT().UpdateSample(ds, sample)
	c.sampleSelected(sample)
}

func TestChooseSameSample(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	sample := mocksamplelib.NewMockSample(ctl)
	audioContext := mock_audio.NewMockContext(ctl)
	audioPlayer := mock_audio.NewMockPlayer(ctl)
	errorHandler := mocktui.NewMockErrorHandler(ctl)
	playQueue := mock_util.NewMockQueue(ctl)
	c := &controller{
		playQueue:     playQueue,
		logger:        log.Default(),
		ac:            audioContext,
		currentPlayer: audioPlayer,
		eh:            errorHandler,
	}

	playQueue.EXPECT().Add(sample)
	c.sampleChosen(sample)

	sample.EXPECT().Equal(gomock.Any()).Return(true)
	audioPlayer.EXPECT().Playing().Return(true)
	audioPlayer.EXPECT().Stop()
	errorHandler.EXPECT().Handle(gomock.Any()).AnyTimes()
	// return sample after the first get
	playQueue.EXPECT().Get().Return(sample, false)
	playQueue.EXPECT().Done(sample)
	// shutdown after the first get
	playQueue.EXPECT().Get().Return(nil, true)
	c.playLoop()

}

func TestController_chooseNodeAndSample(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	audioContext := mock_audio.NewMockContext(ctl)
	ds := mocksamplelib.NewMockDataSource(ctl)
	nodeView := mockview.NewMockNodeView(ctl)
	infoView := mockview.NewMockInfoView(ctl)
	node := mocksamplelib.NewMockNode(ctl)
	sample := mocksamplelib.NewMockSample(ctl)
	errorHandler := mocktui.NewMockErrorHandler(ctl)
	playQueue := mock_util.NewMockQueue(ctl)

	c := &controller{
		playQueue: playQueue,
		eh:        errorHandler,
		ac:        audioContext,
		ds:        ds,
		nv:        nodeView,
		iv:        infoView,
		logger:    util.NewLogger(log.Default()),
	}

	errorHandler.EXPECT().Handle(gomock.Any()).AnyTimes()

	// Test nodeChosen(Node)
	nodeView.EXPECT().UpdateNode(ds, node, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())
	infoView.EXPECT().UpdateNode(ds, node)
	node.EXPECT().Name().Return("node name").AnyTimes()
	c.nodeChosen(node)

	// Test sampleChosen(Sample)
	playQueue.EXPECT().Add(sample)
	c.sampleChosen(sample)

	// test play loop
	audioPlayer := mock_audio.NewMockPlayer(ctl)
	samplePath := "sample path"
	sample.EXPECT().Path().MinTimes(1).Return(samplePath)
	// return the sample to play the first time through the play loop
	playQueue.EXPECT().Get().Return(sample, false)
	playQueue.EXPECT().Done(sample)
	// return shutdown the next time through the loop
	playQueue.EXPECT().Get().Return(nil, true)
	audioContext.EXPECT().PlayerFor(samplePath).Return(audioPlayer, nil)
	audioPlayer.EXPECT().Play(gomock.Any())
	log.Println("About to call playLoop...")
	c.playLoop()
}
