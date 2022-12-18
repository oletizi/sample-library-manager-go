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
	"errors"
	"github.com/golang/mock/gomock"
	mocktui "github.com/oletizi/samplemgr/mocks/tui"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorHandler(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	logger := mocktui.NewMockLogger(ctl)

	handler := NewErrorHandler(logger)
	assert.NotNil(t, handler)

	err := errors.New("my error")

	logger.EXPECT().Print(err)
	handler.Handle(err)
}
