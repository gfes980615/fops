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
	"github.com/gfes980615/fops/helper"
	"github.com/spf13/cobra"
	"hash"
	"io"
	"os"
)

func init() {
	rootCmd.AddCommand(checksumCmd)
	checksumCmd.Flags().StringVarP(&fileName, "file", "f", "", "enter file name")
	checksumCmd.MarkFlagRequired("file")
	checksumCmd.Flags().BoolVar(&md5FlagBool, "md5", false, "md5 checksum")
	checksumCmd.Flags().BoolVar(&sha1FlagBool, "sha1", false, "sha1 checksum")
	checksumCmd.Flags().BoolVar(&sha256FlagBool, "sha256", false, "sha256 checksum")
}

const (
	MD5    = "md5"
	SHA1   = "sha1"
	SHA256 = "sha256"
)

// checksumCmd represents the checksum command
var (
	checksumCmd = &cobra.Command{
		Use:     "checksum",
		Short:   "Print checksum of file",
		Long:    `Print checksum of file`,
		RunE:    runChecksumCommand,
		Example: "  fops checksum -f [filename] [checksum flag]",
	}
	md5FlagBool    bool
	sha1FlagBool   bool
	sha256FlagBool bool
)

func runChecksumCommand(cmd *cobra.Command, args []string) error {
	var method string
	if md5FlagBool {
		method = MD5
	} else if sha1FlagBool {
		method = SHA1
	} else if sha256FlagBool {
		method = SHA256
	} else {
		return fmt.Errorf("Please enter one checksum flag: '--md5', '--sha1', '--sha256'")
	}

	result, err := checksumFunc(method, fileName)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func checksumFunc(method, file string) (string, error) {
	if err := helper.CheckFileExist(file); err != nil {
		return "", err
	}

	var h hash.Hash
	switch method {
	case MD5:
		h = md5.New()
	case SHA1:
		h = sha1.New()
	case SHA256:
		h = sha256.New()
	default:
		return "", fmt.Errorf("Please enter one checksum flag: '--md5', '--sha1', '--sha256'")
	}

	f, _ := os.Open(file)
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
