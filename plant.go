package main

import (
	"math/rand"

	"github.com/gotk3/gotk3/cairo"
)

var directionModifiers = [][]int{
	[]int{-1, 1},
	[]int{0, 1},
	[]int{1, 1},
	[]int{-1, 0},
	[]int{1, 0},
	[]int{-1, -1},
	[]int{0, -1},
	[]int{1, -1},
}

// Plant represents a single plant for our purposes
type Plant struct {
	x, y, energy int
	color        RGB
	altruistic   bool
}

func (p *Plant) act(g *Grid) bool {
	// Age
	if threshhold := rand.Intn(100); threshhold < 5 {
		direction := rand.Intn(8)
		cX := p.x - directionModifiers[direction][0]
		cY := p.y - directionModifiers[direction][1]
		if a, inbounds := g.getCell(cX, cY); inbounds {
			if a == nil {
				if p.energy > 100 {
					newPlant := p.split(cX, cY)
					g.append(&newPlant, newPlant.x, newPlant.y)
				}
			} else {
				if p.altruistic && p.energy > a.getEnergy() {
					split := (p.energy + a.getEnergy()) / 2
					p.energy = split
					a.setEnergy(split)
				} else {
					p.energy -= 5
				}
			}
		}
		return true
	} else if threshhold < 60 && p.energy < 121 {
		p.energy++
	} else if p.energy == 0 {
		g.remove(p.x, p.y)
	} else {
		p.energy--
	}
	return false
}

func (p *Plant) setEnergy(newEnergy int) {
	p.energy = newEnergy
}

func (p *Plant) getEnergy() int {
	return p.energy
}

func (p *Plant) split(x, y int) Plant {
	// Check for mutation
	r, g, b, a := p.color.r, p.color.g, p.color.b, p.altruistic

	if threshhold := rand.Intn(400); threshhold < 1 {
		flip := rand.Intn(7)
		switch flip {
		case 0:
			x := r
			r = b
			b = x
		case 1:
			x := r
			r = g
			g = x
		case 2:
			x := b
			b = g
			g = x
		case 3:
			x := b
			b = r
			r = x
		case 4:
			x := g
			g = r
			r = x
		case 5:
			x := g
			g = b
			b = x
		case 6:
			a = !a
		}
	}
	newPlant := Plant{x, y, (p.energy / 2) - 5, RGB{r, g, b}, a}
	p.energy = p.energy/2 - 5
	return newPlant
}

func (p *Plant) getX() int {
	return p.x
}

func (p *Plant) getY() int {
	return p.y
}

func (p *Plant) setX(cX int) {
	p.x = cX
}

func (p *Plant) setY(cY int) {
	p.y = cY
}

func (p *Plant) create() {
	p.x, p.y = rand.Intn(63), rand.Intn(47)
	p.energy = 50
	p.color = RGB{rand.Float64(), rand.Float64(), rand.Float64()}
	p.altruistic = rand.Intn(2) == 1
}

func (p *Plant) draw(unitsize float64, cr *cairo.Context) {
	cr.SetSourceRGBA(p.color.r, p.color.g, p.color.b, p.getBrightness())
	x, y := float64(p.x)*unitsize, float64(p.y)*unitsize
	cr.Rectangle(x, y, unitsize, unitsize)
	cr.Fill()
}

func (p *Plant) getBrightness() float64 {
	var retValue float64
	switch {
	case p.energy <= 25:
		retValue = .25
	case p.energy <= 50:
		retValue = .50
	case p.energy <= 75:
		retValue = .75
	default:
		retValue = 1.0
	}
	return retValue
}
