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

//go:generate mockgen -destination ../../mocks/audio/context.go . Context
type Context interface {
	PlayerFor(url string) (Player, error)
}

//go:generate mockgen -destination ../../mocks/audio/player.go . Player
type Player interface {
	Playing() bool                    // Playing returns true if the player is currently playing.
	Play(callBack *func())            // Play sound.
	Loop(times int, callback *func()) // Loop the sound `times` number of times. times = -1: loop forever.
	Pause()                           // Pause or resume playback from previous pause.
	Stop()                            // Stop playback and reset transport.
	Close()                           // Close the player and underlying resources.
}
