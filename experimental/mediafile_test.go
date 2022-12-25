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

package experimental

import (
	"github.com/h2non/filetype"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path"
	"testing"
)

func TestFileFormats(t *testing.T) {
	pre := path.Join("..", "test", "data", "library", "one-level")
	file, err := os.Open(path.Join(pre, "hh.mov"))
	assert.Nil(t, err)

	// read enough of the file to get the header
	head := make([]byte, 1024)
	read, err := file.Read(head)
	assert.Nil(t, err)
	log.Printf("Bytes read: %v", read)
	//assert.Equal(t, 1024, read)

	// match the header
	match, err := filetype.Match(head)
	log.Printf("Match: %v:", match)
	log.Printf("Extension: %v", match.Extension)
	log.Printf("Type: %v", match.MIME.Type)
	log.Printf("Value: %v", match.MIME.Value)
	log.Printf("Subtype: %v", match.MIME.Subtype)

}
