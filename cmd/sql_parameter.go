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
	_ "gorm.io/driver/mysql"
	"regexp"
)

func init() {
	rootCmd.AddCommand(parameterCmd)
	parameterCmd.Flags().StringVarP(&fileName, "file", "f", "", "enter sql file")
	parameterCmd.Flags().StringVar(&folder, "folder", "", "enter folder")
}

// checksumCmd represents the checksum command
var (
	parameterCmd = &cobra.Command{
		Use:   "sql_parameter",
		Short: "get parameter from {}",
		Long:  `get parameter from {}`,
		RunE:  runParameterCommand,
	}
)

func runParameterCommand(cmd *cobra.Command, args []string) error {
	filePath := helper.GetAllFileInFolder(folder)
	parameterFileContent := ""
	parameterMap := make(map[string]bool)
	for _, file := range filePath {
		sql := helper.ReadFileToString(file)
		r, _ := regexp.Compile("\\{([a-z_0-9A-Z]+)}")
		matches := r.FindAllStringSubmatch(sql, -1)
		for _, match := range matches {
			if !parameterMap[match[0]] {
				parameterFileContent += fmt.Sprintf("%s:\n", match[0])
				parameterMap[match[0]] = true
			}
		}
	}
	fmt.Printf("get %d parameter > parameter.txt\n"+
		"Please enter your value by yourself.\n"+
		"And you can print command sql_replace let your sql parameter to replace", len(parameterMap))
	return helper.WriteToFile(parameterFileContent, "parameter.txt")
}
