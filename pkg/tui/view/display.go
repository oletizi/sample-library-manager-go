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
	"bytes"
	_ "embed"
	"github.com/oletizi/samplemgr/pkg/samplelib"
	"github.com/oletizi/samplemgr/pkg/tui"
	"log"
	"text/template"
)

//go:embed templates/nodeText.tmpl
var nodeTextTemplateString string

//go:embed templates/sampleText.tmpl
var sampleTextTemplateString string

//go:embed templates/listing.tmpl
var listingTemplateString string

//go:generate mockgen -destination=../../../mocks/tui/view/displayer.go . Display
type Display interface {
	DisplayNodeAsText(node samplelib.Node) string
	DisplayNodeAsListing(node samplelib.Node, isParent bool) string
	DisplaySampleAsListing(sample samplelib.Sample) string
	DisplaySampleAsText(sample samplelib.Sample) string
}

type display struct {
	logger                *log.Logger
	errorHandler          tui.ErrorHandler
	nodeTextTemplate      *template.Template
	nodeListingTemplate   *template.Template
	sampleTextTemplate    *template.Template
	sampleListingTemplate *template.Template
}

func (d *display) DisplaySampleAsListing(sample samplelib.Sample) string {
	data := struct{ Name string }{sample.Name()}
	return render(d.sampleListingTemplate, data, d.errorHandler)
}

func (d *display) DisplaySampleAsText(sample samplelib.Sample) string {
	data := struct {
		Name string
		Path string
	}{
		sample.Name(),
		sample.Path(),
	}
	return render(d.sampleTextTemplate, data, d.errorHandler)
}

func render(template *template.Template, data any, handler tui.ErrorHandler) string {
	buf := new(bytes.Buffer)
	err := template.Execute(buf, data)
	if err != nil {
		handler.Handle(err)
		return "error"
	}
	return buf.String()
}

func (d *display) DisplayNodeAsText(node samplelib.Node) string {
	data := struct {
		Name string
		Path string
	}{
		node.Name(),
		node.Path(),
	}
	return render(d.nodeTextTemplate, data, d.errorHandler)
}

func (d *display) DisplayNodeAsListing(node samplelib.Node, isParent bool) string {
	if isParent {
		return ".."
	}
	data := struct{ Name string }{node.Name()}
	return render(d.nodeListingTemplate, data, d.errorHandler)
}

func NewDisplay(logger *log.Logger, errorHandler tui.ErrorHandler) (Display, error) {
	nodeTextTemplate, err := template.New("nodeTextTemplate").Parse(nodeTextTemplateString)
	errorHandler.Handle(err)

	sampleTextTemplate, err := template.New("sampleTextTemplate").Parse(sampleTextTemplateString)
	errorHandler.Handle(err)

	nodeListingTemplate, err := template.New("listTemplate").Parse(listingTemplateString)
	errorHandler.Handle(err)

	// For now, these are the same
	sampleListingTemplate := nodeListingTemplate

	return &display{
		logger:                logger,
		nodeTextTemplate:      nodeTextTemplate,
		nodeListingTemplate:   nodeListingTemplate,
		sampleTextTemplate:    sampleTextTemplate,
		sampleListingTemplate: sampleListingTemplate,
		errorHandler:          errorHandler}, nil
}
