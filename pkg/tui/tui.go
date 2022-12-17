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

package tui

import (
	"github.com/oletizi/samplemgr/pkg/tui/view"
	"github.com/rivo/tview"
	"log"
)

//go:generate mockgen -destination=../../mocks/tui/application.go . Application
type Application interface {
	Run() error
}

//go:generate mockgen -destination=../../mocks/tui/userinterface.go . UserInterface
type UserInterface interface {
	NodeView() view.NodeView
	InfoView() view.InfoView
	LogView() view.LogView
}

type tviewUi struct {
	app      *tview.Application
	nodeView *tview.List
	infoView *tview.TextView
	logView  *tview.TextView
	logger   *log.Logger
}

func (t *tviewUi) Run() error {
	return t.app.Run()
}

func NewTviewInterface() Application {
	app := tview.NewApplication()

	nodeView := tview.NewList()
	nodeView.SetBorder(true).SetTitle(" Node ")

	infoView := tview.NewTextView()
	infoView.SetBorder(true).SetTitle(" Info ")

	logView := tview.NewTextView()
	logView.SetBorder(true).SetTitle(" Log ")

	flex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nodeView, 0, 1, true).
		AddItem(infoView, 0, 1, false).
		AddItem(logView, 0, 1, false)

	app.SetRoot(flex, true)
	return &tviewUi{
		app:      app,
		nodeView: nodeView,
		infoView: infoView,
		logView:  logView,
		logger:   log.New(logView, "", 0),
	}
}
