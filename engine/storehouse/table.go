package storehouse

type Table interface {
	Fields() []Field
}
