package schedule

import "time"

type Schedule struct {
	ID        string
	UserID    string
	Date      string
	Time      string
	Title     string
	Duration  string
	Type      string
	Notes     string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewSchedule(userID, date, schedTime, title, duration, typ, notes string) *Schedule {
	now := time.Now()
	return &Schedule{
		UserID:    userID,
		Date:      date,
		Time:      schedTime,
		Title:     title,
		Duration:  duration,
		Type:      typ,
		Notes:     notes,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (s *Schedule) Update(date, schedTime, title, duration, typ, notes string) {
	s.Date = date
	s.Time = schedTime
	s.Title = title
	s.Duration = duration
	s.Type = typ
	s.Notes = notes
	s.UpdatedAt = time.Now()
}

func (s *Schedule) Toggle() {
	s.Completed = !s.Completed
	s.UpdatedAt = time.Now()
}
