package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

type Config struct {
	Name       string
	Port       string
	ReplicaSet int
	ImageName  string
	Tag        string
	EnvPath    string
}

func Task03() {
	yamlConfiTemplate, err := ioutil.ReadFile("/Users/u19502010/rebrainme/golandbase/module07/03_task/assets/template/config_template.yml")
	if err != nil {
		panic(err)
	}

	_ = ConfigGenerate(string(yamlConfiTemplate), "/Users/u19502010/rebrainme/golandbase/module07/03_task/config.yml")
}

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

func generate(tmpl string, outfilePath string, fields interface{}) error {
	t, err := template.New("test").Parse(string(tmpl))
	if err != nil {
		panic(err)
	}

	outFile, err := os.Create(outfilePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	if err := t.Execute(outFile, fields); err != nil {
		panic(err)
	}

	return nil
}

func main() {
	// Раскоментировать, когда дойдем до задания номер 3
	Task03()
	// Раскоментировать, когда дойдем до задания номер 4
	// Task04()
	// Раскоментировать, когда дойдем до задания номер 5
	// Task05()

	fmt.Println("Hello, from 07 module")
}
