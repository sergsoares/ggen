package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/karrick/godirwalk"
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

	dirname := "examples"
	err = godirwalk.Walk(dirname, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			if !strings.Contains(osPathname, "/") {
				return nil
			}
			// TODO: Open each file and apply template with data params.
			fmt.Printf("%s %s\n", de.ModeType(), osPathname)
			return nil
		},
	})
	if err != nil {
		panic(err)
	}

	evaluteFile(config.TemplatePath, config.OutputPath, out.Data)
}

func evaluteFile(path string, outputPath string, data interface{}) {
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
	f, _ := os.Create(outputPath)
	defer f.Close()

	err = t.Execute(f, data)
	if err != nil {
		panic(err)
	}
}
