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

package audio

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/golang/mock/gomock"
	mockbp "github.com/oletizi/samplemgr/mocks/audio/bp"
	mockutil "github.com/oletizi/samplemgr/mocks/util"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestBeepPlayer_Playing(t *testing.T) {
	player := &beepPlayer{}

	playing := player.Playing()
	assert.False(t, playing)

	ctrl := &beep.Ctrl{}
	ctrl.Paused = false
	player.ctl = ctrl

	playing = player.Playing()
	assert.True(t, playing)
}

func TestBeepContext_PlayerFor(t *testing.T) {
	ctx := &beepContext{}
	player, err := ctx.PlayerFor("../../test/data/library/multi-level/hh.wav")
	assert.Nil(t, err)
	assert.NotNil(t, player)
}

func TestBeepPlayer_InterfaceFunctions(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	controlQueue := mockutil.NewMockQueue(ctl)
	streamer := mockbp.NewMockStreamer(ctl)
	speaker := mockbp.NewMockSpeaker(ctl)
	player := &beepPlayer{
		controlQueue:      controlQueue,
		spk:               speaker,
		speakerSampleRate: 44100,
		streamer:          streamer,
		format:            &beep.Format{SampleRate: 44100, NumChannels: 2, Precision: 4},
	}

	callback := func() {}
	// XXX: Can't figure out how to get arg matchers to work right
	controlQueue.EXPECT().Add(gomock.Any())
	player.Play(&callback)

	controlQueue.EXPECT().Add(gomock.Any())
	player.Loop(1, &callback)

	controlQueue.EXPECT().Add(gomock.Any())
	player.Pause()

	controlQueue.EXPECT().Add(gomock.Any())
	player.Stop()

	controlQueue.EXPECT().Add(gomock.Any())
	player.Close()

}

func TestTransportOperation(t *testing.T) {
	operation := transportOperation{op: doStop}
	assert.NotNil(t, operation.timestamp)
	fmt.Printf("Timestampe: %s\n", operation.timestamp)
	fmt.Printf("Timestampe: %s\n", operation.timestamp)
}

func TestControlLoop(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	speaker := mockbp.NewMockSpeaker(ctl)
	controlQueue := mockutil.NewMockQueue(ctl)
	streamer := mockbp.NewMockStreamer(ctl)
	beepCtl := beep.Ctrl{Streamer: streamer, Paused: false}
	player := &beepPlayer{
		logger:            log.Default(),
		spk:               speaker,
		controlQueue:      controlQueue,
		streamer:          streamer,
		ctl:               &beepCtl,
		format:            &beep.Format{SampleRate: 44100, NumChannels: 2, Precision: 4},
		speakerSampleRate: 44100,
	}
	callback := func() {}
	// test loop
	loopOp := newTransportOperation(doLoop, 1, &callback)
	controlQueue.EXPECT().Get().Return(loopOp, false)
	controlQueue.EXPECT().Done(loopOp)
	speaker.EXPECT().Lock()
	streamer.EXPECT().Seek(0)
	speaker.EXPECT().Unlock()
	speaker.EXPECT().Play(gomock.Any())
	// shut down the queue after first get
	controlQueue.EXPECT().Get().Return(nil, true)
	player.controlLoop()

	// test pause
	pauseOp := newTransportOperation(doPause, 0, nil)
	controlQueue.EXPECT().Get().Return(pauseOp, false)
	controlQueue.EXPECT().Done(pauseOp)
	speaker.EXPECT().Lock()
	speaker.EXPECT().Unlock()
	// shut down the queue after first get
	controlQueue.EXPECT().Get().Return(nil, true)
	player.controlLoop()

	// test stop
	stopOp := newTransportOperation(doStop, 0, nil)
	controlQueue.EXPECT().Get().Return(stopOp, false)
	controlQueue.EXPECT().Done(stopOp)
	speaker.EXPECT().Lock()
	streamer.EXPECT().Seek(0)
	speaker.EXPECT().Unlock()
	// shut down the queue after first get
	controlQueue.EXPECT().Get().Return(nil, true)
	player.controlLoop()

	// test close
	closeOp := newTransportOperation(doClose, 0, nil)
	controlQueue.EXPECT().Get().Return(closeOp, false)
	controlQueue.EXPECT().Done(closeOp)
	controlQueue.EXPECT().ShutDown()
	speaker.EXPECT().Lock()
	speaker.EXPECT().Unlock()
	streamer.EXPECT().Close()

	// shut down the queue after first get
	controlQueue.EXPECT().Get().Return(nil, true)
	player.controlLoop()

	// test queue shutdown
	controlQueue.EXPECT().Get().Return(nil, true)
	player.controlLoop()

}
