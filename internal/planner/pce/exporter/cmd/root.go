package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd 在没有任何子命令的情况下调用时的基本命令
var rootCmd = &cobra.Command{
	Use:   "exporter",
	Short: "An exporter suitable for exporting xlsx configuration templates into go language configuration code and json data files. | 一个适合将 xlsx 配置模板导出为 go 语言配置代码和 json 数据文件的导出器。",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

// Execute 将所有子命令添加到根命令并适当设置标志。这是由 main.main() 调用的。 rootCmd 只需要发生一次
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
