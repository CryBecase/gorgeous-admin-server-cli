package log

import (
	"github.com/k0kubun/go-ansi"
	"github.com/mitchellh/colorstring"

	"gorgeous-admin-server-cli/global"
)

const (
	DebugTitle   = "[blue][Debug][reset] "
	InfoTitle    = "[green][Info][reset] "
	WarningTitle = "[magenta][Warning][reset] "
	ErrorTitle   = "[red][Error][reset] "
)

func Debug(str string) {
	if !global.Debug {
		return
	}

	_, _ = ansi.Println(colorstring.Color(DebugTitle + str))
}

func Info(str string) {
	_, _ = ansi.Println(colorstring.Color(InfoTitle + str))
}

func Warning(str string) {
	_, _ = ansi.Println(colorstring.Color(WarningTitle + str))
}

func Error(str string) {
	_, _ = ansi.Println(colorstring.Color(ErrorTitle + str))
}
