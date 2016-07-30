package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	W_WIDTH  int32 = 800
	W_HEIGHT int32 = 600
)

// Key press constants
const (
	KEY_PRESS_SURFACE_DEFAULT = iota
	KEY_PRESS_SURFACE_UP
	KEY_PRESS_SURFACE_DOWN
	KEY_PRESS_SURFACE_LEFT
	KEY_PRESS_SURFACE_RIGHT
	KEY_PRESS_SURFACE_TOTAL
)

func initWindowAndSurface(title string, w, h int) (*sdl.Window, *sdl.Surface, error) {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow(
		title,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		w, h, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, nil, err
	}

	surface, err := window.GetSurface()
	if err != nil {
		window.Destroy()
		return nil, nil, err
	}

	return window, surface, nil
}

func loadBmpFile(path string) *sdl.Surface {
	s, err := sdl.LoadBMP(path)
	if err != nil {
		panic(err)
	}

	return s
}

func main() {
	var currentSurface *sdl.Surface

	window, windowSurface, err := initWindowAndSurface("02", int(W_WIDTH), int(W_HEIGHT))
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// Do stuff with windowSurface
	clearRect := sdl.Rect{0, 0, W_WIDTH, W_HEIGHT}
	windowSurface.FillRect(&clearRect, 0x00000000)

	surfaces := []*sdl.Surface{
		loadBmpFile("./press.bmp"),
		loadBmpFile("./up.bmp"),
		loadBmpFile("./down.bmp"),
		loadBmpFile("./left.bmp"),
		loadBmpFile("./right.bmp")}

	// Rect where image should be shown
	imgRect := sdl.Rect{80, 60, 640, 480}

	currentSurface = surfaces[0]
	currentSurface.Blit(&currentSurface.ClipRect, windowSurface, &imgRect)

	window.UpdateSurface()

	// Main loop start
	running := true
	var event sdl.Event

	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {

			// Special events
			case *sdl.QuitEvent:
				running = false
				fmt.Println("Quit")

			// Keyboard events
			case *sdl.KeyDownEvent:
				switch t.Keysym.Sym {
				case sdl.K_UP:
					currentSurface = surfaces[KEY_PRESS_SURFACE_UP]
				case sdl.K_DOWN:
					currentSurface = surfaces[KEY_PRESS_SURFACE_DOWN]
				case sdl.K_LEFT:
					currentSurface = surfaces[KEY_PRESS_SURFACE_LEFT]
				case sdl.K_RIGHT:
					currentSurface = surfaces[KEY_PRESS_SURFACE_RIGHT]
				default:
					currentSurface = surfaces[KEY_PRESS_SURFACE_DEFAULT]
				}
			}
		}

		// Set updated surface and update windowSurface
		currentSurface.Blit(&currentSurface.ClipRect, windowSurface, &imgRect)
		window.UpdateSurface()
	}

	// Main loop end

	sdl.Quit()
}
