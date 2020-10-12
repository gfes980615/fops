package helper

import (
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"os"
	"regexp"
	"strings"
)

const (
	File   = "file"
	Folder = "folder"
)

func CheckFileIsBinary(b []byte) bool {
	detectedMIME := mimetype.Detect(b)
	isBinary := true
	for mime := detectedMIME; mime != nil; mime = mime.Parent() {
		if mime.Is("text/plain") {
			isBinary = false
		}
	}
	return isBinary
}

func CheckFileExist(file string) error {
	f := distinguishFileOrFolder(file)
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

func distinguishFileOrFolder(f string) string {
	sub := strings.Split(f, "/")
	if len(sub) == 1 {
		return File
	}

	match, _ := regexp.MatchString("(.*?)\\.(.*?)", sub[len(sub)-1])

	if match {
		return File
	} else {
		return Folder
	}
	return ""
}
