package main

import (
	"fmt"
)

const maxX = 64
const maxY = 48

// Grid that stores the actors for the program
type Grid struct {
	cells [maxX][maxY]Actor
}

func (g *Grid) append(a Actor, x, y int) {
	if x >= maxX || x < 0 || y >= maxY || y < 0 {
		return
	}
	g.cells[x][y] = a
}

func (g *Grid) move(a Actor, x, y int) bool {
	if x >= maxX || x < 0 || y >= maxY || y < 0 {
		return false
	}

	if oX, oY := a.getX(), a.getY(); g.cells[oX][oY] != nil {
		if g.cells[x][y] == nil {
			// Move that goober
			g.cells[x][y] = a
			// Delete the source cell
			g.cells[oX][oY] = nil
			a.setX(x)
			a.setY(y)
			return true
		}

		fmt.Printf("boop\n")
	} else {
		fmt.Printf("beep\n")
	}
	return false
}

func (g *Grid) actors() []Actor {
	var retValues []Actor
	for i := range g.cells {
		for j := range g.cells[i] {
			if g.cells[i][j] != nil {
				retValues = append(retValues, g.cells[i][j])
			}
		}
	}
	return retValues
}

func (g *Grid) neighbors(x int, y int) []Actor {
	var retValues []Actor
	for i := x - 1; i <= x+1; i++ {
		if i >= maxX || i < 0 {
			continue
		}
		for j := y - 1; j <= y+1; j++ {
			if j >= maxY || j < 0 {
				continue
			}
			if g.cells[i][j] != nil {
				retValues = append(retValues, g.cells[i][j])
			}
		}
	}

	return retValues
}

func (g *Grid) getCell(x int, y int) (Actor, bool) {
	if x >= maxX || x < 0 || y >= maxY || y < 0 {
		return nil, false
	}
	return g.cells[x][y], true
}

func (g *Grid) remove(x, y int) bool {
	if x < maxX || x >= 0 || y < maxY || y >= 0 {
		a := g.cells[x][y]
		g.cells[x][y] = nil
		return !(a == nil)
	}
	return false
}
