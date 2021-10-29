package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	pathin  string = "/Users/u19502010/golndbase/golandbase/module02/07_task/07_task_in.txt"
	pathout string = "/Users/u19502010/golndbase/golandbase/module02/07_task/data_out.txt"
	result  string
	i       = 0
)

func main() {
	wrеOnFile(pathin, pathout)
	//	defer reOutFile()
}

func wrеOnFile(pathin string, pathout string) {
	f, _ := os.Open(pathin)
	defer f.Close()
	s := bufio.NewScanner(f)
	fo, err := os.OpenFile(pathout, os.O_APPEND|os.O_WRONLY, 0644)
	defer fo.Close()
	if err != nil {
		fmt.Println(err)
	}
	w := bufio.NewWriter(fo)
	for s.Scan() {
		words := strings.Split(s.Text(), "|")
		i++
		for _, word := range words {
			if word == "" {
				//				dat, _ := os.ReadFile(pathout)
				//				fmt.Print(string(dat))
				errMwssage := fmt.Sprintf("parse error: empty field on string %d", i)
				//				recover(errText)
				fmt.Println(errMwssage)
			}
		}

		result = fmt.Sprintf("Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n", i, words[0], words[1], words[2])
		fmt.Print(result)
		//w.WriteString(result)
	}
	defer w.Flush()
}

//func reOutFile() {
//	dat, _ := os.ReadFile(pathout)
//	fmt.Print(string(dat))
//}
