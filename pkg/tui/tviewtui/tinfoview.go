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
	"github.com/oletizi/samplemgr/pkg/samplelib"
	"github.com/oletizi/samplemgr/pkg/tui"
	"github.com/oletizi/samplemgr/pkg/tui/view"
	"github.com/oletizi/samplemgr/pkg/util"
	"github.com/rivo/tview"
)

type tInfoView struct {
	textView *tview.TextView
	logger   util.Logger
	eh       tui.ErrorHandler
	display  view.Display
}

func (t *tInfoView) Update(v string) {
	t.textView.Clear()
	_, err := t.textView.Write([]byte(v))
	t.eh.Handle(err)
}

func (t *tInfoView) UpdateNode(ds samplelib.DataSource, node samplelib.Node) {
	t.Update(t.display.DisplayNodeAsText(ds, node))
}

func (t *tInfoView) UpdateSample(ds samplelib.DataSource, sample samplelib.Sample) {
	t.Update(t.display.DisplaySampleAsText(ds, sample))
}
