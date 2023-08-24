package storage

// Warehouse 数据仓库接口，用于数据生产及持久化
type Warehouse[K comparable, D any] interface {
	// GenerateZero 生成一个空的数据对象
	GenerateZero() D

	// Init 用于数据初始化，例如启动时从数据库中加载所有玩家的离线数据等情况
	Init() (map[K][]byte, error)

	// Query 查询特定 key 的数据
	//  - 返回的 err 应该排除数据不存在的错误情况，例如：sql.ErrNoRows
	Query(key K) (data []byte, err error)

	// Create 创建特定 key 的数据
	Create(key K, data []byte) error

	// Save 保存特定 key 的数据
	Save(key K, data []byte) error
}
