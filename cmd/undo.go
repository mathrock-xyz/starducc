package main

import "github.com/spf13/cobra"

var cmd = &cobra.Command{
	Use:   "undo",
	Short: "",
	RunE:  func(cmd *cobra.Command, args []string) error {},
}
