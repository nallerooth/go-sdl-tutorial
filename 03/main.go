package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	W_WIDTH  int32 = 800
	W_HEIGHT int32 = 600
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

func main() {
	window, surface, err := initWindowAndSurface("02", int(W_WIDTH), int(W_HEIGHT))
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// Do stuff with surface
	clearRect := sdl.Rect{0, 0, W_WIDTH, W_HEIGHT}
	surface.FillRect(&clearRect, 0x00000000)

	helloWorld, err := sdl.LoadBMP("./hello_world.bmp")
	if err != nil {
		panic(err)
	}

	// Rect where image should be shown
	helloRect := sdl.Rect{80, 60, 640, 480}
	helloWorld.Blit(&helloWorld.ClipRect, surface, &helloRect)

	window.UpdateSurface()

	// Main loop start
	running := true
	var event sdl.Event

	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
				fmt.Println("Quit")
			case *sdl.KeyDownEvent:
				fmt.Printf("Sym:%c\tMod:%d\n", t.Keysym.Sym, t.Keysym.Mod)
			}
		}
	}

	// Main loop end

	sdl.Quit()
}
