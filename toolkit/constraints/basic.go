package constraints

type Basic interface {
	Ordered
	bool | []byte | rune | byte
	~bool | ~[]byte | ~rune | ~byte
}
