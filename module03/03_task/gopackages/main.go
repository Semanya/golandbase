package main

import (
	"fmt"
	newcolor "gopackages/pkg/color"
	"gopackages/pkg/wordz"

	"github.com/fatih/color"
	"github.com/huandu/xstrings"

	"github.com/Semanya/utils" //Импортируем нашу библиотеку как стороннюю зависимость
	utilsV2 "github.com/Semanya/utils/v2"
	utilsV3 "github.com/Semanya/utils/v3"
)

func main() {

	//С помощью нашей новой библиотеки проверяем слайс на наличие определенного значения
	//Если его там находим, то сообщим это в stdout и закончим выполнение программы
	isExistIntV3 := utilsV3.InSliceInt([]int{1, 2, 3, 4, 5}, 5)
	if isExistIntV3 {
		fmt.Println("Slice Int contain finding value")
		return
	}

	isExistV2 := utilsV2.InSlice(wordz.Words, "Two")
	if isExistV2 {
		fmt.Println("Using utilsV2.InSlice and find value ")
		return
	}

	isExistInt := utils.ContainsInt([]int{1, 2, 3, 4, 5}, 5)
	if isExistInt {
		fmt.Println("Slice Int contain finding value")
		return
	}
	isExist := utils.Contains(wordz.Words, "Two")
	if isExist {
		fmt.Println("Slice Words contain finding value")
		return
	}

	newcolor.Greet()
	fmt.Println("Hello world")
	color.Red("Hello world again")

	fmt.Println(wordz.Hello)
	wordz.Words = []string{"Moscow", "New-York", "Amsterdam", "Barcelona", "Paris"}
	fmt.Println(wordz.Random())

	wordz.Prefix = ""
	fmt.Println("Moscow")
	fmt.Println("One")

	fmt.Println(xstrings.Shuffle("Kiev"))
	fmt.Println(xstrings.Shuffle("Two"))

}
