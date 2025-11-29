package main

import (
	"os/user"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

var whoami = &cobra.Command{
	Use:   "whoami",
	Short: "Show current logged-in user",
	RunE:  func(cmd *cobra.Command, args []string) (err error) {},
}
