package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init()  {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:"version",
	Short:"打印命令的版本",
	Long:"All software has versions. This is Hugo's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Matryoshka 平台 v4.0 -- Weiyi")
	},
}
