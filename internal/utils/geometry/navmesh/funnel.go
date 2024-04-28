package navmesh

import (
	"github.com/kercylan98/minotaur/utils/generic"
	"github.com/kercylan98/minotaur/utils/geometry"
)

type funnel[V generic.SignedNumber] struct {
	path    []geometry.Point[V]
	portals [][2]geometry.Point[V]
}

func (slf *funnel[V]) push(point1, point2 geometry.Point[V]) {
	slf.portals = append(slf.portals, [2]geometry.Point[V]{point1, point2})
}

func (slf *funnel[V]) pushSingle(point geometry.Point[V]) {
	slf.portals = append(slf.portals, [2]geometry.Point[V]{point, point})
}

func (slf *funnel[V]) stringPull() []geometry.Point[V] {
	var (
		portals     = slf.portals
		points      []geometry.Point[V]
		apexIndex   = 0
		leftIndex   = 0
		rightIndex  = 0
		portalApex  = portals[0][0]
		portalLeft  = portals[0][0]
		portalRight = portals[0][1]
	)

	points = append(points, portalApex)
	for i := 1; i < len(portals); i++ {
		lr := portals[i]
		left, right := lr[0], lr[1]

		if geometry.CalcTriangleTwiceArea(portalApex, portalRight, right) <= V(0) {
			if portalApex.Equal(portalRight) || geometry.CalcTriangleTwiceArea(portalApex, portalLeft, right) > V(0) {
				portalRight = right
				rightIndex = i
			} else {
				points = append(points, portalLeft)
				portalApex = portalLeft
				apexIndex = leftIndex

				portalLeft = portalApex
				portalRight = portalApex
				leftIndex = apexIndex
				rightIndex = apexIndex

				i = apexIndex
				continue
			}
		}

		if geometry.CalcTriangleTwiceArea(portalApex, portalLeft, left) >= V(0) {
			if portalApex.Equal(portalLeft) || geometry.CalcTriangleTwiceArea(portalApex, portalRight, left) < V(0) {
				portalLeft = left
				leftIndex = i
			} else {
				points = append(points, portalRight)
				portalApex = portalRight
				apexIndex = rightIndex

				portalLeft = portalApex
				portalRight = portalApex
				leftIndex = apexIndex
				rightIndex = apexIndex

				i = apexIndex
				continue
			}
		}
	}

	if len(points) == 0 || !points[len(points)-1].Equal(portals[len(portals)-1][0]) {
		points = append(points, portals[len(portals)-1][0])
	}

	slf.path = points
	return slf.path
}
