package helper

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFileToString(file string) string {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(buf)
}

func WriteToFile(value, fileName string) error {
	sqlFile := []byte(value)

	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.Write(sqlFile)
	if err != nil {
		return err
	}
	return nil
}

func GetAllFileInFolder(rootFolder string) []string {
	paths := []string{}
	folders, err := ioutil.ReadDir(rootFolder)
	if err != nil {
		return nil
	}
	for _, f := range folders {
		subFolder := rootFolder + "/" + f.Name()
		if f.IsDir() {
			paths = append(paths, GetAllFileInFolder(subFolder)...)
		} else {
			paths = append(paths, subFolder)
		}
	}
	return paths
}

func CreateNewFolder(mean, rootFolder string) error {
	folders, err := ioutil.ReadDir(rootFolder)
	if err != nil {
		return err
	}
	for _, f := range folders {
		subFolder := rootFolder + "/" + f.Name()
		if f.IsDir() {
			os.MkdirAll(mean+subFolder, os.ModePerm)
			CreateNewFolder(mean, subFolder)
		}
	}
	return nil
}
