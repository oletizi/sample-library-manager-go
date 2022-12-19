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

package experimental

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"
)

func TestBeep(t *testing.T) {
	f, err := os.Open("../../test/data/library/multi-level/hh.wav")
	assert.Nil(t, err)

	streamer, format, err := wav.Decode(f)
	assert.Nil(t, err)
	defer func(streamer beep.StreamSeekCloser) {
		err := streamer.Close()
		if err != nil {
			log.Panic(err)
		}
	}(streamer)

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Println("Can't open speaker. Probably no sound card. Shouldn't be testing that in CI anyway.")
		return
	}

	done := make(chan bool)

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	select {
	case c := <-done:
		fmt.Println("Done: ", c)
	case <-time.After(1000 * time.Millisecond):
		//log.Panic("Timeout!")
		log.Printf("Timeout!")
	}
}
