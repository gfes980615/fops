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

func GetAllFileInFolder(rootFolder string, folders []os.FileInfo) []string {
	paths := []string{}
	fn := func(rootFolder string, folders []os.FileInfo) {}
	fn = func(rootFolder string, folders []os.FileInfo) {
		for _, folder := range folders {
			file := rootFolder + "/" + folder.Name()
			if folder.IsDir() {
				subFolder, _ := ioutil.ReadDir(file)
				fn(file, subFolder)
			} else {
				paths = append(paths, file)
			}
		}
	}
	fn(rootFolder, folders)

	return paths
}

func CreateNewFolder(mean, rootFolder string, folders []os.FileInfo) {
	newFolderRoot := mean + rootFolder
	os.MkdirAll(newFolderRoot, os.ModePerm)
	fn := func(rootFolder string, folders []os.FileInfo) {}
	fn = func(rootFolder string, folders []os.FileInfo) {
		for _, folder := range folders {
			file := newFolderRoot + "/" + folder.Name()
			if folder.IsDir() {
				subFolder, _ := ioutil.ReadDir(file)
				os.MkdirAll(file, os.ModePerm)
				fn(file, subFolder)
			}
		}
	}
	fn(rootFolder, folders)
}
