package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var delete = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) < 1 {
			return fmt.Errorf("This command takes one argument")
		}

		name := args[0]
		if name == "" {
			return fmt.Errorf("This command takes one argument")
		}

		file, err := os.Stat(name)
		if err != nil {
			return
		}

		if file.IsDir() {
			return fmt.Errorf("This command cannot accept folder")
		}

		req, _ := http.NewRequest("DELETE", "http://app.starducc.mathrock.xyz/delete/"+file.Name(), nil)

		token, err := bearer()
		if err != nil {
			return
		}

		req.Header.Set("Authorization", "Bearer "+token)

		request := new(http.Client)

		res, _ := request.Do(req)

		if res.StatusCode != http.StatusOK {
			msg, err := parse(res.Body)
			if err != nil {
				return err
			}

			return fmt.Errorf(msg)
		}

		log.Info("Succes")
		return
	},
}
