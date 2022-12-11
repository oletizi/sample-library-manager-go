package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"log"
)

func main() {
	boxLeft := tview.NewBox().SetBorder(true).SetTitle("Left Side")
	boxRight := tview.NewBox().SetBorder(true).SetTitle("Right Side")
	logView := tview.NewTextView()
	logView.SetBorder(true).SetTitle("Log")
	logView.Clear()
	logger := &Logger{view: logView}

	flex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(boxLeft, 0, 1, true).
		AddItem(boxRight, 0, 1, true).
		AddItem(logView, 0, 1, true)

	app := tview.NewApplication()
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		logger.Println(fmt.Sprintf("Event name: %s, rune: %d", event.Name(), event.Rune()))
		return event
	})
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

type Logger struct {
	view *tview.TextView
}

func (l *Logger) Println(msg any) {
	w := l.view.BatchWriter()
	defer func(w tview.TextViewWriter) {
		err := w.Close()
		if err != nil {
			log.Print(err)
		}
	}(w)
	_, err := fmt.Fprintln(w, msg)
	if err != nil {
		log.Print(err)
	}
}
