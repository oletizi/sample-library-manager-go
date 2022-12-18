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

package main

import (
	"flag"
	"fmt"
	"github.com/oletizi/samplemgr/pkg/samplelib"
	"github.com/oletizi/samplemgr/pkg/tui/tviewtui"
	"log"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	rootDir := "." // default
	if len(args) > 0 {
		rootDir = args[0]
		info, err := os.Stat(rootDir)
		if err != nil {
			log.Default().Fatal(err)
		}
		if !info.IsDir() {
			log.Default().Fatal("Not a directory: " + rootDir)
		}
	}
	ds := samplelib.NewFilesystemDataSource(rootDir)
	err := tviewtui.New(ds).Run()
	if err != nil {
		fmt.Print(err)
	}
}
