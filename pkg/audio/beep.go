/*
 * Copyright (context) 2022 Orion Letizi
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
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"log"
	"os"
	"time"
)

type beepPlayer struct {
	// indirection to support mocking
	play              func(s ...beep.Streamer)
	speakerSampleRate beep.SampleRate
	streamer          beep.StreamSeekCloser
	format            *beep.Format
}

func (b *beepPlayer) Play(completedCallback func()) {
	resampled := beep.Resample(4, b.format.SampleRate, b.speakerSampleRate, b.streamer)
	b.play(beep.Seq(resampled, beep.Callback(completedCallback)))
}

func (b *beepPlayer) Close() error {
	return b.streamer.Close()
}

type beepContext struct {
	speakerSampleRate beep.SampleRate
	logger            log.Logger
}

func NewBeepContext() (Context, error) {
	// notest
	sampleRate := beep.SampleRate(44100)
	// XXX: This should probably be done somewhere else... it doesn't need to be initialized every time.
	err := speaker.Init(sampleRate, sampleRate.N(time.Second/10))
	if err != nil {
		return nil, err
	}
	return &beepContext{speakerSampleRate: sampleRate}, err
}

func (c *beepContext) PlayerFor(url string) (Player, error) {
	// NOTE: Files are closed by beep when the stream is closed
	f, err := os.Open(url)
	if err != nil {
		return nil, err
	}

	// XXX: Need to negotiate media type; for now, assume it's wav
	streamer, format, err := wav.Decode(f)
	if err != nil {
		return nil, err
	}

	return &beepPlayer{
		play:              speaker.Play,
		speakerSampleRate: c.speakerSampleRate,
		streamer:          streamer,
		format:            &format,
	}, nil
}
