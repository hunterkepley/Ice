package main

import (
	"fmt"
	_ "image"
	_ "image/jpeg"
	_ "image/png"
	_ "os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

var (
	tiles    []Tile
	players  []Player
	frames   = 0
	second   = time.Tick(time.Second)
	ingame   = false
	menuPlay = true
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
			tiles = append(tiles, Tile{0, pixel.V(float64(x*TileSize), float64(y*TileSize)), len(tiles), 3.0})
		}
	}

	// Create players
	players = append(players, newPlayer(pixel.V(tiles[30].pos.X+float64(TileSize/2), tiles[30].pos.Y+float64(TileSize/2)), 30, "art/player1.png", 0))
	players = append(players, newPlayer(pixel.V(tiles[162].pos.X+float64(TileSize/2), tiles[162].pos.Y+float64(TileSize/2)), 162, "art/player2.png", 1))

	win.SetSmooth(false)

	// Menu definitions
	menuPic, err := loadPicture("art/bg.jpg")
	if err != nil {
		panic(err)
	}
	menubg := pixel.NewSprite(menuPic, menuPic.Bounds())

	playPic, err := loadPicture("art/playButton.png")
	if err != nil {
		panic(err)
	}
	playButton := pixel.NewSprite(playPic, playPic.Bounds())

	quitPic, err := loadPicture("art/quitButton.png")
	if err != nil {
		panic(err)
	}
	quitButton := pixel.NewSprite(quitPic, quitPic.Bounds())

	arrowPic, err := loadPicture("art/arrow.png")
	if err != nil {
		panic(err)
	}
	arrow := pixel.NewSprite(arrowPic, arrowPic.Bounds())

	// end ^

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
				if timer > 0 {
					tiles[i].update(dt)
				}
				tiles[i].render(imd)
			}

			drawLines(WinWidth, WinHeight, imd)

			imd.Draw(win) // Draw shapes

			for i := 0; i < len(players); i++ {
				if timer > 0 {
					players[i].update(&tiles[players[i].tileID], win, dt)
				}
				players[i].render(win)
			}

			timer -= 1.0 * dt

			timerTxt := text.New(pixel.V(300, WinHeight-130), basicAtlas)

			if timer > 0 {
				timerTxt = text.New(pixel.V(280, WinHeight-130), basicAtlas)
				timerTxt.Color = colornames.Darkturquoise
				fmt.Fprintln(timerTxt, fmt.Sprintf("Time left: %d", int(timer)))
			} else {
				restartTimer -= 1 * dt
				if restartTimer <= 0 { // Reset
					for i := 0; i < len(tiles); i++ {
						tiles[i].state = 0
					}
					players[0].tileID = 30
					players[1].tileID = 162
					players[0].pos = pixel.V(tiles[30].pos.X+float64(TileSize/2), tiles[30].pos.Y+float64(TileSize/2))
					players[1].pos = pixel.V(tiles[162].pos.X+float64(TileSize/2), tiles[162].pos.Y+float64(TileSize/2))
					players[0].nextTileID = 30
					players[1].nextTileID = 162
					players[0].dir = "right"
					players[1].dir = "left"
					timer = 120
					restartTimer = 5
				}
				timerTxt = text.New(pixel.V(300, WinHeight-130), basicAtlas)
				timerTxt.Color = colornames.Chocolate
				fmt.Fprintln(timerTxt, "Game Over!")
				for i := 0; i < len(tiles); i++ {
					if tiles[i].state == 1 || tiles[i].state == 3 {
						amt1++
					} else if tiles[i].state == 2 || tiles[i].state == 4 {
						amt2++
					}
				}
				if amt1 > amt2 {
					winTxt = text.New(pixel.V(280, WinHeight-130), basicAtlas)
					fmt.Fprintln(timerTxt, fmt.Sprintf("Blue wins!\n [%d-%d]", amt1, amt2))
					winTxt.Draw(win, pixel.IM.Moved(pixel.V(300, WinHeight-160)))
				} else if amt1 < amt2 {
					winTxt = text.New(pixel.V(280, WinHeight-130), basicAtlas)
					fmt.Fprintln(timerTxt, fmt.Sprintf("Red wins!\n [%d-%d]", amt1, amt2))
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

			mat := pixel.IM.Moved(pixel.V(0+(menubg.Frame().Size().X/2), 0+(menubg.Frame().Size().Y/2)))
			menubg.Draw(win, mat)

			mat = pixel.IM.Moved(pixel.V(0+(playButton.Frame().Size().X/2), 350+(playButton.Frame().Size().Y/2)))
			playButton.Draw(win, mat)

			mat = pixel.IM.Moved(pixel.V(0+(quitButton.Frame().Size().X/2), (350-playButton.Frame().Size().Y)+(playButton.Frame().Size().Y/2)))
			quitButton.Draw(win, mat)

			if menuPlay {
				mat = pixel.IM.Moved(pixel.V((10+(quitButton.Frame().Size().X))+(arrow.Frame().Size().X/2), 350+(arrow.Frame().Size().Y/2)))
			} else {
				mat = pixel.IM.Moved(pixel.V((10+(quitButton.Frame().Size().X))+(arrow.Frame().Size().X/2), (350-playButton.Frame().Size().Y)+(arrow.Frame().Size().Y/2)))
			}
			arrow.Draw(win, mat)

			if win.JustPressed(pixelgl.KeyDown) || win.JustPressed(pixelgl.KeyUp) {
				if menuPlay {
					menuPlay = false
				} else {
					menuPlay = true
				}
			}
			if win.JustPressed(pixelgl.KeyEnter) && menuPlay {
				timer = 120
				restartTimer = 5
				amt1 = 0
				amt2 = 0
				ingame = true
			} else if win.JustPressed(pixelgl.KeyEnter) && !menuPlay {
				win.Destroy()
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
