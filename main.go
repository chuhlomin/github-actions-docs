package main

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
)

type config struct {
	WorkingDir []string `env:"INPUT_WORKING_DIR" envDefault:"." envSeparator:","`
	OutputFile string   `env:"INPUT_OUTPUT_FILE" envDefault:"README.md"`
}

func main() {
	log.Println("Starting...")

	if err := run(); err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	log.Println("Done.")
}

func run() error {
	log.Println("Parsing environment variables...")
	var c config
	if err := env.Parse(&c); err != nil {
		return errors.Wrap(err, "failed to parse environment variables")
	}

	app, err := newApp("template.md")
	if err != nil {
		return errors.Wrap(err, "failed to create app")
	}

	var numChanged uint
	for _, dir := range c.WorkingDir {
		log.Printf("Processing directory: %s", dir)
		if err := app.ProcessDir(dir, c.OutputFile); err != nil {
			return errors.Wrap(err, "failed to process directory")
		}
		numChanged++
	}

	if numChanged == 0 {
		log.Println("No changes detected.")
		return nil
	}

	fmt.Printf("::set-output name=num-changed::%d", numChanged)

	return nil
}
