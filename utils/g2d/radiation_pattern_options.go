package g2d

type RadiationPatternOption[ItemType comparable, Item RadiationPatternItem[ItemType]] func(rp *RadiationPattern[ItemType, Item])

func WithRadiationPatternExclude[ItemType comparable, Item RadiationPatternItem[ItemType]](itemType ...ItemType) RadiationPatternOption[ItemType, Item] {
	return func(rp *RadiationPattern[ItemType, Item]) {
		if rp.excludes == nil {
			rp.excludes = map[ItemType]bool{}
		}
		for _, t := range itemType {
			rp.excludes[t] = true
		}
	}
}
