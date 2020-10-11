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
	"github.com/gabriel-vasile/mimetype"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// linecountCmd represents the linecount command
var (
	linecountCmd = &cobra.Command{
		Use:   "linecount",
		Short: "Print line count of file",
		Long:  `Print line count of file`,
		Run: func(cmd *cobra.Command, args []string) {
			count, err := fileLineCount(fileName)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(count)
		},
	}
	fileName string
)

const (
	File   = "file"
	Folder = "folder"
)

func init() {
	rootCmd.AddCommand(linecountCmd)
	linecountCmd.Flags().StringVarP(&fileName, "file", "f", "", "file")
}

func fileLineCount(file string) (int, error) {
	text, err := readFile(file)
	if err != nil {
		return -1, err
	}
	return len(strings.Split(text, "\n")), nil
}

func readFile(file string) (string, error) {
	if err := checkFileExist(file); err != nil {
		return "", err
	}

	b, _ := ioutil.ReadFile(file)
	if checkFileIsBinary(b) {
		return "", fmt.Errorf("error: Cannot do linecount for binary file '%s'", file)
	}
	return string(b), nil
}

func checkFileIsBinary(b []byte) bool {
	detectedMIME := mimetype.Detect(b)
	isBinary := true
	for mime := detectedMIME; mime != nil; mime = mime.Parent() {
		if mime.Is("text/plain") {
			isBinary = false
		}
	}
	return isBinary
}

func checkFileExist(file string) error {
	f := checkFileOrFolder(file)
	if _, err := os.Stat(file); os.IsNotExist(err) {
		switch f {
		case File:
			if strings.HasSuffix(err.Error(), "The system cannot find the file specified.") {
				err = fmt.Errorf("error: No such file '%s'", file)
			}
			return err
		case Folder:
			if strings.HasSuffix(err.Error(), "The system cannot find the file specified.") {
				err = fmt.Errorf("error: Expected file got directory '%s'", file)
			}
			return err
		}
	}
	return nil
}

func checkFileOrFolder(f string) string {
	sub := strings.Split(f, "/")
	if len(sub) == 1 {
		return File
	}

	match, err := regexp.MatchString("(.*?)\\.(.*?)", sub[len(sub)-1])
	if err != nil {
		fmt.Println(err)
	}
	if match {
		return File
	} else {
		return Folder
	}
	return ""
}
