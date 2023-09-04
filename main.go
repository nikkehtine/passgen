package main

import (
	"fmt"

	"gitlab.com/nikkehtine/passgen/checker"
	"gitlab.com/nikkehtine/passgen/generator"
)

func main() {
	config := new(generator.GeneratorConfig)
	fmt.Printf("Generator: %s\n", generator.Generate(config))
	fmt.Printf("Checker: %s\n", checker.Check("TODO"))
}
