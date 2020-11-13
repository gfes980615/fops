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
	"strings"
)

func init() {
	rootCmd.AddCommand(replaceCmd)
	replaceCmd.Flags().StringVarP(&fileName, "file", "f", "", "enter sql file")
	replaceCmd.Flags().StringVar(&folder, "folder", "", "enter folder")
}

const (
	folderMean_test = "test_"
)

var (
	replaceCmd = &cobra.Command{
		Use:   "sql_replace",
		Short: "replace parameter has {}",
		Long:  `replace parameter has {}`,
		RunE:  runReplaceCommand,
	}
)

func runReplaceCommand(cmd *cobra.Command, args []string) error {
	filePath := helper.GetAllFileInFolder(folder)
	helper.CreateNewFolder(folderMean_test, folder)
	initFolderMap(folderMean_test, filePath)
	pMap := getParameter()
	for _, file := range filePath {
		sql := helper.ReadFileToString(file)
		for key, value := range pMap {
			sql = strings.ReplaceAll(sql, key, value)
		}
		sql += ";"
		createTestSQL(file, sql)
	}
	return nil
}

func createTestSQL(file, sql string) {
	if err := helper.WriteToFile(sql, sqlFolderMap[file]); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("create test %s sql \n", file)
}

func getParameter() map[string]string {
	parameter := helper.ReadFileToString("parameter.txt")
	parameters := strings.Split(parameter, "\n")
	parameterMap := make(map[string]string)
	for _, p := range parameters {
		ps := strings.Split(p, ":")
		if len(ps) < 2 {
			continue
		}
		parameterMap[ps[0]] = ps[1]
	}
	return parameterMap
}
