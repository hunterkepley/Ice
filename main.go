package main

import (
	"fmt"
	_ "image"
	_ "os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

var (
	tiles   []Tile
	players []Player
	frames  = 0
	second  = time.Tick(time.Second)
	ingame  = false
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

	face, err := loadTTF("fonts/chintzy.ttf", 35)
	if err != nil {
		panic(err)
	}

	basicAtlas := text.NewAtlas(face, text.ASCII)
	timerTxt := text.New(pixel.V(300, 300), basicAtlas)
	winTxt := text.New(pixel.V(300, 300), basicAtlas)
	timerTxt.Color = colornames.Darkturquoise
	winTxt.Color = colornames.Gold

	amt1 := 0 // Amount of tiles at the end [p1]
	amt2 := 0 // Amount of tiles at the end [p2]

	last := time.Now()
	timer := 120.0
	restartTimer := 5.0
	for !win.Closed() {
		if ingame {
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

			timer -= 1.0 * dt

			timerTxt := text.New(pixel.V(300, WinHeight-130), basicAtlas)

			if timer > 0 {
				timerTxt = text.New(pixel.V(280, WinHeight-130), basicAtlas)
				timerTxt.Color = colornames.Darkturquoise
				fmt.Fprintln(timerTxt, fmt.Sprintf("Time left: %d", int(timer)))
			} else {
				timerTxt = text.New(pixel.V(300, WinHeight-130), basicAtlas)
				timerTxt.Color = colornames.Crimson
				fmt.Fprintln(timerTxt, "Game Over!")
				for i := 0; i < len(tiles); i++ {
					if tiles[i].state == 1 {
						amt1++
					} else if tiles[i].state == 2 {
						amt2++
					}
				}
				if amt1 > amt2 {
					winTxt = text.New(pixel.V(280, WinHeight-130), basicAtlas)
					fmt.Fprintln(timerTxt, "Player 1 wins!")
					winTxt.Draw(win, pixel.IM.Moved(pixel.V(300, WinHeight-160)))
				} else if amt1 < amt2 {
					winTxt = text.New(pixel.V(280, WinHeight-130), basicAtlas)
					fmt.Fprintln(timerTxt, "Player 2 wins!")
					winTxt.Draw(win, pixel.IM.Moved(pixel.V(300, WinHeight-160)))
				} else {
					winTxt = text.New(pixel.V(280, WinHeight-130), basicAtlas)
					fmt.Fprintln(timerTxt, "Tied!")
					winTxt.Draw(win, pixel.IM.Moved(pixel.V(300, WinHeight-160)))
				}
				amt1 = 0
				amt2 = 0

			}
			timerTxt.Draw(win, pixel.IM.Moved(pixel.V(10, 100)))
		} else {
			last = time.Now()
			win.Clear(colornames.Darkgoldenrod)
			if win.Pressed(pixelgl.KeyEnter) {
				timer = 10
				restartTimer = 5
				amt1 = 0
				amt2 = 0
				ingame = true
			}
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
