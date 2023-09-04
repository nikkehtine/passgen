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

// Generates a password based on provided config
func Generate(config *GeneratorConfig) string {
	return "TODO"
}
