package astar

import (
	"fmt"
	"testing"
)

func TestPathfinding3D(t *testing.T) {
	grid := [][][]int{
		{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
		{
			{0, 0, 0, 0, 0},
			{0, 1, 1, 1, 0},
			{0, 1, 0, 1, 0},
			{0, 1, 1, 1, 0},
			{0, 0, 0, 0, 0},
		},
		{
			{1, 0, 0, 0, 0},
			{0, 1, 0, 1, 0},
			{0, 0, 0, 0, 0},
			{0, 1, 0, 1, 0},
			{0, 0, 0, 0, 0},
		},
		{
			{0, 0, 0, 0, 0},
			{0, 1, 1, 1, 0},
			{0, 1, 0, 1, 0},
			{0, 1, 1, 1, 0},
			{0, 0, 0, 0, 0},
		},
		{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
	}

	start := Point3D{2, 1, 0}
	target := Point3D{2, 2, 2}
	collisionRange := 0

	path, err := Pathfinding3D(grid, start, target, collisionRange)
	if err != nil {
		panic(err)
	}

	fmt.Println(path)
	printGrid(grid, start)
}

func printGrid(grid [][][]int, start Point3D) {
	for z := 0; z < len(grid); z++ {
		fmt.Printf("z = %d\n", z)
		for y := 0; y < len(grid[z]); y++ {
			for x := 0; x < len(grid[z][y]); x++ {
				if x == start.X && y == start.Y && z == start.Z {
					fmt.Print("S ")
				} else if grid[z][y][x] == 0 {
					fmt.Print(". ")
				} else {
					fmt.Print("# ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
