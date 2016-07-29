package main

import "github.com/veandco/go-sdl2/sdl"

const (
	W_WIDTH  int32 = 800
	W_HEIGHT int32 = 600
)

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int(W_WIDTH), int(W_HEIGHT), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	rect1 := sdl.Rect{0, 0, W_WIDTH / 2, W_HEIGHT / 2}
	rect2 := sdl.Rect{W_WIDTH / 2, 0, W_WIDTH / 2, W_HEIGHT / 2}
	rect3 := sdl.Rect{0, W_HEIGHT / 2, W_WIDTH / 2, W_HEIGHT / 2}
	rect4 := sdl.Rect{W_WIDTH / 2, W_HEIGHT / 2, W_WIDTH / 2, W_HEIGHT / 2}

	surface.FillRect(&rect1, 0xffff0000)
	surface.FillRect(&rect2, 0xffffff00)
	surface.FillRect(&rect3, 0xff00ff00)
	surface.FillRect(&rect4, 0xff0000ff)

	window.UpdateSurface()

	sdl.Delay(3000)
	sdl.Quit()
}
