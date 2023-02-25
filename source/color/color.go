package color

import "github.com/fatih/color"

var Green func(a ...interface{}) string = color.New(color.FgGreen).Add(color.BgBlack).Add(color.Bold).SprintFunc()
var WhiteItalic func(a ...interface{}) string = color.New(color.FgWhite).Add(color.Italic).SprintFunc()
