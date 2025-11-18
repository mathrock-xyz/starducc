package main

import "github.com/spf13/cobra"

var lock = &cobra.Command{
	Use:   "lock",
	Short: "",
	RunE:  func(cmd *cobra.Command, args []string) error {},
}
