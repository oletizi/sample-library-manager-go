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
	"github.com/oletizi/samplemgr/pkg/tui/controller"
	"github.com/oletizi/samplemgr/pkg/tui/view"
	"github.com/rivo/tview"
	"log"
)

type tApp struct {
	app    *tview.Application
	logger *log.Logger
}

// notest
func (t *tApp) Run() error {
	return t.app.Run()
}

func New(ds samplelib.DataSource) (tui.Application, error) {

	app := tview.NewApplication()
	logView := &tLogView{textView: tview.NewTextView()}
	logView.textView.SetBorder(true)

	l := log.New(logView, "", 0)
	logger := tui.NewLogger(l)
	errorHandler := tui.NewErrorHandler(logger)

	display, err := view.NewDisplay(logger, errorHandler)
	if err != nil {
		return nil, err
	}

	nodeView := newTNodeView(tview.NewList(), display, logger, errorHandler)
	infoView := newTInfoView(tview.NewTextView(), display, logger, errorHandler)

	flex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nodeView.list, 0, 1, true).
		AddItem(infoView.textView, 0, 1, false).
		AddItem(logView.textView, 0, 1, false)

	app.SetRoot(flex, true)

	rootNode, err := ds.RootNode()
	errorHandler.Handle(err)

	ctl := controller.New(ds, errorHandler, nodeView, infoView, logView)
	ctl.UpdateNode(rootNode)

	return &tApp{
		app:    app,
		logger: log.New(logView, "", 0),
	}, nil
}

func newTInfoView(
	textView *tview.TextView,
	display view.Display,
	logger tui.Logger,
	handler tui.ErrorHandler,
) *tInfoView {
	textView.SetBorder(true)
	textView.SetTitle(" Info ")
	return &tInfoView{
		textView: textView,
		display:  display,
		logger:   logger,
		eh:       handler,
	}
}
