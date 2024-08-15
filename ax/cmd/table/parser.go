package table

type TypeParser interface {
	Parse(input string) []Type
}

type DataParser interface {
	Parse(input string) string
}

type CodeParser interface {
	Parse(configs []*Config) string
}
