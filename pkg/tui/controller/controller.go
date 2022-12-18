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
	"github.com/oletizi/samplemgr/pkg/samplelib"
	"github.com/oletizi/samplemgr/pkg/tui"
	"github.com/oletizi/samplemgr/pkg/tui/view"
	"log"
)

//go:generate mockgen -destination=../../../mocks/tui/controller/controller.go . Controller
type Controller interface {
	UpdateNode(node samplelib.Node)
}

type controller struct {
	ds     samplelib.DataSource
	eh     tui.ErrorHandler
	nv     view.NodeView
	iv     view.InfoView
	lv     view.LogView
	logger *log.Logger
}

// UpdateNode tells the controller to update the UI for a new node
func (c *controller) UpdateNode(node samplelib.Node) {
	c.logger.Print("Calling UpdateNode on node: " + node.Name())
	c.nv.UpdateNode(c.ds, node, c.nodeSelected, c.sampleSelected, c.nodeChosen, c.sampleChosen)
	c.iv.UpdateNode(node)
}

// nodeSelected callback function for when a node is selected in the node view
func (c *controller) nodeSelected(node samplelib.Node) {
	c.iv.UpdateNode(node)
}

// sampleSelected callback function for when a sample is selected in the node view
func (c *controller) sampleSelected(sample samplelib.Sample) {
	c.iv.UpdateSample(sample)
}

// nodeChosen callback function for when a node is chosen in the node view
func (c *controller) nodeChosen(node samplelib.Node) {
	c.logger.Print("In controller nodeChosen: " + node.Name())
	c.UpdateNode(node)
}

// sampleChosen callback function for when a sample is chosen in the node view
func (c *controller) sampleChosen(sample samplelib.Sample) {
	// for now, choosing a sample is the same as selecting a sample
	c.sampleSelected(sample)
}

func New(ds samplelib.DataSource, eh tui.ErrorHandler, nodeView view.NodeView, infoView view.InfoView, logView view.LogView) Controller {
	return &controller{
		ds:     ds,
		eh:     eh,
		nv:     nodeView,
		iv:     infoView,
		lv:     logView,
		logger: log.New(logView, "", 0)}
}
