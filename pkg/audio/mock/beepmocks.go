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

package mock

import "github.com/faiface/beep"

// FakeStreamer is a proxy for beep.StreamSeekCloser just for mocking. It turns out you can't declare it
// in a *_test.go file, so it's here
//
//go:generate mockgen -destination ../../../mocks/audio/mock/fakestreamer.go . FakeStreamer
type FakeStreamer interface {
	beep.StreamSeekCloser
}

// FakeSpeaker is a proxy for functions in the beep "speaker" package just for mocking.
//
//go:generate mockgen -destination ../../../mocks/audio/mock/fakespeaker.go . FakeSpeaker
type FakeSpeaker interface {
	Play(s ...beep.Streamer)
}
