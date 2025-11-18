package main

import "github.com/spf13/cobra"

var cp = &cobra.Command{
	Use:   "cp",
	Short: "",
	RunE:  func(cmd *cobra.Command, args []string) error {},
}
