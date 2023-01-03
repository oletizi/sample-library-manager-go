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
	"github.com/oletizi/samplemgr/pkg/audio"
	"github.com/oletizi/samplemgr/pkg/samplelib"
	"github.com/oletizi/samplemgr/pkg/tui"
	"github.com/oletizi/samplemgr/pkg/tui/controller"
	"github.com/oletizi/samplemgr/pkg/tui/view"
	"github.com/oletizi/samplemgr/pkg/util"
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

// New creates a new instance of the tview-based TUI implementation
func New(ds samplelib.DataSource) (tui.Application, error) {

	app := tview.NewApplication()
	logView := &tLogView{textView: tview.NewTextView().SetScrollable(false)}
	logView.textView.SetBorder(true)

	l := log.New(logView, "", 0)
	logger := util.NewLogger(l)
	errorHandler := tui.NewErrorHandler(logger)

	display, err := view.NewDisplay(logger, errorHandler)
	if err != nil {
		return nil, err
	}

	nodeView := newTNodeView(app, tview.NewList(), display, logger, errorHandler)
	infoView := newTInfoView(app, tview.NewTextView(), display, logger, errorHandler)

	// Set up the display interface
	horizontalLayout := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nodeView.list, 0, 1, true).
		AddItem(infoView.textView, 0, 1, false).
		AddItem(logView.textView, 0, 1, false)

	// Set up the control panel view
	controlPanelLayout := tview.NewFlex().SetDirection(tview.FlexColumn)
	controlPanelLayout.SetBorder(true)

	// The vertical layout holds everything
	verticalLayout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(horizontalLayout, 0, 1, false).
		AddItem(controlPanelLayout, 3, 0, false)

	app.SetRoot(verticalLayout, true)
	rootNode, err := ds.RootNode()
	errorHandler.Handle(err)

	audioContext, err := audio.NewBeepContext(logger)
	if err != nil {
		log.Fatal(err)
	}

	// XXX: Not sure that it's worth the isolation of a constructor when there's this mutual dependency
	// At least the mess is contained in a single function.
	//controlPanel := newControlPanel(logger, app, controlPanelLayout)
	controlPanelConstructor := func(ctl controller.Controller) view.ControlPanel {
		controlPanel := newControlPanel(logger, app, controlPanelLayout, ctl)
		return &controlPanel
	}
	ctl := controller.New(audioContext, ds, errorHandler, nodeView, infoView, logView, controlPanelConstructor)
	ctl.UpdateNode(rootNode)
	ctl.StartPlayLoop()

	return &tApp{
		app:    app,
		logger: log.New(logView, "", 0),
	}, nil
}
