package main

import (
	"fmt"

	"gitlab.com/nikkehtine/passgen/generator"
)

func main() {
	config := new(generator.GeneratorConfig)
	fmt.Println(generator.Generate(config))
}
