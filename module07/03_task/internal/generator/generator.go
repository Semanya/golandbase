package generator

import (
	"os"
	"text/template"
)

// const (
// 	marshallStructKey       = "marshallStruct"
// 	marshallStructFieldsKey = "marshallStructFields"
// )

// Task04 - функция для генерации маршалера структуры в мапу
// func MarshallerGenerator(marshallerTemplate string, structName string, inFilePath string, outFilePath string) error {
// 	return nil
// }

// Task03 - структура для yaml конфигурации
// Если нужно расширить yaml файл, тогда эту структуру нужно также расширить
// нужными параметрами, yaml файл и эта структура должны быть идентичными.
type Config struct {
	Name       string
	Port       string
	ReplicaSet int
	ImageName  string
	Tag        string
	EnvPath    string
}

// Task03 - функция генерации конфига
func ConfigGenerate(tmpl string, outFilePath string) error {
	config := &Config{
		Name:       "etcd-druid",
		Port:       "9090",
		ReplicaSet: 3,
		ImageName:  "eu.gcr.io/gardener-project/gardener/etcd-druid:v0.5.2",
		Tag:        "prod",
		EnvPath:    "/tmp/test",
	}

	errFromGen := generate(tmpl, outFilePath, config)
	if errFromGen != nil {
		return errFromGen
	}
	return nil
}

// // Основная функция которая должна производить генерацию
func generate(tmpl string, outfilePath string, fields interface{}) error {
	t, err := template.New("test").Parse(string(tmpl))
	if err != nil {
		panic(err)
	}

	outFile, err := os.Create(outfilePath)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	if err := t.Execute(outFile, fields); err != nil {
		panic(err)
	}
	return nil
}
