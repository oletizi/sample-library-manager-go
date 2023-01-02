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

package controller

import (
	"github.com/oletizi/samplemgr/pkg/audio"
	"github.com/oletizi/samplemgr/pkg/samplelib"
	"github.com/oletizi/samplemgr/pkg/tui"
	"github.com/oletizi/samplemgr/pkg/tui/view"
	"github.com/oletizi/samplemgr/pkg/util"
	"k8s.io/client-go/util/workqueue"
	"log"
)

type editContext struct {
	start  func()
	commit func() error
	cancel func()
}

func newNullEditContext(logger util.Logger) editContext {
	return editContext{
		func() {
			logger.Println("Null start edit.")
		},
		func() error {
			logger.Println("Null commit edit.")
			return nil
		},
		func() {
			logger.Println("Null cancel edit.")
		},
	}
}

//go:generate mockgen -destination=../../../mocks/tui/controller/controller.go . Controller
type Controller interface {
	UpdateNode(node samplelib.Node)
	SetControlPanel(cp view.ControlPanel)
	StartPlayLoop()
	EditStart()
	EditCommit()
	EditCancel()
}

type controller struct {
	playQueue     workqueue.Interface
	ac            audio.Context
	ds            samplelib.DataSource
	eh            tui.ErrorHandler
	nv            view.NodeView
	iv            view.InfoView
	lv            view.LogView
	logger        util.Logger
	controlPanel  view.ControlPanel
	currentPlayer audio.Player
	currentSample samplelib.Sample
	editContext   editContext
}

// UpdateNode tells the controller to update the UI for a new node
func (c *controller) UpdateNode(node samplelib.Node) {
	c.nv.UpdateNode(c.ds, node, c.nodeSelected, c.sampleSelected, c.nodeChosen, c.sampleChosen)
	c.iv.UpdateNode(c.ds, node)
}

func (c *controller) SetControlPanel(controlPanel view.ControlPanel) {
	c.controlPanel = controlPanel
}

func (c *controller) StartPlayLoop() {
	// notest
	go c.playLoop()
}

func (c *controller) EditStart() {
	c.editContext.start()
}

func (c *controller) EditCommit() {
	c.eh.Handle(c.editContext.commit())
}

func (c *controller) EditCancel() {
	c.editContext.cancel()
}

func (c *controller) playLoop() {
	for {
		c.logger.Println("Reading from play queue...")
		item, shutdown := c.playQueue.Get()
		if shutdown {
			c.logger.Println("Queue shutdown. Return.")
			return
		}
		c.logger.Println("Done reading from play queue.")
		c.playQueue.Done(item)
		newSample := item.(samplelib.Sample) // not sure why workqueue interface isn't generic

		var err error

		// stop current playback, if any
		if c.currentPlayer != nil && c.currentPlayer.Playing() {
			c.logger.Println("Current player is playing, stopping!")
			c.currentPlayer.Stop()
			c.logger.Println("Current player stopped.")
			c.eh.Handle(err)
			// If the current sample is the same as the newSample, don't play the sample again.
			// This is the play/pause toggle condition.
			if newSample.Equal(c.currentSample) {
				c.logger.Println("Was already playing chosen sample. Continue.")
				continue
			}
		}

		// if the chosen newSample is different than the current newSample, create a new newPlayer
		// and start playback
		c.logger.Printf("Creating a new player for %s", newSample.Path())
		newPlayer, err := c.ac.PlayerFor(newSample.Path())
		if err != nil {
			// notest
			c.eh.Handle(err)
			continue
		}
		// Play the chosen newSample
		c.logger.Println("Calling play...")
		callback := func() {
			// notest
			c.logger.Println("Done playing newSample! Closing the newPlayer...")
			newPlayer.Close()
			c.eh.Handle(err)
		}
		newPlayer.Play(&callback)
		c.logger.Println("Done calling play.")
		if err != nil {
			// notest
			c.eh.Handle(err)
			continue
		}
		c.logger.Println("Setting current sample and player")
		c.currentSample = newSample
		c.currentPlayer = newPlayer
		c.logger.Println("End of play loop; back to the beginning...")
	}
}

// nodeSelected callback function for when a node is selected in the node view
func (c *controller) nodeSelected(node samplelib.Node) {
	// set the edit context
	c.editContext = editContext{
		start: func() {
			c.logger.Printf("Start edit on node: %v", node.Name())
		},
		commit: func() error {
			c.logger.Printf("Commit edit on node: %v", node.Name())
			return nil
		},
		cancel: func() {
			c.logger.Println("Cancel edit on node: %v", node.Name())
		},
	}
	// update the info view
	c.iv.UpdateNode(c.ds, node)
}

// sampleSelected callback function for when a sample is selected in the node view
func (c *controller) sampleSelected(sample samplelib.Sample) {
	// set the edit context
	c.editContext = editContext{
		start: func() {
			c.logger.Printf("Start edit on sample: %v", sample.Name())
		},
		commit: func() error {
			c.logger.Printf("Commit edit on sample: %v", sample.Name())
			return nil
		},
		cancel: func() {
			c.logger.Printf("Cancel edit on sample: %v", sample.Name())
		},
	}
	// update the info view
	c.iv.UpdateSample(c.ds, sample)
}

// nodeChosen callback function for when a node is chosen in the node view
func (c *controller) nodeChosen(node samplelib.Node) {
	c.UpdateNode(node)
}

// sampleChosen callback function for when a sample is chosen in the node view
func (c *controller) sampleChosen(newSample samplelib.Sample) {
	c.playQueue.Add(newSample)
}

func New(
	ac audio.Context,
	ds samplelib.DataSource,
	eh tui.ErrorHandler,
	nodeView view.NodeView,
	infoView view.InfoView,
	logView view.LogView,
) Controller {
	logger := log.New(logView, "", 0)
	rv := &controller{
		playQueue:   workqueue.New(),
		ac:          ac,
		ds:          ds,
		eh:          eh,
		nv:          nodeView,
		iv:          infoView,
		lv:          logView,
		logger:      logger,
		editContext: newNullEditContext(logger),
	}
	nodeView.Focus()
	return rv
}
