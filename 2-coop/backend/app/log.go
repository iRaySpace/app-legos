package app

import (
	"github.com/fatih/color"
)

func LogError(err error) {
	color.Red("Err: %s", err.Error())
}
