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
	"github.com/oletizi/samplemgr/pkg/audio/bp"
	"github.com/oletizi/samplemgr/pkg/util"
	"k8s.io/client-go/util/workqueue"
	"log"
	"os"
	"time"
)

type operation string

const (
	doLoop  = "doLoop"
	doPause = "doPause"
	doStop  = "doStop"
	doClose = "doClose"
)

type transportOperation struct {
	timestamp time.Time
	op        operation
	times     int
	callback  *func()
}

func newTransportOperation(op operation, times int, callback *func()) transportOperation {
	return transportOperation{
		op:        op,
		times:     times,
		callback:  callback,
		timestamp: time.Now(),
	}
}

type beepPlayer struct {
	logger            util.Logger
	controlQueue      workqueue.Interface
	spk               bp.Speaker
	speakerSampleRate beep.SampleRate
	streamer          beep.StreamSeekCloser
	format            *beep.Format
	ctl               *beep.Ctrl
}

func (b *beepPlayer) Playing() bool {
	return b.ctl != nil && !b.ctl.Paused
}

func (b *beepPlayer) Play(completedCallback *func()) {
	b.Loop(1, completedCallback)
}

func (b *beepPlayer) Loop(times int, completedCallback *func()) {
	b.controlQueue.Add(newTransportOperation(doLoop, times, completedCallback))
}

func (b *beepPlayer) Stop() {
	b.controlQueue.Add(newTransportOperation(doStop, 0, nil))
}

func (b *beepPlayer) Pause() {
	b.controlQueue.Add(newTransportOperation(doPause, 0, nil))
}

func (b *beepPlayer) Close() {
	b.controlQueue.Add(newTransportOperation(doClose, 0, nil))
}

func (b *beepPlayer) controlLoop() {
	for {
		item, shutdown := b.controlQueue.Get()
		if shutdown {
			return
		}
		b.controlQueue.Done(item)
		var op transportOperation
		op = item.(transportOperation)

		switch op.op {
		case doLoop:
			b.stop()
			resampled := beep.Resample(4, b.format.SampleRate, b.speakerSampleRate, beep.Loop(op.times, b.streamer))
			b.ctl = &beep.Ctrl{Streamer: beep.Seq(resampled, beep.Callback(*op.callback)), Paused: false}
			b.spk.Play(b.ctl)
		case doPause:
			if b.ctl != nil {
				// lock the speaker while we fiddle with the control...
				b.spk.Lock()
				b.ctl.Paused = !b.ctl.Paused
				// unlock the speaker
				b.spk.Unlock()
			}
		case doStop:
			b.stop()
		case doClose:
			b.controlQueue.ShutDown()
			// lock the speaker while we tear down resources...
			b.spk.Lock()
			err := b.streamer.Close()
			// unlock the speaker.
			b.spk.Unlock()
			if err != nil {
				// notest
				log.Print(err)
			}
		}
	}
}

// stop is not thread safe. Should only be called inside the control loop.
func (b *beepPlayer) stop() {
	if b.ctl == nil {
		// notest
		return
	}
	// lock the speaker while we fiddle with the control & streamer
	b.spk.Lock()
	b.ctl.Paused = true
	err := b.streamer.Seek(0)
	b.ctl = nil

	// unlock the speaker so it can keep doing stuff
	b.spk.Unlock()

	if err != nil {
		// notest
		log.Panic(err)
	}
}

type beepContext struct {
	speakerSampleRate beep.SampleRate
	logger            util.Logger
}

func NewBeepContext(logger util.Logger) (Context, error) {
	// notest
	sampleRate := beep.SampleRate(44100)
	// XXX: This should probably be done somewhere else... it doesn't need to be initialized every time.
	err := speaker.Init(sampleRate, sampleRate.N(time.Second/10))
	if err != nil {
		return nil, err
	}
	return &beepContext{speakerSampleRate: sampleRate, logger: logger}, err
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

	rv := &beepPlayer{
		controlQueue:      workqueue.New(),
		logger:            c.logger,
		spk:               bp.NewSpeaker(),
		speakerSampleRate: c.speakerSampleRate,
		streamer:          streamer,
		format:            &format,
	}
	// start the control loop
	go rv.controlLoop()
	return rv, nil
}
