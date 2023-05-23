package astar

import (
	"container/heap"
	"fmt"
	"math"
)

// Pathfinding3D 寻路
//   - grid: X、Y、Z
func Pathfinding3D(grid [][][]int, start, target Point3D, collisionRange int) ([]Point3D, error) {
	path, err := aStar3D(grid, start, target, collisionRange)
	if err != nil {
		return nil, err
	}
	return path, nil
}

func aStar3D(grid [][][]int, start, goal Point3D, collisionRange int) ([]Point3D, error) {
	startNode := &Node3D{point: start, g: 0, h: heuristic(start, goal)}
	startNode.f = startNode.g + startNode.h
	openList := make(Node3DList, 0)
	heap.Init(&openList)
	heap.Push(&openList, startNode)
	closedList := make(map[Point3D]*Node3D)

	for openList.Len() > 0 {
		current := heap.Pop(&openList).(*Node3D)
		if current.point == goal {
			path := make([]Point3D, 0)
			for current != nil {
				path = append([]Point3D{current.point}, path...)
				current = current.parent
			}
			return path, nil
		}

		closedList[current.point] = current
		for _, neighborPoint := range neighbors(current.point, grid, collisionRange) {
			if _, ok := closedList[neighborPoint]; ok {
				continue
			}
			tentativeG := current.g + 1
			neighborNode := &Node3D{point: neighborPoint, parent: current, g: tentativeG, h: heuristic(neighborPoint, goal)}
			neighborNode.f = neighborNode.g + neighborNode.h
			for _, openNode := range openList {
				if openNode.point == neighborPoint && tentativeG >= openNode.g {
					continue
				}
			}
			heap.Push(&openList, neighborNode)
		}
	}
	return nil, fmt.Errorf("no path found")
}

func heuristic(a, b Point3D) float64 {
	return math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)) + math.Abs(float64(a.Z-b.Z))
}

func inBounds(point Point3D, grid [][][]int) bool {
	return point.X >= 0 && point.X < len(grid) && point.Y >= 0 && point.Y < len(grid[0]) && point.Z >= 0 && point.Z < len(grid[0][0])
}

func passable(point Point3D, grid [][][]int, collisionRange int) bool {
	for x := -collisionRange; x <= collisionRange; x++ {
		for y := -collisionRange; y <= collisionRange; y++ {
			for z := -collisionRange; z <= collisionRange; z++ {
				newPoint := Point3D{point.X + x, point.Y + y, point.Z + z}
				if inBounds(newPoint, grid) && grid[newPoint.X][newPoint.Y][newPoint.Z] == 1 {
					return false
				}
			}
		}
	}
	return true
}

func neighbors(point Point3D, grid [][][]int, collisionRange int) []Point3D {
	neighborPoints := []Point3D{
		{point.X - 1, point.Y, point.Z}, {point.X + 1, point.Y, point.Z},
		{point.X, point.Y - 1, point.Z}, {point.X, point.Y + 1, point.Z},
		{point.X, point.Y, point.Z - 1}, {point.X, point.Y, point.Z + 1},
	}
	result := make([]Point3D, 0)
	for _, neighbor := range neighborPoints {
		if inBounds(neighbor, grid) && passable(neighbor, grid, collisionRange) {
			result = append(result, neighbor)
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
