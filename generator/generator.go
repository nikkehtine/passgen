package generator

const (
	lowercase_chars = "abcdefghijklmnopqrstuvwxyz"
	uppercase_chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers         = "0123456789"
	symbols         = "!@#$%^&*-_=+"
	similar         = "iIl1oO0vw"
)

type GeneratorConfig struct {
	length          int
	lowercase_chars bool
	uppercase_chars bool
	numbers         bool
	symbols         bool
	exclude_similar bool
}

// Generates a password based on provided config
func Generate(config *GeneratorConfig) string {
	return "TODO"
}
