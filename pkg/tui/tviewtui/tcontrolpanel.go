/*
 * Copyright (c) 2023 Orion Letizi
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

package tviewtui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/oletizi/samplemgr/pkg/tui/controller"
	"github.com/oletizi/samplemgr/pkg/tui/view"
	"github.com/rivo/tview"
)

// displayMargin the number of spaces between each control
const displayMargin = 2

type controlPanel struct {
	ctl      controller.Controller
	controls []view.Control
	layout   *tview.Flex
	app      *tview.Application
}

func (c *controlPanel) EditControls() {
	c.controls = []view.Control{
		{
			Label: "[::r]F1:[::-] Save",
			Key:   "F1",
			Action: func() {
				c.ctl.EditCommit()
				c.MainControls()
			},
		},
		{
			Label: "[::r]F2:[::-] Cancel",
			Key:   "F2",
			Action: func() {
				c.ctl.EditCancel()
				c.MainControls()
			},
		},
	}
	c.update()
}

func (c *controlPanel) MainControls() {
	c.controls = []view.Control{
		{
			Label: "[::r]F1:[::-] Edit",
			Key:   "F1",
			Action: func() {
				c.ctl.EditStart()
				c.EditControls()
			},
		},
	}
	c.update()
}

func (c *controlPanel) update() {
	c.layout.Clear()
	for _, control := range c.controls {
		v := tview.NewTextView()
		v.SetDynamicColors(true)
		v.SetText(control.Label)
		c.layout.AddItem(v, len(v.GetText(true))+displayMargin, 0, false)
	}
	// TODO: Figure out how to get the control panel flex layout to update
}

func (c *controlPanel) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	for _, control := range c.controls {
		if control.Key == event.Name() {
			control.Action()
			break
		}
	}
	return event
}

func newControlPanel(app *tview.Application, layout *tview.Flex, ctl controller.Controller) view.ControlPanel {
	cp := controlPanel{
		app:    app,
		layout: layout,
		ctl:    ctl,
	}
	cp.MainControls()
	app.SetInputCapture(cp.inputCapture)
	return &cp
}
