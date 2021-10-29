package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var (
	pathin       string = "/Users/u19502010/golndbase/golandbase/module02/06_task/06_task_in.txt"
	pathout      string = "/Users/u19502010/golndbase/golandbase/module02/06_task/06_task_out.txt"
	i, b         int
	resultString string
	start        = time.Now()
)

func main() {
	CheckFileExist(pathout)
	wrеOnFile(pathin, pathout)
	defer logTime()
}

func logTime() {
	duration := time.Since(start)
	fmt.Println(duration)
}

func wrеOnFile(pathin string, pathout string) {
	f, _ := os.Open(pathin)
	defer f.Close()
	var lines []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	fo, err := os.OpenFile(pathout, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	w := bufio.NewWriter(fo)
	for _, row := range lines {
		i++
		resultString = fmt.Sprint(i) + " " + row + "\n"
		nn, _ := w.WriteString(resultString)
		b = b + nn
	}
	defer fo.Close()
	defer fmt.Println("Количество байт", b, "Количество строк", i)
	defer w.Flush()
}

func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		_, err = os.Create(fileName)
	}
	return true
}
