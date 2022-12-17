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
	"github.com/rivo/tview"
	"log"
)

//go:generate mockgen -destination=../../mocks/tui/application.go . Application
type Application interface {
	Run() error
}

type tviewApp struct {
	app    *tview.Application
	logger *log.Logger
}

func (t *tviewApp) Run() error {
	return t.app.Run()
}

func NewTviewApplication() Application {
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
	return &tviewApp{
		app:    app,
		logger: log.New(logView, "", 0),
	}
}
