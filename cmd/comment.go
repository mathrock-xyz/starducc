package main

import "github.com/spf13/cobra"

var comment = &cobra.Command{
	Use:   "comment",
	Short: "",
	RunE:  func(cmd *cobra.Command, args []string) error {},
}
