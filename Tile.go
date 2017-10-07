package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	_ "github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	TileSize = 50 // W and H of tiles
	TilesX   = 16 // Amount of tiles on the x axis [800/50]
	TilesY   = 12 // Amount of tiles on the y axis [600/50]
)

type Tile struct {
	state int // 0 = nothing, 1 = purple, 2 = gray
	pos   pixel.Vec
	ID    int
}

func (t *Tile) update() { // Updates a tile

}

func (t Tile) render(imd *imdraw.IMDraw) { // Draws a tile
	switch t.state {
	case 0:
		imd.Color = colornames.Snow
	case 1:
		imd.Color = colornames.Aqua
	case 2:
		imd.Color = colornames.Lightsteelblue
	default:
	}
	dPos1, dPos2, dSize := t.getRectangle()
	imd.Push(dPos1, dPos2)
	imd.Line(dSize)
}

// Returns the start (x,y), end (x,y), and thickness
// This was made because I did not need a full polygon made,
// And just making a single line most likely is more effecient
// Than making a full polygon for a rectangle/square and filling it in.
func (t Tile) getRectangle() (pixel.Vec, pixel.Vec, float64) { // Returns start pt, end pt, and thickness
	return pixel.V(t.pos.X+TileSize/2, t.pos.Y), pixel.V(t.pos.X+TileSize/2, t.pos.Y+TileSize), TileSize
}
