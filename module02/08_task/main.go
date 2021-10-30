package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	workPath = "/Users/u19502010/golndbase/golandbase"
)

var (
	pathin = workPath + "/module02/08_task/08_task_in.txt"
	result string
	i      = 0
)

func main() {
	readFile(pathin)
}

func readFile(pathin string) {
	f, err := os.Open(pathin)
	if err != nil {
		panic("Не удалось открыть файл")
	}
	defer f.Close()
	s := bufio.NewReader(f)
	for {
		_, err := s.ReadString('\n')
		i++
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
	}
	fmt.Printf("Total strings: %d", i)
}
