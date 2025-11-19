package main

import (
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(up)
	root.AddCommand(save)
	root.AddCommand(rm)
	root.AddCommand(del)
}

var root = &cobra.Command{
	Use:   "star",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
