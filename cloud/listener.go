package cloud

import (
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type selfListener struct{}

// ProgressChangedCallback 自定义进度回调，需要实现 ProgressChangedCallback 方法
func (l *selfListener) ProgressChangedCallback(event *cos.ProgressEvent) {
	bar := progressbar.NewOptions64(event.TotalBytes,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("[文件上传:]"),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
	bar.Add64(event.ConsumedBytes)
}
