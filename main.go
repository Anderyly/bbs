package main

import (
	"bbs/cmd/serve"
	"bbs/cmd/sql"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "hm",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func main() {
	rootCmd.AddCommand(sql.Cmd())
	rootCmd.AddCommand(serve.Cmd())

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
