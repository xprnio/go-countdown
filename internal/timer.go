package internal

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Timer struct {
	*Pomodoro
}

func NewTimer(sessionDuration, pauseDuration time.Duration) *Timer {
	return &Timer{
		Pomodoro: &Pomodoro{
			SessionDuration: sessionDuration,
			PauseDuration:   pauseDuration,
		},
	}
}

func (t Timer) Draw(center rl.Vector2) {
	t.DrawSessions(center)
	t.DrawBreakLabel(center)
	t.DrawTimer(center)
	t.DrawPomodoro()
}

func (t Timer) DrawSessions(center rl.Vector2) {
	screenHeight := float32(rl.GetScreenHeight())
	color := rl.Lime

	if t.IsPaused() {
		color = rl.Blue
	}

	for i := range t.Sessions {
		const padding = 10
		const size = 10
		const spacing = 2
		offset := float32(i+1)*size + float32(i)*spacing
		rect := rl.NewRectangle(
			padding, screenHeight-padding-offset,
			40, 10,
		)

		active := i == len(t.Sessions)-1
		if active {
			rl.DrawRectangleLinesEx(rect, 1, color)
		} else {
			rl.DrawRectangleRec(rect, color)
		}
	}
}

func (t Timer) DrawPomodoro() {
	windowWidth := rl.GetScreenWidth()
	windowHeight := rl.GetScreenHeight()
	bounds := rl.NewRectangle(
		float32(windowWidth)*.15,
		float32(windowHeight)*.8,
		float32(windowWidth)*.7,
		float32(windowHeight)*.1,
	)

	if t.IsPaused() {
		rl.DrawRectangleRec(bounds, rl.ColorAlpha(rl.Black, 0.2))
	} else {
		rl.DrawRectangleRec(bounds, rl.ColorAlpha(rl.White, 0.2))
	}

	progressRect := rl.NewRectangle(
		bounds.X, bounds.Y,
		bounds.Width*t.Progress(),
		bounds.Height,
	)
	switch t.State {
	case PomodoroStatePaused:
		rl.DrawRectangleRec(progressRect, rl.Red)
	case PomodoroStateRunning:
		rl.DrawRectangleRec(progressRect, rl.Magenta)
	}
}

func (t Timer) DrawBreakLabel(center rl.Vector2) {
	font := rl.GetFontDefault()
	windowHeight := rl.GetScreenHeight()

	if t.IsPaused() {
		fontSize := float32(windowHeight) * .2
		fontSpacing := fontSize * .2
		size := rl.MeasureTextEx(font, "TAKE A BREAK", fontSize, fontSpacing)
		position := rl.NewVector2(center.X-size.X/2, 4)
		rl.DrawTextEx(font, "TAKE A BREAK", position, fontSize, fontSpacing, rl.ColorAlpha(rl.Black, 0.8))
	}
}

func (t Timer) DrawTimer(center rl.Vector2) {
	font := rl.GetFontDefault()
	dur := DrawableDuration(t.Duration())

	color := rl.White
	if t.IsPaused() {
		color = rl.Black
	}

	rects := dur.Rectangles(center)
	for _, r := range rects {
		fontSize := r.Height * .8
		fontSpacing := fontSize * .2
		position := rl.Vector2Add(
			rl.NewVector2(r.X, r.Y),
			rl.Vector2Scale(rl.Vector2Subtract(
				rl.NewVector2(r.Width, r.Height),
				rl.MeasureTextEx(font, r.Text, fontSize, fontSpacing),
			), 0.5),
		)

		rl.DrawTextEx(font, r.Text, position, fontSize, fontSpacing, color)
	}
}
