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

package view

import (
	"github.com/oletizi/samplemgr/pkg/samplelib"
	"io"
)

type View interface {
	Focus()
}

type TextView interface {
	View
	Update(v string)
}

//go:generate mockgen -destination=../../../mocks/tui/view/nodview.go . NodeView
type NodeView interface {
	View
	UpdateNode(
		ds samplelib.DataSource,
		node samplelib.Node,
		nodeSelected func(node samplelib.Node),
		sampleSelected func(sample samplelib.Sample),
		nodeChosen func(node samplelib.Node),
		sampleChosen func(sample samplelib.Sample),
	)
}

//go:generate mockgen -destination=../../../mocks/tui/view/infoview.go . InfoView
type InfoView interface {
	View
	TextView
	UpdateNode(ds samplelib.DataSource, node samplelib.Node)
	UpdateSample(ds samplelib.DataSource, sample samplelib.Sample)
}

//go:generate mockgen -destination=../../../mocks/tui/view/logview.go . LogView
type LogView interface {
	io.Writer
}

type ControlPanel interface {
	ShowMainControls()
	ShowEditControls()
}

type Control struct {
	Label  string
	Keys   []string
	Action func()
}
