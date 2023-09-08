package hash

// Chunk 将哈希表按照指定大小分块
//   - m: 待分块的哈希表
//   - size: 每块的大小
func Chunk[K comparable, V any](m map[K]V, size int) []map[K]V {
	if len(m) == 0 {
		return nil
	}

	var (
		i int
		j int
	)
	chunks := make([]map[K]V, (len(m)-1)/size+1)
	for i = 0; i < len(m); i += size {
		chunks[j] = make(map[K]V, size)
		for key, value := range m {
			if i <= j*size && j*size < i+size {
				chunks[j][key] = value
			}
		}
		j++
	}

	return chunks
}
