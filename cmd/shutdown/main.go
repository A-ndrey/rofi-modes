package main

import (
	"io"
	"log/slog"
	"os"
	"os/exec"

	"github.com/A-ndrey/rofi-modes"
)

var (
	PoweroffCmd = &rofi.Entry{
		Name: "poweroff",
		Run: func(env rofi.Env) error {
			return exec.Command("systemctl", "poweroff").Start()
		},
	}

	RebootCmd = &rofi.Entry{
		Name: "reboot",
		Run: func(env rofi.Env) error {
			return exec.Command("systemctl", "reboot").Start()
		},
	}
)

func main() {
	entries := []*rofi.Entry{
		PoweroffCmd,
		RebootCmd,
	}

	logWriter := os.Stderr
	logFile, err := os.OpenFile("/tmp/rofi-modes.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		logWriter = logFile
	}

	log := slog.New(slog.NewTextHandler(io.Writer(logWriter), nil))

	if err := rofi.Run(entries); err != nil {
		log.Error(err.Error())
	}
}
