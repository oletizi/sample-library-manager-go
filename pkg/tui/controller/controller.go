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
	"log"
)

//go:generate mockgen -destination=../../../mocks/tui/controller/controller.go . Controller
type Controller interface {
	UpdateNode(node samplelib.Node)
}

type controller struct {
	ds     samplelib.DataSource
	ui     tui.UserInterface
	eh     tui.ErrorHandler
	logger *log.Logger
}

func (c *controller) UpdateNode(node samplelib.Node) {
	children, err := c.ds.ChildrenOf(node)
	if err != nil {
		c.eh.Print(err)
	}
	for _, child := range children {
		c.logger.Printf("Child! %T", child)
	}
}

func New(ds samplelib.DataSource, ui tui.UserInterface) Controller {
	return &controller{ds: ds, ui: ui, logger: log.New(ui.LogView(), "", 0)}
}
