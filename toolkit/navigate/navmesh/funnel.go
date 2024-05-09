package navmesh

import (
	"github.com/kercylan98/minotaur/toolkit/geometry"
)

type funnel struct {
	path    []geometry.Vector2
	portals [][2]geometry.Vector2
}

func (slf *funnel) push(point1, point2 geometry.Vector2) {
	slf.portals = append(slf.portals, [2]geometry.Vector2{point1, point2})
}

func (slf *funnel) pushSingle(point geometry.Vector2) {
	slf.portals = append(slf.portals, [2]geometry.Vector2{point, point})
}

func (slf *funnel) stringPull() []geometry.Vector2 {
	var (
		portals     = slf.portals
		points      []geometry.Vector2
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

		if geometry.CalcTriangleAreaTwice(portalApex, portalRight, right) <= 0 {
			if portalApex.Equal(portalRight) || geometry.CalcTriangleAreaTwice(portalApex, portalLeft, right) > 0 {
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

		if geometry.CalcTriangleAreaTwice(portalApex, portalLeft, left) >= 0 {
			if portalApex.Equal(portalLeft) || geometry.CalcTriangleAreaTwice(portalApex, portalRight, left) < 0 {
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
