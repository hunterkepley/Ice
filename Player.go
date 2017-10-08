package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	"math"
)

const (
	PlayerSpeed = 250
)

type Player struct {
	pos          pixel.Vec
	dir          string
	size         pixel.Vec
	tileID       int
	pic          pixel.Picture
	sprite       *pixel.Sprite
	moving       bool
	canMove      bool
	ID           int // 0 == player 1 [wasd], 1 == player 2 [arrow keys]
	nextTileID   int
	distanceLeft float64
}

func newPlayer(pos pixel.Vec, tID int, loc string, ID int) Player { // Constructor for Player
	pic, err := loadPicture(loc)
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(pic, pic.Bounds())
	if ID == 0 {
		return Player{pos, "right", pixel.V(50, 50), tID, pic, sprite, false, true, ID, 0, 50}
	} else if ID == 1 {
		return Player{pos, "left", pixel.V(50, 50), tID, pic, sprite, false, true, ID, 0, 50}
	}

	return Player{pos, "right", pixel.V(50, 50), tID, pic, sprite, false, true, ID, 0, 50}

	// Returns a basic Player
}

func (p *Player) update(tile *Tile, win *pixelgl.Window, dt float64) { // Updates a tile
	if !p.moving {
		tile.state = p.ID + 1
		// TileSize/2 bc the x,y is at the center of the picture
	} else {
	}
	if p.canMove {
		if !p.moving {
			if p.ID == 0 {
				if win.Pressed(pixelgl.KeyA) {
					p.dir = "left"
					if p.tileID >= 0 && p.tileID <= (TilesY-1) {
						p.nextTileID = (TilesX * TilesY) - (12 - p.tileID)
					} else {
						p.nextTileID = p.tileID - TilesY
					}
					p.moving = true
				}
				if win.Pressed(pixelgl.KeyD) {
					p.dir = "right"
					if p.tileID >= (TilesX*TilesY)-12 && p.tileID <= (TilesX*TilesY)-1 {
						p.nextTileID = 12 - ((TilesX * TilesY) - p.tileID)
					} else {
						p.nextTileID = p.tileID + TilesY
					}
					p.moving = true
				}
				if win.Pressed(pixelgl.KeyS) {
					p.dir = "down"
					if p.tileID == 0 || p.tileID%12 == 0 {
						p.nextTileID = p.tileID + 11
					} else {
						p.nextTileID = p.tileID - 1
					}
					p.moving = true
				}
				if win.Pressed(pixelgl.KeyW) {
					p.dir = "up"
					if (p.tileID%12)-11 == 0 {
						p.nextTileID = p.tileID - 11
					} else {
						p.nextTileID = p.tileID + 1
					}
					p.moving = true
				}
			} else if p.ID == 1 {
				if win.Pressed(pixelgl.KeyLeft) {
					p.dir = "left"
					if p.tileID >= 0 && p.tileID <= (TilesY-1) {
						p.nextTileID = (TilesX * TilesY) - (12 - p.tileID)
					} else {
						p.nextTileID = p.tileID - TilesY
					}
					p.moving = true
				}
				if win.Pressed(pixelgl.KeyRight) {
					p.dir = "right"
					if p.tileID >= (TilesX*TilesY)-12 && p.tileID <= (TilesX*TilesY)-1 {
						p.nextTileID = 12 - ((TilesX * TilesY) - p.tileID)
					} else {
						p.nextTileID = p.tileID + TilesY
					}
					p.moving = true
				}
				if win.Pressed(pixelgl.KeyDown) {
					p.dir = "down"
					if p.tileID == 0 || p.tileID%12 == 0 {
						p.nextTileID = p.tileID + 11
					} else {
						p.nextTileID = p.tileID - 1
					}
					p.moving = true
				}
				if win.Pressed(pixelgl.KeyUp) {
					p.dir = "up"
					if (p.tileID%12)-11 == 0 {
						p.nextTileID = p.tileID - 11
					} else {
						p.nextTileID = p.tileID + 1
					}
					p.moving = true
				}
			}
		} else {
			switch p.dir {
			case "up":
				if p.distanceLeft > 0 {
					p.pos = pixel.V(p.pos.X, p.pos.Y+(PlayerSpeed*dt))
					p.distanceLeft -= float64(PlayerSpeed * dt)
				} else {
					p.tileID = p.nextTileID
					p.pos = pixel.V(tiles[p.tileID].pos.X+float64(TileSize/2), tiles[p.tileID].pos.Y+float64(TileSize/2))
					p.distanceLeft = 50
					p.moving = false
				}
			case "down":
				if p.distanceLeft > 0 {
					p.pos = pixel.V(p.pos.X, p.pos.Y-(PlayerSpeed*dt))
					p.distanceLeft -= float64(PlayerSpeed * dt)
				} else {
					p.tileID = p.nextTileID
					p.pos = pixel.V(tiles[p.tileID].pos.X+float64(TileSize/2), tiles[p.tileID].pos.Y+float64(TileSize/2))
					p.distanceLeft = 50
					p.moving = false
				}
			case "left":
				if p.distanceLeft > 0 {
					p.pos = pixel.V(p.pos.X-(PlayerSpeed*dt), p.pos.Y)
					p.distanceLeft -= float64(PlayerSpeed * dt)
				} else {
					p.tileID = p.nextTileID
					p.pos = pixel.V(tiles[p.tileID].pos.X+float64(TileSize/2), tiles[p.tileID].pos.Y+float64(TileSize/2))
					p.distanceLeft = 50
					p.moving = false
				}
			case "right":
				if p.distanceLeft > 0 {
					p.pos = pixel.V(p.pos.X+(PlayerSpeed*dt), p.pos.Y)
					p.distanceLeft -= float64(PlayerSpeed * dt)
				} else {
					p.tileID = p.nextTileID
					p.pos = pixel.V(tiles[p.tileID].pos.X+float64(TileSize/2), tiles[p.tileID].pos.Y+float64(TileSize/2))
					p.distanceLeft = 50
					p.moving = false
				}
			default:
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
