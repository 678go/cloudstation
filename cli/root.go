package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version bool
var (
	ossProvider string
	//id          string
	//key         string
)

// RootCmd 根指令
var RootCmd = &cobra.Command{
	Use:  "cloud-cli",
	Long: "cloud-cloud",
	RunE: func(cmd *cobra.Command, args []string) error {
		if version {
			fmt.Println("v0.0.1")
		}
		return nil
	},
}

func init() {
	// 永久命令 不分是否是子命令
	RootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "版本信息")
	RootCmd.PersistentFlags().StringVarP(&ossProvider, "ossProvider", "o", "", "上传对象支持[tencent/aliyah]")
	//	RootCmd.PersistentFlags().StringVarP(&id, "secretId", "i", "", "secretId信息")
	//	RootCmd.PersistentFlags().StringVarP(&key, "secretKey", "k", "", "secretId信息")
}
