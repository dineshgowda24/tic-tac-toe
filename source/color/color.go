package color

import "github.com/fatih/color"

var Green func(a ...interface{}) string = color.New(color.FgGreen).Add(color.BgBlack).Add(color.Bold).SprintFunc()
var BlackItalic func(a ...interface{}) string = color.New(color.FgBlack).Add(color.Italic).SprintFunc()
