package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/kfelter/termpi/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
