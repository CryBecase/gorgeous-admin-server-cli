package main

import (
	"gorgeous-admin-server-cli/cmd"
	"gorgeous-admin-server-cli/internal/log"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			if e, ok := err.(error); ok {
				log.Error(e.Error())
			} else {
				log.Error("未知错误")
			}
		}
	}()
	cmd.Execute()
}
