package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ylinyang/cloudstation/cloud"
	"github.com/ylinyang/cloudstation/test"
)

var filename, path string

// upload 子指令
var upload = &cobra.Command{
	Use:   "upload",
	Short: "upload 文件上传",
	Long:  "upload 文件上传",
	RunE: func(cmd *cobra.Command, args []string) error {
		switch ossProvider {
		case "tencent":
			fmt.Println("上传盘：tencent")
			err := cloud.NewTenCent(test.BucketUrl, test.SecretId, test.SecretKey).UpLoad(path, filename, test.SecretId, test.SecretKey)
			return err
		case "aliyah":
			fmt.Println("没有实现[aliyah]的方法, 因为这只是一个测试, 仅此而已")
		default:
			fmt.Println("无法向该云盘存储文件, 因为没有它的存在, ", ossProvider)
		}
		return nil
	},
}

func init() {
	// 向root注册子命令
	RootCmd.AddCommand(upload)
	// 向upload命令注入参数
	upload.PersistentFlags().StringVarP(&path, "path", "p", "", "云上文件路径")
	upload.PersistentFlags().StringVarP(&filename, "filename", "f", "", "文件名")
}
