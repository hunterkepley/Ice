package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	"math"
)

type Player struct {
	pos     pixel.Vec
	dir     string
	size    pixel.Vec
	tileID  int
	pic     pixel.Picture
	sprite  *pixel.Sprite
	moving  bool
	canMove bool
	ID      int // 0 == player 1 [wasd], 1 == player 2 [arrow keys]
}

func newPlayer(tID int, loc string, ID int) Player { // Constructor for Player
	pic, err := loadPicture(loc)
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(pic, pic.Bounds())
	if ID == 0 {
		return Player{pixel.ZV, "right", pixel.V(50, 50), tID, pic, sprite, false, true, ID}
	} else if ID == 1 {
		return Player{pixel.ZV, "left", pixel.V(50, 50), tID, pic, sprite, false, true, ID}
	}

	return Player{pixel.ZV, "right", pixel.V(50, 50), tID, pic, sprite, false, true, ID}

	// Returns a basic Player
}

func (p *Player) update(tile *Tile, win *pixelgl.Window) { // Updates a tile
	if !p.moving {
		p.pos = pixel.V(tile.pos.X+float64(TileSize/2), tile.pos.Y+float64(TileSize/2))
		tile.state = p.ID + 1
		// TileSize/2 bc the x,y is at the center of the picture
	} else {
	}

	if p.canMove {
		if p.ID == 0 {
			if win.Pressed(pixelgl.KeyA) {
				p.dir = "left"
			}
			if win.Pressed(pixelgl.KeyD) {
				p.dir = "right"
			}
			if win.Pressed(pixelgl.KeyS) {
				p.dir = "down"
			}
			if win.Pressed(pixelgl.KeyW) {
				p.dir = "up"
			}
		} else if p.ID == 1 {
			if win.Pressed(pixelgl.KeyLeft) {
				p.dir = "left"
			}
			if win.Pressed(pixelgl.KeyRight) {
				p.dir = "right"
			}
			if win.Pressed(pixelgl.KeyDown) {
				p.dir = "down"
			}
			if win.Pressed(pixelgl.KeyUp) {
				p.dir = "up"
			}
		}
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
