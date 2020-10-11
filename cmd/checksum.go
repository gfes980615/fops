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
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"hash"
	"io"
	"os"
)

// checksumCmd represents the checksum command
var (
	checksumCmd = &cobra.Command{
		Use:   "checksum",
		Short: "Print checksum of file",
		Long:  `Print checksum of file`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Please enter one checksum method: 'md5', 'sha1', 'sha256'")
				return
			}
			result, err := checksumFunc(args[0], fileName)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(result)
		},
	}
	checkSumFlag string
)

func init() {
	rootCmd.AddCommand(checksumCmd)
	checksumCmd.Flags().StringVarP(&fileName, "file", "f", "", "enter file name")
}

func checksumFunc(method, file string) (string, error) {
	if err := checkFileExist(file); err != nil {
		return "", err
	}

	var h hash.Hash
	switch method {
	case "md5":
		h = md5.New()
	case "sha1":
		h = sha1.New()
	case "sha256":
		h = sha256.New()
	default:
		return "", errors.New("Please enter one checksum method: 'md5', 'sha1', 'sha256'")

	}

	f, _ := os.Open(file)
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
