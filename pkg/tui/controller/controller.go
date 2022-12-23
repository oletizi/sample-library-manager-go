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
	"k8s.io/client-go/util/workqueue"
	"log"
)

//go:generate mockgen -destination=../../../mocks/tui/controller/controller.go . Controller
type Controller interface {
	UpdateNode(node samplelib.Node)
	StartPlayLoop()
}

type controller struct {
	playQueue     workqueue.Interface
	ac            audio.Context
	ds            samplelib.DataSource
	eh            tui.ErrorHandler
	nv            view.NodeView
	iv            view.InfoView
	lv            view.LogView
	logger        tui.Logger
	currentPlayer audio.Player
	currentSample samplelib.Sample
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
			err = c.currentPlayer.Stop()
			c.logger.Println("Current player stopped.")
			c.eh.Handle(err)
			// If the current sample is the same as the newSample, don't play the sample again.
			// This is the play/pause toggle condition.
			if newSample.Equal(c.currentSample) {
				c.logger.Println("Was already playing chosen sample. Break.")
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
		err = newPlayer.Play(func() {
			// notest
			c.logger.Println("Done playing newSample! Closing the newPlayer...")
			err := newPlayer.Close()
			c.eh.Handle(err)
		})
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

// UpdateNode tells the controller to update the UI for a new node
func (c *controller) UpdateNode(node samplelib.Node) {
	c.logger.Print("Calling UpdateNode on node: " + node.Name())
	c.nv.UpdateNode(c.ds, node, c.nodeSelected, c.sampleSelected, c.nodeChosen, c.sampleChosen)
	c.iv.UpdateNode(c.ds, node)
}

func (c *controller) StartPlayLoop() {
	// notest
	go c.playLoop()
}

// nodeSelected callback function for when a node is selected in the node view
func (c *controller) nodeSelected(node samplelib.Node) {
	c.iv.UpdateNode(c.ds, node)
}

// sampleSelected callback function for when a sample is selected in the node view
func (c *controller) sampleSelected(sample samplelib.Sample) {
	c.iv.UpdateSample(c.ds, sample)
}

// nodeChosen callback function for when a node is chosen in the node view
func (c *controller) nodeChosen(node samplelib.Node) {
	c.logger.Print("In controller nodeChosen: " + node.Name())
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
	rv := &controller{
		playQueue: workqueue.New(),
		ac:        ac,
		ds:        ds,
		eh:        eh,
		nv:        nodeView,
		iv:        infoView,
		lv:        logView,
		logger:    log.New(logView, "", 0)}
	return rv
}
