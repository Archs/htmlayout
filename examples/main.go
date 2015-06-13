package main

import (
	"github.com/Archs/go-htmlayout"
	. "github.com/Archs/htmlayout/declarative"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	gohl.EnableDebug()
	if _, err := (MainWindow{
		Layout: VBox{},
		Children: []Widget{
			HtmLayout{
				PageUrl: "a.html",
				OnKeyDown: func(key walk.Key) {
					println("key down", key)
				},
				// MinSize: Size{800, 600},
			},
			DateEdit{},
		},
	}).Run(); err != nil {
		panic(err)
	}
}
