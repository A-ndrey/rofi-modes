package rofi

import (
	"errors"
	"os"
)

type Entry struct {
	Name       string
	RowOptions []RowOption
	Run        func(env Env) error
}

type CustomHandler func(args []string, env Env) error

type HotKeysHandler func(hotKey int, env Env) error

func Run(entries []*Entry, modeOpts ...ModeOption) error {
	if len(entries) == 0 {
		return errors.New("no commands specified")
	}

	opts := newOptions()
	opts.addMode(modeOpts)

	env := ParseEnv()

	if env.Retv == 0 {
		for _, entry := range entries {
			opts.addRow(entry.Name, entry.RowOptions)
		}

		_, err := os.Stdout.WriteString(opts.String())
		return err
	}

	args := os.Args[1:]

	if env.Retv == 1 {
		for _, entry := range entries {
			if entry.Name == args[0] {
				return entry.Run(env)
			}
		}
	}

	if env.Retv == 2 {
		return opts.customHandler(args, env)
	}

	if env.Retv >= 10 && env.Retv <= 28 {
		return opts.hotKeyHandler(env.Retv-9, env)
	}

	return errors.New("invalid return value")
}
