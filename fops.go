package main
//
//import (
//	"flag"
//	"fmt"
//	"io/ioutil"
//	"strings"
//)
//
//var (
//	fileName string
//)
//
//func init() {
//	flag.StringVar(&fileName, "f", "", "the file")
//	//flag.Usage = usage
//}
//
//// 印出預設的說明
////func usage() {
////	fmt.Fprintf(os.Stderr, "Usage: math [options] [root]\n")
////	fmt.Fprintf(os.Stderr, "  Currently, there are four URI routes could be used:\n")
////	flag.PrintDefaults()
////}
//
//func main() {
//	flag.Parse()
//	fmt.Println(flag.Args())
//	b, err := ioutil.ReadFile(fileName)
//	if err != nil {
//		fmt.Print(err)
//	}
//	str := string(b)
//	fmt.Println(str)
//	fmt.Println(fileLineCount(str))
//}
//
//func fileLineCount(text string) int {
//	return len(strings.Split(text, "\n"))
//}
