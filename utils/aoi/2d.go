package aoi

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/geometry"
	"github.com/kercylan98/minotaur/utils/hash"
	"math"
	"sync"
)

func NewTwoDimensional[EID generic.Basic, PosType generic.SignedNumber, E TwoDimensionalEntity[EID, PosType]](width, height, areaWidth, areaHeight int) *TwoDimensional[EID, PosType, E] {
	aoi := &TwoDimensional[EID, PosType, E]{
		event:  new(event[EID, PosType, E]),
		width:  float64(width),
		height: float64(height),
		focus:  map[EID]map[EID]E{},
	}
	aoi.SetAreaSize(areaWidth, areaHeight)
	return aoi
}

type TwoDimensional[EID generic.Basic, PosType generic.SignedNumber, E TwoDimensionalEntity[EID, PosType]] struct {
	*event[EID, PosType, E]
	rw               sync.RWMutex
	width            float64
	height           float64
	areaWidth        float64
	areaHeight       float64
	areaWidthLimit   int
	areaHeightLimit  int
	areas            [][]map[EID]E
	focus            map[EID]map[EID]E
	repartitionQueue []func()
}

func (slf *TwoDimensional[EID, PosType, E]) AddEntity(entity E) {
	slf.rw.Lock()
	slf.addEntity(entity)
	slf.rw.Unlock()
}

func (slf *TwoDimensional[EID, PosType, E]) DeleteEntity(entity E) {
	slf.rw.Lock()
	slf.deleteEntity(entity)
	slf.rw.Unlock()
}

func (slf *TwoDimensional[EID, PosType, E]) Refresh(entity E) {
	slf.rw.Lock()
	defer slf.rw.Unlock()
	slf.refresh(entity)
}

func (slf *TwoDimensional[EID, PosType, E]) GetFocus(id EID) map[EID]E {
	slf.rw.RLock()
	defer slf.rw.RUnlock()
	return hash.Copy(slf.focus[id])
}

func (slf *TwoDimensional[EID, PosType, E]) SetSize(width, height int) {
	fw, fh := float64(width), float64(height)
	if fw == slf.width && fh == slf.height {
		return
	}
	slf.rw.Lock()
	defer slf.rw.Unlock()
	slf.width = fw
	slf.height = fh
	slf.setAreaSize(int(slf.areaWidth), int(slf.areaHeight))
}

func (slf *TwoDimensional[EID, PosType, E]) SetAreaSize(width, height int) {
	fw, fh := float64(width), float64(height)
	if fw == slf.areaWidth && fh == slf.areaHeight {
		return
	}
	slf.rw.Lock()
	defer slf.rw.Unlock()
	slf.setAreaSize(width, height)
}

func (slf *TwoDimensional[EID, PosType, E]) setAreaSize(width, height int) {

	// 旧分区备份
	var oldAreas = make([][]map[EID]E, len(slf.areas))
	for w := 0; w < len(slf.areas); w++ {
		hs := slf.areas[w]
		ohs := make([]map[EID]E, len(hs))
		for h := 0; h < len(hs); h++ {
			es := map[EID]E{}
			for g, e := range hs[h] {
				es[g] = e
			}
			ohs[h] = es
		}
		oldAreas[w] = ohs
	}

	// 清理分区
	for i := 0; i < len(oldAreas); i++ {
		area := slf.areas[i]
		for a := 0; a < len(area); a++ {
			entities := area[a]
			for _, entity := range entities {
				slf.deleteEntity(entity)
			}
		}
	}

	// 生成区域
	slf.areaWidth = float64(width)
	slf.areaHeight = float64(height)
	slf.areaWidthLimit = int(math.Ceil(slf.width / slf.areaWidth))
	slf.areaHeightLimit = int(math.Ceil(slf.height / slf.areaHeight))
	areas := make([][]map[EID]E, slf.areaWidthLimit+1)
	for i := 0; i < len(areas); i++ {
		entities := make([]map[EID]E, slf.areaHeightLimit+1)
		for e := 0; e < len(entities); e++ {
			entities[e] = map[EID]E{}
		}
		areas[i] = entities
	}
	slf.areas = areas

	// 重新分区
	for i := 0; i < len(oldAreas); i++ {
		area := oldAreas[i]
		for a := 0; a < len(area); a++ {
			entities := area[a]
			for _, entity := range entities {
				slf.addEntity(entity)
			}
		}
	}
}

