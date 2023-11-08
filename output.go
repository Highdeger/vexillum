package vexillum

import (
	"log"
	"os"
)

var (
	loggerWarning = log.New(os.Stdout, "flag warning: ", 0)
	loggerError   = log.New(os.Stderr, "flag error: ", 0)
)
