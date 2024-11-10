package rofi

import (
	"os"
	"strconv"
)

type Env struct {
	Retv int
	Info string
	Data string
}

func ParseEnv() Env {
	e := Env{}
	e.Info = os.Getenv("ROFI_INFO")
	e.Data = os.Getenv("ROFI_DATA")
	if retv, err := strconv.Atoi(os.Getenv("ROFI_RETV")); err == nil {
		e.Retv = retv
	}

	return e
}
