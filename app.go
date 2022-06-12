package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

var fm = template.FuncMap{
	"val": val, // format value
}

type app struct {
	tmpl *template.Template
}

func newApp(tmplContent string) (*app, error) {
	tmpl := template.Must(template.New("readme").Funcs(fm).Parse(tmplContent))
	return &app{tmpl: tmpl}, nil
}

func (a *app) ProcessDir(dir string, outputFile string) error {
	// read action file
	file := dir + "/action.yml"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("WARN: No action.yml file found in directory: %q", dir)
			return nil
		}
		return fmt.Errorf("failed to read file %q: %v", file, err)
	}

	// parse "action.yml" file
	var act action
	if err := yaml.Unmarshal(content, &act); err != nil {
		return fmt.Errorf("failed to parse %q: %v", file, err)
	}

	// create "README.md" file
	outputPath := dir + "/" + outputFile
	readmeFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create %q: %v", outputPath, err)
	}

	// write "README.md" file
	if err := a.writeReadme(readmeFile, act); err != nil {
		return fmt.Errorf("failed to write %q: %v", outputPath, err)
	}

	return nil
}

func (a *app) writeReadme(f *os.File, act action) error {
	// replace placeholders in template
	if err := a.tmpl.Execute(f, act); err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	return nil
}

func val(v interface{}) string {
	return fmt.Sprintf("%q", v)
}
