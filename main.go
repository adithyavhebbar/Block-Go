package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(600, 400, "Blocks")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
