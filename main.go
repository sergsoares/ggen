package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

type Config struct {
	TemplatePath string `yaml:"template_path"`
	OutputPath   string `yaml:"output_path"`
}

type Params struct {
	Data interface{} `yaml:"data"`
}

const config_path = "ggen.yml"

func main() {
	file_config, err := os.Open(config_path)
	if err != nil {
		fmt.Errorf("Failure with path: %s", err)
	}

	content_config, err := ioutil.ReadAll(file_config)
	if err != nil {
		fmt.Errorf("Failure with file content: %s", err)
	}

	config := Config{}
	if err := yaml.Unmarshal(content_config, &config); err != nil {
		log.Fatal(err)
	}

	out := Params{}
	if err := yaml.Unmarshal(content_config, &out); err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(config.TemplatePath)
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

	// TODO: Create file or path too
	f, _ := os.Create(config.OutputPath)
	defer f.Close()

	err = t.Execute(f, out.Data)
	if err != nil {
		panic(err)
	}
}
