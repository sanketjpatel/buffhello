package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/sanketjpatel/buffhello/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
