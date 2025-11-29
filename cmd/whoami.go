package main

import "github.com/spf13/cobra"

var whoami = &cobra.Command{
	Use:   "whoami",
	Short: "Show current logged-in user",
	RunE:  func(cmd *cobra.Command, args []string) error {},
}
