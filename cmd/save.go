package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var save = &cobra.Command{
	Use:   "save",
	Short: "",
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

		descriptor, err := os.Open(name)
		defer descriptor.Close()

		if err != nil {
			return
		}

		var data bytes.Buffer
		writer := multipart.NewWriter(&data)
		defer writer.Close()

		fw, err := writer.CreateFormField("file")
		if err != nil {
			return
		}

		if _, err = io.Copy(fw, descriptor); err != nil {
			return
		}

		req, _ := http.NewRequest("POST", "http://app.starducc.mathrock.xyz", &data)

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
