package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	pathin       string = "/Users/u19502010/golndbase/golandbase/module02/07_task/07_task_in.txt"
	pathout      string = "/Users/u19502010/golndbase/golandbase/module02/07_task/data_out.txt"
	resultString string
	i            = 0
)

func main() {
	wrеOnFile(pathin, pathout)
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
	//	for s.Scan() {
	//		resultString = strings.SplitAfter(s.Text(), "|")
	//		fmt.Println(s.Text())
	//		fmt.Printf(" and has type: %T\n", s.Text())
	//        strings.SplitAfterN(s.Text(), "|", 3)
	//		fmt.Println(strings.SplitAfterN(s.Text(), "|", 3))
	//		w.WriteString(resultString)
	//	}
	for s.Scan() {
		words := strings.Split(s.Text(), "|")
		i++
		fmt.Println("Row:", i)
		//		fmt.Println(words)
		//fmt.Printf("Row: %d\n", words)
		//		fmt.Printf("Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n", words)
		for _, word := range words {
			//			println(idx)
			fmt.Printf("is: %s\n", word)
		}
	}
	defer w.Flush()
}
