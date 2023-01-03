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
	"github.com/golang/mock/gomock"
	mock_controller "github.com/oletizi/samplemgr/mocks/tui/controller"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

func TestControlPanel_constructorAndInputCapture(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	logger := log.Default()
	app := tview.NewApplication()
	layout := tview.NewFlex()
	controller := mock_controller.NewMockController(ctl)
	c := newControlPanel(logger, app, layout, controller)
	c.ctl = controller

	assert.NotNil(t, c)
	assert.NotNil(t, c.controls)
	assert.NotNil(t, c.logger)
	assert.NotNil(t, c.textView)
	assert.Equal(t, app, c.app)

	assert.Equal(t, 1, layout.GetItemCount())

	// test F1 key event, "edit" in the main control context
	controller.EXPECT().EditStart()
	event := tcell.NewEventKey(tcell.KeyF1, ' ', tcell.ModNone)
	c.inputCapture(event)

	// test F1 key event, "save" in the edit control context
	c.ShowEditControls()
	controller.EXPECT().EditCommit()
	c.inputCapture(event)

	// test F2 key event, "cancel" in the edit control context
	event = tcell.NewEventKey(tcell.KeyF2, ' ', tcell.ModNone)
	c.ShowEditControls()
	controller.EXPECT().EditCancel()
	c.inputCapture(event)
}

func TestControlPanel_ShowEditControls(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	controller := mock_controller.NewMockController(ctl)

	textView := tview.NewTextView()
	cp := &controlPanel{
		textView: textView,
		ctl:      controller,
	}
	cp.ShowEditControls()
	text := textView.GetText(true)
	assert.True(t, strings.Contains(text, "Save"))
	assert.True(t, strings.Contains(text, "Cancel"))

	// test the controls
	assert.Equal(t, 2, len(cp.controls))

	// test save control action
	saveControl := cp.controls[0]
	cancelControl := cp.controls[1]

	controller.EXPECT().EditCommit()
	saveControl.Action()

	controller.EXPECT().EditCancel()
	cancelControl.Action()
}

func TestControlPanel_ShowMainControls(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	controller := mock_controller.NewMockController(ctl)
	textView := tview.NewTextView()
	cp := &controlPanel{
		textView: textView,
		ctl:      controller,
	}
	cp.ShowMainControls()
	text := textView.GetText(true)
	assert.True(t, strings.Contains(text, "Edit"))

	// test the controls
	assert.Equal(t, 1, len(cp.controls))

	// test the control action
	control := cp.controls[0]

	controller.EXPECT().EditStart()
	control.Action()

}
