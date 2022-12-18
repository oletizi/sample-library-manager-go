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
	mock_view "github.com/oletizi/samplemgr/mocks/tui/view"
	"github.com/oletizi/samplemgr/pkg/tui"
	"github.com/rivo/tview"
	"log"
	"testing"
)

func TestTInfoView_UpdateMethods(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	display := mock_view.NewMockDisplay(ctl)
	ti := &tInfoView{
		textView: tview.NewTextView(),
		eh:       tui.NewErrorHandler(log.Default()),
		display:  display,
	}
	ti.Update("Some text")

	node := mock_samplelib.NewMockNode(ctl)
	display.EXPECT().DisplayNodeAsText(node)

	ti.UpdateNode(node)

	sample := mock_samplelib.NewMockSample(ctl)
	display.EXPECT().DisplaySampleAsListing(sample)
	ti.UpdateSample(sample)
}
