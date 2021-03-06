package components

import "github.com/VathsalaM/GoSecret/Interface"

type horse struct{
	position Interface.Position
	id int
	colour string
}

func NewHorse(id int,colour string)(newHorse horse){
	newHorse.colour = colour
	newHorse.id = id
	newHorse.position= Interface.Position{}
	return
}

func (h *horse)UpdatePosition(position Interface.Position){
	h.position = position
}

func (h *horse)Position()Interface.Position{
	return h.position
}

