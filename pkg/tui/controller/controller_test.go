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

package controller

import (
	"bytes"
	"errors"
	"github.com/golang/mock/gomock"
	mock_samplelib "github.com/oletizi/samplemgr/mocks/samplelib"
	mock_tui "github.com/oletizi/samplemgr/mocks/tui"
	mock_view "github.com/oletizi/samplemgr/mocks/tui/view"
	"github.com/oletizi/samplemgr/pkg/samplelib"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	ctl := gomock.NewController(t)
	ds := mock_samplelib.NewMockDataSource(ctl)
	ui := mock_tui.NewMockUserInterface(ctl)
	logView := mock_view.NewMockLogView(ctl)

	ui.EXPECT().LogView().Times(1).Return(logView)
	c := New(ds, ui)
	assert.NotNil(t, c)
}

func TestController_UpdateNode(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	ds := mock_samplelib.NewMockDataSource(ctl)
	ui := mock_tui.NewMockUserInterface(ctl)
	eh := mock_tui.NewMockErrorHandler(ctl)
	node := mock_samplelib.NewMockNode(ctl)
	child := mock_samplelib.NewMockNode(ctl)

	children := make([]samplelib.Node, 0)
	children = append(children, child)
	ds.EXPECT().ChildrenOf(node).Times(1).Return(children, nil)

	// make a new controller
	c := &controller{
		ds:     ds,
		ui:     ui,
		eh:     eh,
		logger: log.Default(),
	}
	assert.NotNil(t, c)
	c.UpdateNode(node)

	// test error condition
	errString := "explosions"
	var buf bytes.Buffer
	c.logger = log.New(&buf, "", 0)
	// return an error from ChildrenOf()
	ds.EXPECT().ChildrenOf(node).Times(1).Return(nil, errors.New(errString))
	// expect controller to send error to handler
	eh.EXPECT().Print(gomock.Any()).Times(1)

	c.UpdateNode(node)
}
