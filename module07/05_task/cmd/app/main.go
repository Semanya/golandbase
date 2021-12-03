package main

import (
	"fmt"
	"module07/internal/monitors"

	"github.com/sirupsen/logrus"
)

// func Task03() {
// 	yamlConfiTemplate, err := ioutil.ReadFile("./assets/template/config_template.yml")
// 	if err != nil {
// 		panic(err)
// 	}

// 	_ = generator.ConfigGenerate(string(yamlConfiTemplate), "./config.yml")
// }

// func Task04() {
// 	marshallerTemplate, err := ioutil.ReadFile("./assets/template/marshaller.gotmpl")
// 	if err != nil {
// 		panic(err)
// 	}

// 	_ = generator.MarshallerGenerator(
// 		string(marshallerTemplate),
// 		"Config",
// 		"./internal/config/config.go",
// 		"./internal/config/codegen_marshaller.go",
// 	)
// }

// Заполнить логикой, когда дойдем до задания номер 5
func Task05() {
	logger := logrus.New()
	logger.SetLevel(6)
	entry := logrus.NewEntry(logger)
	w := monitors.NewMonitorWithLogrus(monitors.NewSimpleMonitor(), entry)
	w.Run()
}

func main() {
	// Раскоментировать, когда дойдем до задания номер 3
	// Task03()
	// Раскоментировать, когда дойдем до задания номер 4
	// Task04()
	// Раскоментировать, когда дойдем до задания номер 5
	Task05()

	fmt.Println("Hello, from 07 module")
}
