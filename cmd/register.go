package main

import "github.com/spf13/cobra"

var register = &cobra.Command{
	Use:   "register",
	Short: "Create new account",
	RunE:  func(cmd *cobra.Command, args []string) error {},
}
