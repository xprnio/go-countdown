package main

import (
	"time"

	"github.com/gen2brain/raylib-go/raylib"
	"github.com/xprnio/countdown/internal"
)

const (
	windowWidth  int32 = 400
	windowHeight int32 = 100
)

// pomodoro constants
const (
	// sessionDuration = 5 * time.Second
	// pauseDuration   = 5 * time.Second
	sessionDuration = 25 * time.Minute
	pauseDuration   = 5 * time.Minute
)

func initWindow() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(windowWidth, windowHeight, "gomodoro")
	rl.SetTargetFPS(60)
	rl.SetWindowMinSize(int(windowWidth), int(windowHeight))
	rl.SetWindowMaxSize(int(windowWidth), int(windowHeight))
	rl.SetExitKey(0)
}

func main() {
	initWindow()
	defer rl.CloseWindow()

	timer := internal.NewTimer(sessionDuration, pauseDuration)
	timer.Start()

	paused := false
	resetting := false
	quitting := false

	quit := false

	for !quit {
		dt := rl.GetFrameTime()

		if !paused && !resetting && !quitting {
			timer.Tick(dt)
		}

		switch {
		case rl.IsKeyPressed(rl.KeySpace):
			switch {
			case quitting:
			case resetting:
			default:
				paused = !paused
			}
		case rl.IsKeyPressed(rl.KeyR):
			switch {
			case quitting:
			case paused:
			default:
				resetting = true
			}
		case rl.IsKeyPressed(rl.KeyEnter):
			switch {
			case resetting:
				resetting = false
				timer.Reset()
			case quitting:
				quit = true
			}
		case rl.IsKeyPressed(rl.KeyEscape):
			switch {
			case paused:
			case resetting:
				resetting = false
			case quitting:
				quitting = false
			default:
				quitting = true
			}
		}

		quit = quit || rl.WindowShouldClose()

		rl.BeginDrawing()
		{
			drawTimer(*timer)

			if quitting {
				windowWidth := rl.GetScreenWidth()
				windowHeight := rl.GetScreenHeight()
				rl.DrawRectangle(0, 0, int32(windowWidth), int32(windowHeight), rl.ColorAlpha(rl.Black, .9))

				const height = 48
				width := rl.MeasureText("QUIT?", height)
				x := (int32(windowWidth) - width) / 2
				y := (int32(windowHeight) - height) / 2
				rl.DrawText("QUIT?", x, y, height, rl.Red)
			}

			if resetting {
				windowWidth := rl.GetScreenWidth()
				windowHeight := rl.GetScreenHeight()
				rl.DrawRectangle(0, 0, int32(windowWidth), int32(windowHeight), rl.ColorAlpha(rl.Black, .9))

				const height = 48
				width := rl.MeasureText("RESET?", height)
				x := (int32(windowWidth) - width) / 2
				y := (int32(windowHeight) - height) / 2
				rl.DrawText("RESET?", x, y, height, rl.Red)
			}

			if paused {
				windowWidth := rl.GetScreenWidth()
				windowHeight := rl.GetScreenHeight()
				rl.DrawRectangle(0, 0, int32(windowWidth), int32(windowHeight), rl.ColorAlpha(rl.Black, .9))

				const height = 48
				width := rl.MeasureText("PAUSED", height)
				x := (int32(windowWidth) - width) / 2
				y := (int32(windowHeight) - height) / 2
				rl.DrawText("PAUSED", x, y, height, rl.Red)
			}
		}
		rl.EndDrawing()
	}
}

func drawTimer(timer internal.Timer) {
	windowWidth := rl.GetScreenWidth()
	windowHeight := rl.GetScreenHeight()
	if timer.IsPaused() {
		rl.ClearBackground(rl.RayWhite)
	} else {
		rl.ClearBackground(rl.Black)
	}

	center := rl.Vector2Scale(
		rl.NewVector2(float32(windowWidth), float32(windowHeight)), 0.5,
	)

	timer.Draw(center)
}
