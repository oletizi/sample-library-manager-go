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

// Constructor for tNodeView. Discourages forgetting to set properties. Wires up listeners.
// Also sets some display defaults
func newTNodeView(
	list *tview.List,
	display view.Display,
	logger tui.Logger,
	handler tui.ErrorHandler,
) *tNodeView {
	list.ShowSecondaryText(false)
	list.SetBorder(true)
	return &tNodeView{
		list:    list,
		display: display,
		logger:  logger,
		eh:      handler,
	}
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

	var nodes []samplelib.Node
	// if the node has a parent, add an item for it.
	if !node.Parent().Null() {
		text := ".."
		parent := node.Parent()
		nodes = append(nodes, parent)
		t.list.AddItem(text, "", 0, func() {
			t.logger.Print("Parent node chosen: " + parent.Name())
			nodeChosen(parent)
		})
	}

	// Get the children of the new node
	children, err := ds.ChildrenOf(node)
	t.eh.Handle(err)

	nodes = append(nodes, children...)
	for _, child := range children {
		text := t.display.DisplayNodeAsListing(child, false)
		thisChild := child
		t.list.AddItem(text, "", 0, func() {
			t.logger.Print("Child node chosen: " + thisChild.Name())
			nodeChosen(thisChild)
		})
	}

	// Get the samples of the new node
	samples, err := ds.SamplesOf(node)
	t.eh.Handle(err)
	for _, sample := range samples {
		text := t.display.DisplaySampleAsListing(sample)
		thisSample := sample
		t.list.AddItem(text, "", 0, func() {
			t.logger.Print("Sample chosen: " + thisSample.Name())
			sampleSelected(thisSample)
		})
	}

	// set the callback function for when a new list element is selected (e.g., w/ arrow keys)
	t.list.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		t.logger.Printf("Node view changed: index: %d", index)
		if index < len(nodes) {
			nodeSelected(nodes[index])
		} else {
			sampleSelected(samples[index-len(nodes)])
		}
	})
}
