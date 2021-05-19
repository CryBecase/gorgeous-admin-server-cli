package file

import (
	"fmt"
)

type Rule struct {
	// 控制是否覆盖
	ally bool
	alln bool
}

func NewRule() *Rule {
	return &Rule{}
}

func (fr *Rule) Matched() bool {
	if fr.AlwaysMatched() {
		return true
	}
	if fr.AlwaysNotMatched() {
		return false
	}

	return fr.canCover()
}

func (fr *Rule) AlwaysMatched() bool {
	return fr.ally
}

func (fr *Rule) AlwaysNotMatched() bool {
	return fr.alln
}

func (fr *Rule) canCover() bool {
	text := ""
	_, _ = fmt.Scanln(&text)
	switch text {
	case "ally":
		fr.ally = true
		return true
	case "alln":
		fr.alln = true
		return true
	case "n":
		return false
	case "y":
		return true
	default:
		fmt.Println("please input y/n/ally/alln")
		return fr.canCover()
	}
}
