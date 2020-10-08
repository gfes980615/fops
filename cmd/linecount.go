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
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"
)

// linecountCmd represents the linecount command
var (
	linecountCmd = &cobra.Command{
		Use:   "linecount",
		Short: "file line count",
		Long:  `file line count`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(fileLineCount(fileName))
		},
	}
	fileName string
)

func init() {
	rootCmd.AddCommand(linecountCmd)
	linecountCmd.Flags().StringVarP(&fileName, "file", "f", "", "file")
}

func fileLineCount(file string) int {
	return len(strings.Split(readFile(file), "\n"))
}

func readFile(file string) string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}
