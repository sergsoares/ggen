package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
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

	err = godirwalk.Walk(config.TemplatePath, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			fmt.Println(osPathname)
			if !strings.Contains(osPathname, "/") {
				return nil
			}

			isDir, err := isDirectory(osPathname)
			if err != nil {
				panic(err)
			}
			if isDir {
				return nil
			}
			pathSplited := strings.SplitN(osPathname, "/", 2)

			pathToSave := filepath.Join(config.OutputPath, pathSplited[1])
			pathDir := path.Dir(pathToSave)
			err = os.MkdirAll(pathDir, os.ModePerm)
			if err != nil {
				panic(err)
			}

			evaluteFile(osPathname, pathToSave, out.Data)
			return nil
		},
	})
	if err != nil {
		panic(err)
	}
}

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
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
