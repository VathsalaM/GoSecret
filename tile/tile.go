package src

import "github.com/VathsalaM/GoSecret/Interface"

type tile struct{
	id int
	color string
}

func NewTile(id int,color string)(newTile tile){
	newTile.color = color
	newTile.id = id
	return
}

func (t *tile)Place(element Interface.Element){

}