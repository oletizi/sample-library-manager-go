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
	"github.com/rivo/tview"
)

type tNodeView struct {
	list    *tview.List
	display view.Display
	eh      tui.ErrorHandler
	logger  tui.Logger
}

func (t *tNodeView) UpdateNode(
	ds samplelib.DataSource,
	node samplelib.Node,
	nodeSelected func(node samplelib.Node),
	sampleSelected func(sample samplelib.Sample),
	nodeChosen func(node samplelib.Node),
	sampleChosen func(sample samplelib.Sample),
) {
	t.list.SetTitle(" " + node.Name() + " ")
	t.list.Clear()
	// if the node has a parent, add an item for it.
	if !node.Parent().Null() {
		text := ".."
		parent := node.Parent()
		t.list.AddItem(text, "", 0, func() {
			t.logger.Print("Parent node chosen: " + parent.Name())
			nodeChosen(parent)
		})
	}

	// Get the children of the new node
	children, err := ds.ChildrenOf(node)
	t.eh.Handle(err)

	for _, child := range children {
		text := t.display.DisplayNodeAsListing(child, false)
		thisChild := child
		t.list.AddItem(text, "", 0, func() {
			t.logger.Print("Child node chosen: " + thisChild.Name())
			nodeChosen(thisChild)
		})
	}

}
