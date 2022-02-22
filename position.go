package main

type Position struct {
	x, y int
}

func (p *Position) Transform() {

}

func (p *Position) Up() {
	p.x -= 1
}

func (p *Position) Down() {
	p.x += 1
}

func (p *Position) Left() {
	p.y -= 1
}

func (p *Position) Right() {
	p.y += 1
}

func (p *Position) Add(pos Position) {
	p.x += pos.x
	p.y += pos.y
}

func (p *Position) Sub(pos Position) {
	if p.x > 0 && p.y > 0 {
		p.x -= pos.x
		p.y -= pos.y
	}
}
