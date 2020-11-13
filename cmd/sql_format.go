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
	"io/ioutil"
	"regexp"
	"strings"
)

func init() {
	rootCmd.AddCommand(formatCmd)
	formatCmd.Flags().StringVarP(&fileName, "file", "f", "", "enter sql file")
	formatCmd.Flags().StringVar(&folder, "folder", "", "enter folder")
	//formatCmd.MarkFlagRequired("file")
}

const (
	folderMean_New = "new_"
)

// checksumCmd represents the checksum command
var (
	formatCmd = &cobra.Command{
		Use:   "sql_format",
		Short: "format your sql",
		Long:  "format your sql, let your column like `table`",
		RunE:  runSqlFormatCommand,
	}
	sqlFolderMap = make(map[string]string)

	specialWordMap = map[string]bool{
		"select":      true,
		"count":       true,
		"coalesce":    true,
		"sum":         true,
		"as":          true,
		"from":        true,
		"and":         true,
		"or":          true,
		"ifnull":      true,
		"where":       true,
		"if":          true,
		"exists":      true,
		"inner":       true,
		"join":        true,
		"left":        true,
		"find_in_set": true,
		"union":       true,
		"all":         true,
		"in":          true,
		"max":         true,
		"limit":       true,
	}
)

func runSqlFormatCommand(cmd *cobra.Command, args []string) error {
	filePath := helper.GetAllFileInFolder(folder)
	helper.CreateNewFolder(folderMean_New, folder)
	initFolderMap(folderMean_New, filePath)
	for _, file := range filePath {
		createFormatSQL(file)
	}
	return nil
}

func initFolderMap(mean string, filePath []string) {
	for _, file := range filePath {
		fmt.Println(file)
		sqlFolderMap[file] = mean + file
	}
}

func createFormatSQL(file string) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	originSQL := string(buf)
	sql := sqlFormat(originSQL)
	if err := helper.WriteToFile(sql, sqlFolderMap[file]); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("create new %s sql \n", file)
}

func sqlFormat(originSQL string) string {
	sqlLineString := strings.Split(originSQL, "\r\n")
	sql := ""
	for _, line := range sqlLineString {
		for _, word := range strings.Split(line, " ") {
			if checkFormat(word) {
				sql += word + " "
				continue
			}
			// 特殊字串轉成大寫
			if specialWordMap[strings.ToLower(word)] {
				sql += strings.ToUpper(word) + " "
				continue
			}
			// COALESCE(SUM(tmp.transfer_amount),0)
			match, _ := regexp.MatchString("([a-zA-Z_]+)\\(([a-zA-Z_]+)\\(([a-zA-Z_]+)\\.([a-zA-Z_]+)\\),(.*)", word)
			if match {
				r, _ := regexp.Compile("([a-zA-Z_]+)\\(([a-zA-Z_]+)\\(([a-zA-Z_]+)\\.([a-zA-Z_]+)\\),(.*)")
				matches := r.FindStringSubmatch(word)
				if specialWordMap[strings.ToLower(matches[1])] {
					matches[1] = strings.ToUpper(matches[1])
				}
				if specialWordMap[strings.ToLower(matches[2])] {
					matches[2] = strings.ToUpper(matches[2])
				}
				sql += fmt.Sprintf("%s(%s(`%s`.`%s`),%s ", matches[1], matches[2], matches[3], matches[4], matches[5])
				continue
			}

			// find_in_set(view_member_reseller_agents.share_login,
			// 解析 '(' , '.' , ',' 之間的英文
			match, _ = regexp.MatchString("(.*)\\((.*)\\.(.*),", word)
			if match {
				r, _ := regexp.Compile("(.*)\\((.*)\\.(.*),")
				matches := r.FindStringSubmatch(word)
				if specialWordMap[strings.ToLower(matches[1])] {
					matches[1] = strings.ToUpper(matches[1])
				}
				sql += fmt.Sprintf("%s(`%s`.`%s`, ", matches[1], matches[2], matches[3])
				continue
			}
			// MAX(withdraw.audit_time)
			// 解析 '(' '.' ')' 之間的英文
			match, _ = regexp.MatchString("(.*)\\((.*)\\.(.*)\\)", word)
			if match {
				r, _ := regexp.Compile("(.*)\\((.*)\\.(.*)\\)")
				matches := r.FindStringSubmatch(word)
				if specialWordMap[strings.ToLower(matches[1])] {
					matches[1] = strings.ToUpper(matches[1])
				}
				sql += fmt.Sprintf("%s(`%s`.`%s`) ", matches[1], matches[2], matches[3])
				continue
			}
			// IFNULL(deposit.transfer_amount
			// 解析 '(' 與 '.' 兩側的英文
			match, _ = regexp.MatchString("(.*)\\((.*)\\.(.*)", word)
			if match {
				r, _ := regexp.Compile("(.*)\\((.*)\\.(.*)")
				matches := r.FindStringSubmatch(word)
				if specialWordMap[strings.ToLower(matches[1])] {
					matches[1] = strings.ToUpper(matches[1])
				}
				sql += fmt.Sprintf("%s(`%s`.`%s` ", matches[1], matches[2], matches[3])
				continue
			}
			// view_member_agent.share_login,
			// 主要解析 '.' ',' 之間的英文
			match, _ = regexp.MatchString("([a-z_]+)\\.([a-z_]+),", word)
			if match {
				r, _ := regexp.Compile("([a-z_]+)\\.([a-z_]+),")
				matches := r.FindStringSubmatch(word)
				sql += fmt.Sprintf("`%s`.`%s`, ", matches[1], matches[2])
				continue
			}
			// view_member_agent.share_login
			// 主要解析 '.' 兩側的英文
			match, _ = regexp.MatchString("([a-z_]+)\\.([a-z_]+)", word)
			if match {
				r, _ := regexp.Compile("([a-z_]+)\\.([a-z_]+)")
				matches := r.FindStringSubmatch(word)
				sql += fmt.Sprintf("`%s`.`%s` ", matches[1], matches[2])
				continue
			}
			// (IFNULL({member_share_login},-99)=-99
			// 主要解析兩個 '(' 之間的英文 (ex IFNULL)
			match, _ = regexp.MatchString("^\\(([a-z_]+)\\((.*)", word)
			if match {
				r, _ := regexp.Compile("^\\(([a-z_]+)\\((.*)")
				matches := r.FindStringSubmatch(word)
				if specialWordMap[strings.ToLower(matches[1])] {
					matches[1] = strings.ToUpper(matches[1])
				}
				sql += fmt.Sprintf("(%s(%s ", matches[1], matches[2])
				continue
			}

			// action_code,
			// 解析 ',' 前的英文
			match, _ = regexp.MatchString("([a-z_]+),", word)
			if match {
				r, _ := regexp.Compile("([a-z_]+),")
				matches := r.FindStringSubmatch(word)
				sql += fmt.Sprintf("`%s`, ", matches[1])
				continue
			}
			// action_code
			// 普通英文的處理
			match, _ = regexp.MatchString("^([a-z_]+)$", word)
			if match {
				r, _ := regexp.Compile("^([a-z_]+)$")
				matches := r.FindStringSubmatch(word)
				sql += fmt.Sprintf("`%s` ", matches[1])
				continue
			}

			sql += word + " "

		}
		sql += "\n"
	}
	return sql
}

func checkFormat(word string) bool {
	match, _ := regexp.MatchString("`(.*)`", word)
	if match {
		return true
	}
	return false
}
