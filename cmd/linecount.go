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
	"fmt"
	"github.com/gfes980615/fops/helper"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strings"
)

func init() {
	rootCmd.AddCommand(linecountCmd)
	linecountCmd.Flags().StringVarP(&fileName, "file", "f", "", "file")
	linecountCmd.MarkFlagRequired("file")
}

// linecountCmd represents the linecount command
var (
	linecountCmd = &cobra.Command{
		Use:     "linecount",
		Short:   "Print line count of file",
		Long:    `Print line count of file`,
		RunE:    runLineCountCommand,
		Example: "  fops linecount -f [filename] ",
	}
	fileName string
)

func runLineCountCommand(cmd *cobra.Command, args []string) error {
	count, err := fileLineCount(fileName)
	if err != nil {
		return err
	}

	fmt.Println(count)

	return nil
}

func fileLineCount(file string) (int, error) {
	text, err := readFileContent(file)
	if err != nil {
		return -1, err
	}
	return len(strings.Split(text, "\n")), nil
}

func readFileContent(file string) (string, error) {
	if err := helper.CheckFileExist(file); err != nil {
		return "", err
	}

	b, _ := ioutil.ReadFile(file)
	if helper.CheckFileIsBinary(b) {
		return "", fmt.Errorf("Cannot do linecount for binary file '%s'", file)
	}
	return string(b), nil
}
