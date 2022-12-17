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
	nv     view.NodeView
	iv     view.InfoView
	lv     view.LogView
	eh     tui.ErrorHandler
	logger *log.Logger
}

func (c *controller) UpdateNode(node samplelib.Node) {
	c.nv.UpdateNode(node, c.nodeSelected, c.sampleSelected, c.nodeChosen, c.sampleChosen)
}

func (c *controller) nodeSelected(node samplelib.Node) {
	c.iv.UpdateNode(node)
}

func (c *controller) sampleSelected(sample samplelib.Sample) {
	c.iv.UpdateSample(sample)
}

func (c *controller) nodeChosen(node samplelib.Node) {
	c.UpdateNode(node)
}

func (c *controller) sampleChosen(sample samplelib.Sample) {
	// for now, choosing a sample is the same as selecting a sample
	c.sampleSelected(sample)
}

func New(ds samplelib.DataSource, nodeView view.NodeView, infoView view.InfoView, logView view.LogView) Controller {
	return &controller{ds: ds,
		nv:     nodeView,
		iv:     infoView,
		lv:     logView,
		logger: log.New(logView, "", 0)}
}
