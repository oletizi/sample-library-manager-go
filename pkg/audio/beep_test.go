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
	mockmock "github.com/oletizi/samplemgr/mocks/audio/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBeepContext_PlayerFor(t *testing.T) {
	ctx := &beepContext{}
	player, err := ctx.PlayerFor("../../test/data/library/multi-level/hh.wav")
	assert.Nil(t, err)
	assert.NotNil(t, player)
}

func TestBeepPlayer_Close(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	streamer := mockmock.NewMockFakeStreamer(ctl)
	player := &beepPlayer{
		streamer: streamer,
	}

	// expect player.CLose() to call streamer.Close()
	streamer.EXPECT().Close()
	err := player.Close()
	assert.Nil(t, err)
}

func TestBeepPlayer_Play(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	fakeSpeaker := mockmock.NewMockFakeSpeaker(ctl)
	streamer := mockmock.NewMockFakeStreamer(ctl)

	player := &beepPlayer{
		play:              fakeSpeaker.Play,
		speakerSampleRate: 44100,
		streamer:          streamer,
		format:            &beep.Format{SampleRate: 44100, NumChannels: 2, Precision: 4},
	}

	fakeSpeaker.EXPECT().Play(gomock.Any())
	player.Play(func() {})

}
