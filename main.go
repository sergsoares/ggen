package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

type Params struct {
	Data interface{} `yaml:"data"`
}

func main() {
	fmt.Println("Lets go")

	path := "examples/fly.io.tpl"
	config_path := "examples/config.yml"

	file_config, err := os.Open(config_path)
	if err != nil {
		fmt.Errorf("Failure with path: %s", err)
	}

	content_config, err := ioutil.ReadAll(file_config)
	if err != nil {
		fmt.Errorf("Failure with file content: %s", err)
	}
	out := Params{}

	if err := yaml.Unmarshal(content_config, &out); err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Errorf("Failure with path: %s", err)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Errorf("Failure with file content: %s", err)
	}

	t, err := template.New("").Parse(string(content))
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, out.Data)
	if err != nil {
		panic(err)
	}
}
