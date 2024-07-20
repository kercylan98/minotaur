package column

import (
	"github.com/kercylan98/minotaur/engine/ecs/storage"
)

func New[PK, Col comparable]() *Storage[PK, Col] {
	return &Storage[PK, Col]{
		primaryKeys: make(map[PK]int),
		columns:     make(map[Col]*column),
	}
}

// Storage 基于列式存储布局的存储器
type Storage[PK, Col comparable] struct {
	columns     map[Col]*column // 列
	primaryKeys map[PK]int      // 主键
	rows        int             // 行数
	invalids    []int           // 无效行
}

// Migrate 迁移行
func (s *Storage[PK, Col]) Migrate(target storage.Storage[PK, Col], primaryKes ...PK) {
	if len(primaryKes) == 0 {
		return
	}

	// 过滤不存在的主键，获得要迁移的行号
	var indexes = make([]int, 0, len(primaryKes))
	var newPrimaryKeys = make([]PK, 0, len(primaryKes))
	var data = make(map[Col][]any)
	for _, primaryKey := range primaryKes {
		index, exists := s.primaryKeys[primaryKey]
		if !exists {
			continue
		}

		delete(s.primaryKeys, primaryKey)
		s.invalids = append(s.invalids, index)
		indexes = append(indexes, index)
		newPrimaryKeys = append(newPrimaryKeys, primaryKey)

		// 生成数据
		for col, c := range s.columns {
			if _, exists := data[col]; !exists {
				data[col] = make([]any, 0, len(primaryKes))
			}
			data[col] = append(data[col], c.data.Get(index))
		}
	}

	// 迁移数据
	target.AddRowsWithValues(newPrimaryKeys, data)
}

// DelRow 删除一行数据
func (s *Storage[PK, Col]) DelRow(primaryKey PK) {
	if index, exists := s.primaryKeys[primaryKey]; exists {
		delete(s.primaryKeys, primaryKey)
		s.invalids = append(s.invalids, index)
	}
}

// DelRows 批量删除数据
func (s *Storage[PK, Col]) DelRows(primaryKeys []PK) {
	for _, primaryKey := range primaryKeys {
		if index, exists := s.primaryKeys[primaryKey]; exists {
			delete(s.primaryKeys, primaryKey)
			s.invalids = append(s.invalids, index)
		}
	}
}

// SetColumn 设置列
func (s *Storage[PK, Col]) SetColumn(col Col, defaultGetter func() any) {
	if _, exists := s.columns[col]; !exists {
		s.columns[col] = newColumn(1024, defaultGetter)
	}
}

// Get 获取特定列数据
func (s *Storage[PK, Col]) Get(primaryKey PK, col Col) any {
	if index, exists := s.primaryKeys[primaryKey]; exists {
		c := s.columns[col]
		v := *c.data.Get(index)
		if v == nil {
			v = c.defaultGetter()
			c.data.Set(index, v)
		}
		return v
	}

	return nil
}

// GetRow 获取一行数据
func (s *Storage[PK, Col]) GetRow(primaryKey PK) map[Col]any {
	var result = make(map[Col]any)

	if index, exists := s.primaryKeys[primaryKey]; exists {
		for col, c := range s.columns {
			v := *c.data.Get(index)
			if v == nil {
				v = c.defaultGetter()
				c.data.Set(index, v)
			}
			result[col] = v
		}
	}

	return result
}

// AddRow 添加一行以默认值填充的数据
func (s *Storage[PK, Col]) AddRow(primaryKey PK) {
	var index int

	if len(s.invalids) > 0 {
		index = s.invalids[0]
		s.invalids = s.invalids[1:]
	} else {
		index = s.rows
		s.rows++
	}
	s.primaryKeys[primaryKey] = index

	// 配置列数据并提前扩容
	for _, c := range s.columns {
		c.data.Grow([]int{index})
	}

}

// AddRows 批量添加数据
func (s *Storage[PK, Col]) AddRows(primaryKeys []PK) {
	var indexes = make([]int, len(primaryKeys))

	// 配置主键且锁行，并且生成默认值
	var index int
	for i, primaryKey := range primaryKeys {
		// 计算索引
		if len(s.invalids) > 0 {
			index = s.invalids[0]
			s.invalids = s.invalids[1:]
		} else {
			index = s.rows
			s.rows++
		}
		indexes[i] = index

		// 设置主键
		s.primaryKeys[primaryKey] = index
	}

	// 扩容列
	for _, c := range s.columns {
		c.data.Grow(indexes)
	}
}

// AddRowsWithValues 批量添加数据
func (s *Storage[PK, Col]) AddRowsWithValues(primaryKeys []PK, values map[Col][]any) {
	var indexes = make([]int, len(primaryKeys))
	var valueGenerators = make(map[Col]func() any)

	// 配置主键
	var index int
	for i, primaryKey := range primaryKeys {
		// 计算索引
		if len(s.invalids) > 0 {
			index = s.invalids[0]
			s.invalids = s.invalids[1:]
		} else {
			index = s.rows
			s.rows++
		}
		indexes[i] = index

		// 设置主键
		s.primaryKeys[primaryKey] = index
	}

	// 配置列数据
	var columnSetter = make([]func(), 0, len(s.columns))
	for col, c := range s.columns {
		col := col
		valueGenerators[col] = c.defaultGetter
		c.data.Grow(indexes)
		columnSetter = append(columnSetter, func() {
			// 行锁内生成数据
			data, exists := values[col]
			if !exists {
				return // 不存在默认 nil
			}
			c.data.BatchSet(indexes, data)
		})
	}

	for _, f := range columnSetter {
		f()
	}
}
