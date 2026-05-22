package entity

import "time"

type Goal struct {
	ID        string
	UserID    string
	Title     string
	Type      string
	Target    int
	Current   int
	Unit      string
	Period    string
	Deadline  string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewGoal(userID, title, typ string, target int, unit, period, deadline string) *Goal {
	now := time.Now()
	return &Goal{
		UserID:    userID,
		Title:     title,
		Type:      typ,
		Target:    target,
		Current:   0,
		Unit:      unit,
		Period:    period,
		Deadline:  deadline,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (g *Goal) Update(title, typ string, target, current int, unit, period, deadline string) {
	g.Title = title
	g.Type = typ
	g.Target = target
	g.Current = current
	g.Unit = unit
	g.Period = period
	g.Deadline = deadline
	g.UpdatedAt = time.Now()
}

func (g *Goal) UpdateProgress(current int) {
	g.Current = current
	if g.Current >= g.Target {
		g.Completed = true
	}
	g.UpdatedAt = time.Now()
}

func (g *Goal) Toggle() {
	g.Completed = !g.Completed
	g.UpdatedAt = time.Now()
}
