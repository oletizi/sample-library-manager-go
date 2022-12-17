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

//go:generate mockgen -destination=../../mocks/tui/logger.go . Logger
type Logger interface {
	Print(v ...any)
	Println(v ...any)
	Printf(format string, v ...any)
	Panic(v ...any)
	Panicln(v ...any)
	Panicf(format string, v ...any)
	Error(v ...any)
	Errorln(v ...any)
	Errorf(v ...any)
}
