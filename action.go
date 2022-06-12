package main

type branding struct {
	Icon  string `yaml:"icon"`
	Color string `yaml:"color"`
}

type input struct {
	Description string `yaml:"description"`
	Required    bool   `yaml:"required"`
	Default     string `yaml:"default"`
}

type output struct {
	Description string `yaml:"description"`
}

type runs struct {
	Using string `yaml:"using"`
	Image string `yaml:"image"`
}

type action struct {
	Name        string            `yaml:"name"`
	Description string            `yaml:"description"`
	Author      string            `yaml:"author"`
	Branding    branding          `yaml:"branding"`
	Inputs      map[string]input  `yaml:"inputs"`
	Outputs     map[string]output `yaml:"outputs"`
	Runs        runs              `yaml:"runs"`
}
