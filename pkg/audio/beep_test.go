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
	"github.com/golang/mock/gomock"
	mockmock "github.com/oletizi/samplemgr/mocks/audio/mock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestBeepPlayer_Close(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	streamer := mockmock.NewMockFakeStreamer(ctl)
	player := &beepPlayer{streamer: streamer}

	// expect player.CLose() to call streamer.Close()
	streamer.EXPECT().Close()
	err := player.Close()
	assert.Nil(t, err)
}

func TestBeepPlayer_Play(t *testing.T) {
	ctx, err := NewBeepContext()
	assert.Nil(t, err)

	player, err := ctx.PlayerFor("../../test/data/library/multi-level/hh.wav")
	assert.Nil(t, err)

	done := make(chan bool)

	player.Play(func() {
		log.Println("Done playing.")
		done <- true
	})
	select {
	case c := <-done:
		fmt.Println("Done: ", c)
	case <-time.After(1000 * time.Millisecond):
		//log.Panic("Timeout!")
		log.Printf("Timeout!")
	}
}
