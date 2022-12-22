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

package bp

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

// Streamer is a proxy for beep.StreamSeekCloser to make it mockable
//
//go:generate mockgen -destination ../../../mocks/audio/bp/streamer.go . Streamer
type Streamer interface {
	beep.StreamSeekCloser
}

//go:generate mockgen -destination ../../../mocks/audio/bp/speaker.go . Speaker
type Speaker interface {
	Play(s ...beep.Streamer)
	Lock()
	Unlock()
}

func NewSpeaker() Speaker {
	// notest
	return &spk{}
}

type spk struct {
}

func (sp *spk) Play(streamers ...beep.Streamer) {
	// notest
	speaker.Play(streamers...)
}

func (sp *spk) Lock() {
	// notest
	speaker.Lock()
}

func (sp *spk) Unlock() {
	// notest
	speaker.Unlock()
}
