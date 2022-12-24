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
	"bytes"
	"github.com/oletizi/samplemgr/pkg/util"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestLogger_Methods(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	v := "v"
	logger := util.NewLogger(log.New(buf, "", 0))
	logger.Print(v)
	assert.Equal(t, v+"\n", buf.String())

	buf.Reset()
	logger.Println(v)
	assert.Equal(t, v+"\n", buf.String())

	buf.Reset()
	logger.Printf("%s", v)
	assert.Equal(t, v+"\n", buf.String())
}
