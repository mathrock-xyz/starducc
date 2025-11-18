package main

import "github.com/spf13/cobra"

var del = &cobra.Command{
	Use:   "del",
	Short: "",
	RunE:  func(cmd *cobra.Command, args []string) error {},
}
