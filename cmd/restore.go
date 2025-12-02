package main

import "github.com/spf13/cobra"

var restore = &cobra.Command{
	Use:   "restore",
	Short: "",
	RunE:  func(cmd *cobra.Command, args []string) error {},
}
