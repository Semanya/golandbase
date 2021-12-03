package generator

import (
	"app/internal/config"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"text/template"
)

// Task04 - функция для генерации маршалера структуры в мапу
func MarshallerGenerator(marshallerTemplate string, structName string, inFilePath string, outFilePath string) error {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, inFilePath, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	Imya := []string{}
	var paket string

	ast.Inspect(node, func(n ast.Node) bool {
		switch s := n.(type) {
		case *ast.File:
			paket = s.Name.Name
			fmt.Println(paket)
		case *ast.StructType:
			mmm := s.Fields.List
			for _, key := range mmm {
				ppp := key.Names
				for _, keyy := range ppp {
					fmt.Println(keyy.Name)
					Imya = append(Imya, keyy.Name)
				}
			}

		}
		return true

	})

	t, err := template.New("test").Parse(string(marshallerTemplate))
	if err != nil {
		panic(err)
	}

	vvv := &config.Config{
		Name:        "aaa",
		Host:        "bbb",
		Port:        123,
		Environment: map[string]string{"bacon": "delicious"},
		Volumes:     []string{"a", "b", "c", "d", "e", "f", "g"},
	}

	type sss struct {
		Packeganame string
		Structura   *config.Config
		Fff         []string
		StrName     string
	}

	svsv := &sss{
		Packeganame: paket,
		Structura:   vvv,
		Fff:         Imya,
		StrName:     structName,
	}

	outFile, err := os.Create(outFilePath)
	if err != nil {
		panic(err)
	}
	// defer os.Remove(outFilePath)
	defer outFile.Close()

	if err := t.Execute(outFile, svsv); err != nil {
		panic(err)
	}
	return nil
}
