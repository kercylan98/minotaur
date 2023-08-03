package arrangement

// Option 编排选项
type Option[ID comparable, AreaInfo any] func(arrangement *Arrangement[ID, AreaInfo])
