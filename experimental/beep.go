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

/*
func main_experimental() {
	flag.Parse()
	args := flag.Args()
	url := "test/data/library/multi-level/hh.wav"
	if len(args) > 0 {
		url = args[0]
	}

	context, err := audio.NewBeepContext()
	if err != nil {
		log.Fatal(err)
	}
	player, err := context.PlayerFor(url)
	if err != nil {
		log.Fatal(err)
	}

	defer func(player audio.Player) {
		err := player.Close()
		if err != nil {
			log.Panic(err)
		}
	}(player)

	const (
		play int = iota
		loop
		pause
		stop
	)
	handler := func(opt wmenu.Opt) error {
		var err error
		switch opt.ID {
		case play:
			fmt.Println("Play!")
			err = player.Play(func() { fmt.Println("Done playing!") })
		case loop:
			fmt.Println("Loop!")
			err = player.Loop(-1, func() { fmt.Println("Done looping!") })
		case pause:
			fmt.Println("Pause!")
			player.Pause()
		case stop:
			fmt.Println("Stop!")
			err = player.Stop()
		}

		return err
	}
	menu := wmenu.NewMenu("Options:")

	menu.Option("Play sound", play, false, handler)
	menu.Option("Loop sound", loop, false, handler)
	menu.Option("Pause/unpause sound", pause, false, handler)
	menu.Option("Stop sound", stop, false, handler)

	for {
		err := menu.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}
*/
