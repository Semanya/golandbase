package main

import (
	"fmt"
	newcolor "gopackages/pkg/color"
	. "gopackages/pkg/wordz" //Добавляем пакет wordz через точку

	"github.com/huandu/xstrings"
)

func main() {
	newcolor.Greet()
	fmt.Println("Hello world")

	fmt.Println(Hello)    //Вызов переменной из пакета wordz
	fmt.Println(Random()) //Вызов функции из пакета wordz
	fmt.Println(xstrings.Shuffle("Moscow"))
}
