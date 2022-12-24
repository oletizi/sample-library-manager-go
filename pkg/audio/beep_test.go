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

func TestBeepPlayer_Close(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	controlQueue := mockutil.NewMockQueue(ctl)
	streamer := mockbp.NewMockStreamer(ctl)
	speaker := mockbp.NewMockSpeaker(ctl)
	player := &beepPlayer{
		controlQueue: controlQueue,
		spk:          speaker,
		streamer:     streamer,
	}

	// expect player.CLose() to call streamer.Close()
	// XXX: Can't figure out how to get an appropriate matcher.
	controlQueue.EXPECT().Add(gomock.Any())
	player.Close()
}

func TestBeepPlayer_Play(t *testing.T) {
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
	controlQueue.EXPECT().Add(gomock.Any())
	player.Play(&callback)
}

func TestTransportOperation(t *testing.T) {
	operation := transportOperation{op: doStop}
	assert.NotNil(t, operation.timestamp)
	fmt.Printf("Timestampe: %s\n", operation.timestamp)
	fmt.Printf("Timestampe: %s\n", operation.timestamp)
}
