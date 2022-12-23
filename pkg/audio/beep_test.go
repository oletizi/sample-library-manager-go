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
	"github.com/faiface/beep"
	"github.com/golang/mock/gomock"
	mockbp "github.com/oletizi/samplemgr/mocks/audio/bp"
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

	streamer := mockbp.NewMockStreamer(ctl)
	speaker := mockbp.NewMockSpeaker(ctl)
	player := &beepPlayer{
		spk:      speaker,
		streamer: streamer,
	}

	// expect player.CLose() to call streamer.Close()
	speaker.EXPECT().Lock()
	streamer.EXPECT().Close()
	speaker.EXPECT().Unlock()

	err := player.Close()
	assert.Nil(t, err)
}

func TestBeepPlayer_Play(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	streamer := mockbp.NewMockStreamer(ctl)
	speaker := mockbp.NewMockSpeaker(ctl)
	player := &beepPlayer{
		spk:               speaker,
		speakerSampleRate: 44100,
		streamer:          streamer,
		format:            &beep.Format{SampleRate: 44100, NumChannels: 2, Precision: 4},
	}

	speaker.EXPECT().Play(gomock.Any())
	err := player.Play(func() {})
	assert.Nil(t, err)
}

func TestBeepPlayer_Transport(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	speaker := mockbp.NewMockSpeaker(ctl)
	streamer := mockbp.NewMockStreamer(ctl)
	player := &beepPlayer{
		spk:               speaker,
		speakerSampleRate: 44100,
		streamer:          streamer,
		format:            &beep.Format{SampleRate: 44100, NumChannels: 2, Precision: 4},
	}
	// If Play() or Loop() haven't been called yet, pause should be a nop
	player.Pause()

	speaker.EXPECT().Play(gomock.Any())
	err := player.Play(func() {})
	assert.Nil(t, err)

	speaker.EXPECT().Lock()
	speaker.EXPECT().Unlock()
	player.Pause()

	speaker.EXPECT().Lock()
	streamer.EXPECT().Seek(0)
	speaker.EXPECT().Unlock()
	speaker.EXPECT().Play(gomock.Any())
	err = player.Loop(1, func() {})
	assert.Nil(t, err)

	speaker.EXPECT().Lock()
	streamer.EXPECT().Seek(0)
	speaker.EXPECT().Unlock()
	err = player.Stop()
	assert.Nil(t, err)
}
