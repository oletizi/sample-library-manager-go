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
	"github.com/golang/mock/gomock"
	mock_samplelib "github.com/oletizi/samplemgr/mocks/samplelib"
	"testing"
)

func TestController_UpdateNode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	node := mock_samplelib.NewMockNode(ctrl)
	node.EXPECT().Null().MinTimes(1)

	node.Null()
}
