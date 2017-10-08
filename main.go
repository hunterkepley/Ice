package main

import (
	"fmt"
	_ "image"
	_ "os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	tiles   []Tile
	players []Player
	frames  = 0
	second  = time.Tick(time.Second)
)

const (
	WinWidth  = 800
	WinHeight = 600
)

func drawLines(w int, h int, imd *imdraw.IMDraw) { // Draws the lines separating the tiles
	amountX := (w / TileSize) - 1 // Amount of lines on the x axis
	amountY := (h / TileSize) - 1 // Amount of lines on the y axis
	imd.Color = colornames.Ivory
	for i := 0; i < amountX; i++ {
		p1 := pixel.V(float64((i+1)*TileSize), 0)
		p2 := pixel.V(float64((i+1)*TileSize), float64(h))
		imd.Push(p1, p2)
		imd.Line(1)
	}
	for i := 0; i < amountY; i++ {
		p1 := pixel.V(0, float64((i+1)*TileSize))
		p2 := pixel.V(float64(w), float64((i+1)*TileSize))
		imd.Push(p1, p2)
		imd.Line(1)
	}
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Ice!",
		Bounds: pixel.R(0, 0, WinWidth, WinHeight),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	for x := 0; x < TilesX; x++ {
		for y := 0; y < TilesY; y++ { // Tiles go up, then to the next column
			tiles = append(tiles, Tile{0, pixel.V(float64(x*TileSize), float64(y*TileSize)), len(tiles)})
		}
	}

	// Create players
	players = append(players, newPlayer(pixel.V(tiles[30].pos.X+float64(TileSize/2), tiles[30].pos.Y+float64(TileSize/2)), 30, "art/player1.png", 0))
	players = append(players, newPlayer(pixel.V(tiles[162].pos.X+float64(TileSize/2), tiles[162].pos.Y+float64(TileSize/2)), 162, "art/player2.png", 1))

	win.SetSmooth(false)

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		_ = dt
		last = time.Now()

		imd.Clear()

		win.Clear(colornames.Seashell)

		for i := 0; i < len(tiles); i++ {
			tiles[i].render(imd)
		}

		drawLines(WinWidth, WinHeight, imd)

		imd.Draw(win) // Draw shapes

		for i := 0; i < len(players); i++ {
			players[i].update(&tiles[players[i].tileID], win, dt)
			players[i].render(win)
		}

		win.Update()

		frames++
		select { // Waits for the block to finish
		case <-second: // A second has passed
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames)) // Appends fps to title for testing
			frames = 0                                                   // Reset it my dude
		default:
		}
	}
}

func main() {
	pixelgl.Run(run)
}
