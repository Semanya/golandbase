package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	workPath = "/Users/u19502010/rebrainme/golandbase"
)

var (
	pathin  = workPath + "/module02/07_task/07_task_in.txt"
	pathout = workPath + "/module02/07_task/data_out.txt"
	result  string
	i       = 0
)

func main() {
	parseFile(pathin, pathout)
}

func parseFile(pathin string, pathout string) {
	f, _ := os.Open(pathin)
	defer f.Close()
	s := bufio.NewScanner(f)
	fo, err := os.OpenFile(pathout, os.O_APPEND|os.O_WRONLY, 0644)
	defer fo.Close()
	if err != nil {
		fmt.Println(err)
	}
	for s.Scan() {
		words := strings.Split(s.Text(), "|")
		i++
		for _, word := range words {
			if word == "" {
				dat, _ := os.ReadFile(pathout)
				fmt.Print(string(dat))
				errMwssage := fmt.Sprintf("parse error: empty field on string %d", i)
				panic(errMwssage)
			}
		}

		result = fmt.Sprintf("Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n", i, words[0], words[1], words[2])
		fo.WriteString(result)
	}
}
