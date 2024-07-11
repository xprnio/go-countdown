package internal

import (
	"time"
)

type PomodoroSession struct {
	Duration time.Duration
	// Start time.Time
	// End   time.Time
}

// func (s PomodoroSession) Duration() time.Duration {
// 	return time.Now().Sub(s.Start)
// }

// func (s PomodoroSession) IsRunning() bool {
//   return s.End.IsZero()
// }

type PomodoroState int

const (
	PomodoroStateStopped PomodoroState = iota
	PomodoroStateRunning
	PomodoroStatePaused
)

type Pomodoro struct {
	SessionDuration time.Duration
	PauseDuration   time.Duration

	State    PomodoroState
	Sessions []*PomodoroSession
	Pause    time.Duration
}

func (p *Pomodoro) Reset() {
	p.State = PomodoroStateRunning
	p.Sessions = nil
	p.Pause = 0
}

func NewPomodoro(sessionDuration, pauseDuration time.Duration) *Pomodoro {
	return &Pomodoro{
		SessionDuration: sessionDuration,
		PauseDuration:   pauseDuration,
	}
}

func (p *Pomodoro) StartSession() *PomodoroSession {
	s := &PomodoroSession{}
	p.Sessions = append(p.Sessions, s)
	p.State = PomodoroStateRunning
	return s
}

func (p *Pomodoro) FinishSession() {
	switch p.State {
	case PomodoroStateRunning:
		p.State = PomodoroStatePaused
	}
}

func (p *Pomodoro) FinishPause() {
	switch p.State {
	case PomodoroStatePaused:
		p.Pause = 0
		p.State = PomodoroStateRunning
		p.StartSession()
	}
}

func (p *Pomodoro) IsPaused() bool {
	return p.State == PomodoroStatePaused
}

func (p *Pomodoro) Tick(dt float32) {
	d := time.Duration(dt * float32(time.Second))
	switch p.State {
	case PomodoroStateRunning:
		s := p.CurrentSession()
		s.Duration += d

		if s.Duration >= p.SessionDuration {
			p.FinishSession()
		}
	case PomodoroStatePaused:
		p.Pause += d

		if p.Pause >= p.PauseDuration {
			p.FinishPause()
		}
	}
}

func (p *Pomodoro) Start() {
	switch p.State {
	case PomodoroStateStopped:
		p.StartSession()
	}
}

func (p *Pomodoro) CurrentSession() *PomodoroSession {
	if len(p.Sessions) == 0 {
		p.StartSession()
	}

	last := len(p.Sessions) - 1
	return p.Sessions[last]
}

func (p *Pomodoro) Duration() time.Duration {
	switch p.State {
	case PomodoroStateRunning:
		s := p.CurrentSession()
		return s.Duration
	case PomodoroStatePaused:
		return p.Pause
	default:
		return 0
	}
}

func (p *Pomodoro) Progress() float32 {
	switch p.State {
	case PomodoroStateRunning:
		return float32(p.Duration()) / float32(p.SessionDuration)
	case PomodoroStatePaused:
		return float32(p.Duration()) / float32(p.PauseDuration)
	default:
		return 0
	}
}
