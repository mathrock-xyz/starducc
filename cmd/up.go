package main

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

var up = &cobra.Command{
	Use:   "up",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if args == nil {
			return errors.New("")
		}

		name := args[0]
		if name == "" {
			return errors.New("")
		}

		file, err := os.Stat(name)
		if err != nil {
			return
		}

		if file.IsDir() {
			return errors.New("")
		}
	},
}
