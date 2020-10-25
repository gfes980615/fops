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
	"database/sql"
	"errors"
	"fmt"
	"github.com/gfes980615/fops/glob"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strings"

	_ "gorm.io/driver/mysql"
)

func init() {
	rootCmd.AddCommand(sqlCmd)
	sqlCmd.Flags().StringVarP(&fileName, "file", "f", "", "enter sql file")
	sqlCmd.Flags().StringVar(&folder, "folder", "", "enter folder")
}

// checksumCmd represents the checksum command
var (
	sqlCmd = &cobra.Command{
		Use:   "mysql",
		Short: "Execute your sql",
		Long:  `Execute your sql`,
		RunE:  runSqlCommand,
	}
	folder string
)

func runSqlCommand(cmd *cobra.Command, args []string) error {
	if len(folder) > 0 {
		folders, err := ioutil.ReadDir(folder)
		if err != nil {
			return err
		}
		for _, f := range folders {
			err = exec(folder + "/" + f.Name())
		}
		return err
	}

	if len(fileName) > 0 {
		return exec(fileName)
	}

	return errors.New("Please enter exec file or folder")
}

func exec(file string) error {
	failedDB, err := dbExec(file)
	if err != nil {
		return err
	}
	if len(failedDB) == 0 {
		fmt.Println("all success")
		return nil
	}
	for db, failedErr := range failedDB {
		fmt.Printf("DB: %s, error: %v", db, failedErr)
	}
	return errors.New("some Error")
}

func dbExec(file string) (map[string]error, error) {
	config := glob.Config
	sqlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	failedDB := make(map[string]error)
	for _, database := range config.Database {
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s",
			config.Username, config.Password, config.Address, database))
		defer db.Close()
		if err != nil {
			if _, ok := failedDB[database]; !ok {
				failedDB[database] = err
			}
			continue
		}
		tx, err := db.Begin()
		defer tx.Rollback()
		if err != nil {
			if _, ok := failedDB[database]; !ok {
				failedDB[database] = err
			}
			break
		}
		execs := strings.Split(string(sqlFile), ";")
		for i := 0; i < len(execs)-1; i++ {
			_, err = tx.Exec(execs[i])
			if err != nil {
				if _, ok := failedDB[database]; !ok {
					failedDB[database] = err
				}
				continue
			}
		}
		tx.Commit()
	}

	return failedDB, nil
}
