package utils

import (
	"github.com/charmbracelet/log"
	"os"
)

var Log = log.New(os.Stdout)

func init() {
	Log.SetPrefix("[raytracer]")
}
