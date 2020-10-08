/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
)

// checksumCmd represents the checksum command
var (
	checksumCmd = &cobra.Command{
		Use:   "checksum",
		Short: "checksum file",
		Long:  `provide three checksum method 1.md5 2.sha1 3.sha256`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(checksumFunc(args[0], fileName))
		},
	}
	check string
)

func init() {
	rootCmd.AddCommand(checksumCmd)
	checksumCmd.Flags().StringVarP(&fileName, "file", "f", "", "file")
}

func checksumFunc(method, file string) string {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	switch method {
	case "md5":
		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			fmt.Println(err)
		}
		return hex.EncodeToString(h.Sum(nil))
	case "sha1":
		h := sha1.New()
		if _, err := io.Copy(h, f); err != nil {
			fmt.Println(err)
		}
		return hex.EncodeToString(h.Sum(nil))
	case "sha256":
		h := sha256.New()
		if _, err := io.Copy(h, f); err != nil {
			fmt.Println(err)
		}
		return hex.EncodeToString(h.Sum(nil))
	}
	return ""
}
