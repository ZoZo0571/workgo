package main

import (
	"errors"
	"log"

	"github.com/awesome-gocui/gocui"
)

const XWIDTH = 80
const YWIDTH = 5

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	// maxX, maxY := g.Size()
	if v, err := g.SetView("main", 0, 0, XWIDTH, YWIDTH, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}

		if _, err := g.SetCurrentView("main"); err != nil {
			return err
		}
		for i := 10; i < 20; i++ {
			v.SetCursor(i, 3)
			v.EditWrite(rune('X'))
		}
		// fmt.Fprintln(v, "Hello world!")
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
