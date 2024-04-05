package main

import (
	"log"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

type colorWriter struct {
	*color.Color
}

func (w *colorWriter) Write(p []byte) (n int, err error) {
	return w.Print(string(p))
}

func getLogger(ctx *cli.Context, colorAttrs ...color.Attribute) *log.Logger {
	if len(colorAttrs) == 0 {
		colorAttrs = []color.Attribute{color.FgYellow}
	}
	c := color.New(colorAttrs...)
	c.SetWriter(ctx.App.ErrWriter)
	wr := colorWriter{c}

	return log.New(&wr, "", log.Ldate|log.Ltime|log.Lshortfile)
}