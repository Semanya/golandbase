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
)

func main() {
	CheckFileExist(pathout)
	logTime()
}

func logTime() {
	defer duration(track("Duration"))
	wrеOnFile(pathin, pathout)
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	fmt.Printf("%v: %v\n", msg, time.Since(start))
}
func wrеOnFile(pathin string, pathout string) {
	f, _ := os.Open(pathin)
	var lines []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	fo, err := os.OpenFile(pathout, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
	defer f.Close()
	defer w.Flush()
}

func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		_, err = os.Create(fileName)
	}
	return true
}
