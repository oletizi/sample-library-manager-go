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
	"github.com/golang/mock/gomock"
	mock_controller "github.com/oletizi/samplemgr/mocks/tui/controller"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

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
