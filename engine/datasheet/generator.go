package datasheet

type CodeGenerator interface {
	Generate(set *Set) error
}
