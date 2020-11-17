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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	_ "gorm.io/driver/mysql"
	"io/ioutil"
	"reflect"
)

func init() {
	rootCmd.AddCommand(jsonCompareCmd)
	jsonCompareCmd.Flags().StringVarP(&fileName, "file", "f", "", "enter sql file")
}

var (
	jsonCompareCmd = &cobra.Command{
		Use:   "json_compare",
		Short: "replace parameter has {}",
		Long:  `replace parameter has {}`,
		RunE:  runJsonCompareCommand,
	}
)

func runJsonCompareCommand(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("enter one compare file.")
	}
	keys := getMapKeys(fileName)
	fileOneMap := getFileDataToMap(fileName, keys)
	fileTwoMap := getFileDataToMap(args[0], keys)

	count := 0
	for key, _ := range fileOneMap {
		if fileTwoMap[key] {
			count++
		}
	}

	if len(fileOneMap) == len(fileTwoMap) && count == len(fileOneMap) && count == len(fileTwoMap) {
		fmt.Println("Two files are equally")
	} else {
		fmt.Println("some different")
	}

	return nil
}

func getFileDataToMap(file string, keys []string) map[string]bool {
	buf, _ := ioutil.ReadFile(file)
	fileResult := []interface{}{}
	json.Unmarshal(buf, &fileResult)

	resultMap := make(map[string]bool)
	for _, result := range fileResult {
		v := reflect.ValueOf(result)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		key := ""
		for _, k := range keys {
			key += fmt.Sprintf("%v/", v.MapIndex(reflect.ValueOf(k)))
		}
		resultMap[key] = true
	}

	return resultMap
}

// 為了固定順序
func getMapKeys(file string) []string {
	buf, _ := ioutil.ReadFile(file)
	fileResult := []interface{}{}
	err := json.Unmarshal(buf, &fileResult)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	v := reflect.ValueOf(fileResult[0])
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	result := []string{}
	for _, value := range v.MapKeys() {
		result = append(result, value.String())
	}
	return result
}
