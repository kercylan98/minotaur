package vivid

type Client interface {
	// Exec 执行远程调用
	Exec(data []byte) error
}
