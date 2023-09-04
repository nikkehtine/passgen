package generator

type GeneratorConfig struct {
	length            int
	lowercase_chars   bool
	uppercase_chars   bool
	numbers           bool
	symbols           bool
	exclude_similar   bool
	exclude_ambiguous bool
}

func Generate(config *GeneratorConfig) string {
	return "TODO"
}
