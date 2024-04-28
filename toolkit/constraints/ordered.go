package constraints

// Ordered 可排序类型
type Ordered interface {
	Int | Float | ~string
}
