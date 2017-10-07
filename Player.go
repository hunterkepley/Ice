package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	"math"
)

type Player struct {
	pos    pixel.Vec
	dir    string
	size   pixel.Vec
	tileID int
	pic    pixel.Picture
	sprite *pixel.Sprite
	moving bool
}

func newPlayer(tID int, loc string) Player { // Constructor for Player
	pic, err := loadPicture(loc)
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(pic, pic.Bounds())
	return Player{pixel.ZV, "right", pixel.V(50, 50), tID, pic, sprite, false}
	// Returns a basic Player
}

func (p *Player) update(tiles []Tile) { // Updates a tile
	if !p.moving {
		p.pos = pixel.V(tiles[p.tileID].pos.X+float64(TileSize/2), tiles[p.tileID].pos.Y+float64(TileSize/2))
		// TileSize/2 bc the x,y is at the center of the picture
	}
}

func (p Player) render(win *pixelgl.Window) { // Draws a tile
	mat := pixel.IM

	switch p.dir {
	case "up":
		mat = mat.Rotated(pixel.ZV, 0)
	case "down":
		mat = mat.Rotated(pixel.ZV, math.Pi)
	case "left":
		mat = mat.Rotated(pixel.ZV, math.Pi/2)
	case "right":
		mat = mat.Rotated(pixel.ZV, (3*math.Pi)/2)
	default:
		mat = mat.Rotated(pixel.ZV, 0)
	}

	mat = mat.Moved(p.pos)
	p.sprite.Draw(win, mat)
}
