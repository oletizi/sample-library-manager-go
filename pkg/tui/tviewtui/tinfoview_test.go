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

package tviewtui

import (
	"github.com/golang/mock/gomock"
	mock_samplelib "github.com/oletizi/samplemgr/mocks/samplelib"
	mock_tui "github.com/oletizi/samplemgr/mocks/tui"
	mock_view "github.com/oletizi/samplemgr/mocks/tui/view"
	"github.com/oletizi/samplemgr/pkg/tui"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestTInfoView_constructor(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	app := tview.NewApplication()
	textView := tview.NewTextView()
	display := mock_view.NewMockDisplay(ctl)
	logger := log.Default()
	errorHandler := mock_tui.NewMockErrorHandler(ctl)

	iv := newTInfoView(app, textView, display, logger, errorHandler)

	assert.NotNil(t, iv)
	assert.Equal(t, app, iv.app)
	assert.Equal(t, textView, iv.textView)
	assert.Equal(t, display, iv.display)
	assert.Equal(t, logger, iv.logger)
	assert.Equal(t, errorHandler, iv.eh)
}

func TestTInfoView_UpdateMethods(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ds := mock_samplelib.NewMockDataSource(ctl)
	display := mock_view.NewMockDisplay(ctl)
	ti := &tInfoView{
		textView: tview.NewTextView(),
		eh:       tui.NewErrorHandler(log.Default()),
		display:  display,
	}
	ti.Update("Some text")

	node := mock_samplelib.NewMockNode(ctl)
	display.EXPECT().DisplayNodeAsText(ds, node)

	ti.UpdateNode(ds, node)

	sample := mock_samplelib.NewMockSample(ctl)
	display.EXPECT().DisplaySampleAsText(ds, sample)
	ti.UpdateSample(ds, sample)
}
