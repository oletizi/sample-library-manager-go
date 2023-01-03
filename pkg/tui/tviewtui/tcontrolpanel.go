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
	"github.com/oletizi/samplemgr/pkg/util"
	"github.com/rivo/tview"
	"strings"
)

// displayMargin the spaces between each control
const displayMargin = "  "

type controlPanel struct {
	logger   util.Logger
	ctl      controller.Controller
	controls []view.Control
	textView *tview.TextView
	layout   *tview.Flex
	app      *tview.Application
}

func (c *controlPanel) ShowEditControls() {
	c.controls = []view.Control{
		{
			Label: "[::r]F1:[::-] Save",
			Keys:  []string{"F1"},
			Action: func() {
				c.ctl.EditCommit()
				c.ShowMainControls()
			},
		},
		{
			Label: "[::r]F2:[::-] Cancel",
			Keys:  []string{"F2", "Esc"},
			Action: func() {
				c.ctl.EditCancel()
				c.ShowMainControls()
			},
		},
	}
	c.update()
}

func (c *controlPanel) ShowMainControls() {
	c.controls = []view.Control{
		{
			Label: "[::r]F1:[::-] Edit",
			Keys:  []string{"F1"},
			Action: func() {
				c.ctl.EditStart()
				c.ShowEditControls()
			},
		},
	}
	c.update()
}

// update updates the control panel view with the current controls
func (c *controlPanel) update() {
	c.textView.Clear()
	var frags []string
	for _, control := range c.controls {
		frags = append(frags, control.Label)
	}
	c.textView.SetText(strings.Join(frags, displayMargin))
}

func (c *controlPanel) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	for _, control := range c.controls {
		for _, key := range control.Keys {
			if key == event.Name() {
				control.Action()
				break
			}
		}
	}
	return event
}

func newControlPanel(logger util.Logger, app *tview.Application, layout *tview.Flex, ctl controller.Controller) controlPanel {
	textView := tview.NewTextView()
	textView.SetDynamicColors(true)
	layout.AddItem(textView, 0, 1, false)
	cp := controlPanel{
		logger:   logger,
		app:      app,
		layout:   layout,
		textView: textView,
		ctl:      ctl,
	}
	cp.ShowMainControls()
	app.SetInputCapture(cp.inputCapture)
	return cp
}
