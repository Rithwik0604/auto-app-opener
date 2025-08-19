package ui

import "github.com/common-nighthawk/go-figure"

func MakeAsciiArt() {

	banner := figure.NewColorFigure("Auto App Opener", "", "green", true)
	banner.Print()

}
