package pb

import (
	"fmt"
	"time"

	"github.com/k0kubun/go-ansi"
	"github.com/mitchellh/colorstring"
	"github.com/schollz/progressbar/v3"

	"gorgeous-admin-server-cli/internal/log"
)

// 默认的进度条主题
// 样式参考 https://github.com/mitchellh/colorstring/blob/master/colorstring.go#L118
var defaultTheme = progressbar.Theme{
	Saucer:        "[blue]=[reset]", // 进度条样式; [reset] 会将之后的文本样式重置
	SaucerHead:    "[blue]>[reset]", // 进度条头部样式
	SaucerPadding: " ",
	BarStart:      "[",
	BarEnd:        "]",
}

func NewProgressBar(max int, description string) *progressbar.ProgressBar {
	finishCall := func() {
		_, _ = fmt.Fprint(ansi.NewAnsiStdout(), colorstring.Color("[yellow]Finished![reset]\n"))
	}

	bar := progressbar.NewOptions(
		max, // 任务总数
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),           // 使 windows console 识别 \x1b，可处理 escape sequences
		progressbar.OptionEnableColorCodes(true),                    // 开启颜色
		progressbar.OptionShowBytes(false),                          // 显示速度
		progressbar.OptionSetWidth(30),                              // 进度条长度
		progressbar.OptionThrottle(65*time.Millisecond),             // 进度条刷新间隔时间
		progressbar.OptionSetRenderBlankState(true),                 // 渲染 0%
		progressbar.OptionSetDescription(log.InfoTitle+description), // 描述
		progressbar.OptionSetTheme(defaultTheme),                    // 设置进度条主题
		progressbar.OptionOnCompletion(finishCall),                  // 完成之后的回调
	)

	return bar
}