func (slf *TwoDimensional[EID, PosType, E]) addEntity(entity E) {
	x, y := entity.GetPosition().GetXY()
	widthArea := int(float64(x) / slf.areaWidth)
	heightArea := int(float64(y) / slf.areaHeight)
	id := entity.GetTwoDimensionalEntityID()
	slf.areas[widthArea][heightArea][id] = entity
	focus := map[EID]E{}
	slf.focus[id] = focus
	slf.rangeVisionAreaEntities(entity, func(eg EID, e E) {
		focus[eg] = e
		slf.OnEntityJoinVisionEvent(entity, e)
		slf.refresh(e)
	})
}

func (slf *TwoDimensional[EID, PosType, E]) refresh(entity E) {
	x, y := entity.GetPosition().GetXY()
	vision := entity.GetVision()
	id := entity.GetTwoDimensionalEntityID()
	focus := slf.focus[id]
	for eg, e := range focus {
		ex, ey := e.GetPosition().GetXY()
		if geometry.CalcDistanceWithCoordinate(float64(x), float64(y), float64(ex), float64(ey)) > vision {
			delete(focus, eg)
			delete(slf.focus[eg], id)
		}
	}

	slf.rangeVisionAreaEntities(entity, func(id EID, e E) {
		if _, exist := focus[id]; !exist {
			focus[id] = e
			slf.OnEntityJoinVisionEvent(entity, e)
		}
	})
}

func (slf *TwoDimensional[EID, PosType, E]) rangeVisionAreaEntities(entity E, handle func(id EID, entity E)) {
	x, y := entity.GetPosition().GetXY()
	widthArea := int(float64(x) / slf.areaWidth)
	heightArea := int(float64(y) / slf.areaHeight)
	vision := entity.GetVision()
	widthSpan := int(math.Ceil(vision / slf.areaWidth))
	heightSpan := int(math.Ceil(vision / slf.areaHeight))
	id := entity.GetTwoDimensionalEntityID()

	sw := widthArea - widthSpan
	if sw < 0 {
		sw = 0
	} else if sw > slf.areaWidthLimit {
		sw = slf.areaWidthLimit
	}
	ew := widthArea - widthSpan
	if ew < sw {
		ew = sw
	} else if ew > slf.areaWidthLimit {
		ew = slf.areaWidthLimit
	}
	for w := sw; w < ew; w++ {
		sh := heightArea - heightSpan
		if sh < 0 {
			sh = 0
		} else if sh > slf.areaHeightLimit {
			sh = slf.areaHeightLimit
		}
		eh := widthArea - widthSpan
		if eh < sh {
			eh = sh
		} else if eh > slf.areaHeightLimit {
			eh = slf.areaHeightLimit
		}
		for h := sh; h < eh; h++ {
			var areaX, areaY float64
			if w < widthArea {
				tempW := w + 1
				areaX = float64(tempW * int(slf.areaWidth))
			} else if w > widthArea {
				areaX = float64(w * int(slf.areaWidth))
			} else {
				areaX = float64(x)
			}
			if h < heightArea {
				tempH := h + 1
				areaY = float64(tempH * int(slf.areaHeight))
			} else if h > heightArea {
				areaY = float64(h * int(slf.areaHeight))
			} else {
				areaY = float64(y)
			}
			areaDistance := geometry.CalcDistanceWithCoordinate(float64(x), float64(y), areaX, areaY)
			if areaDistance <= vision {
				for eg, e := range slf.areas[w][h] {
					if eg == id {
						continue
					}
					if ex, ey := e.GetPosition().GetXY(); geometry.CalcDistanceWithCoordinate(float64(x), float64(y), float64(ex), float64(ey)) > vision {
						continue
					}
					handle(eg, e)
				}
			}
		}
	}
}

func (slf *TwoDimensional[EID, PosType, E]) deleteEntity(entity E) {
	x, y := entity.GetPosition().GetXY()
	widthArea := int(float64(x) / slf.areaWidth)
	heightArea := int(float64(y) / slf.areaHeight)
	id := entity.GetTwoDimensionalEntityID()
	focus := slf.focus[id]
	for g, e := range focus {
		slf.OnEntityLeaveVisionEvent(entity, e)
		slf.OnEntityLeaveVisionEvent(e, entity)
		delete(slf.focus[g], id)
	}
	delete(slf.focus, id)
	delete(slf.areas[widthArea][heightArea], id)
}
