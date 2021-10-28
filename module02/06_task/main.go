package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	pathin := "/Users/u19502010/golndbase/golandbase/module02/06_task/06_task_in.txt"
	pathout := "/Users/u19502010/golndbase/golandbase/module02/06_task/06_task_out.txt"
	CheckFileExist(pathout)
	getFile(pathin)
}

func getFile(path string) {
	f, _ := os.Open(path)
	defer f.Close()

	s := bufio.NewScanner(f)
	i := 0
	for s.Scan() {
		i++
		fmt.Println(i, s.Text())
	}
}

func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		_, err = os.Create(fileName)
	}
	return true
}
