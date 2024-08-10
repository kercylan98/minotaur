package storehouse

type Storage interface {
	CreateTable(name string) error
}
