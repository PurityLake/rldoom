package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 600, "RLDoom")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("RLDoom", 10, 10, 20, rl.DarkGray)

		rl.EndDrawing()
	}
}
