package dto

type ScheduleResult struct {
	ID        string
	UserID    string
	Date      string
	Time      string
	Title     string
	Duration  string
	Type      string
	Notes     string
	Completed bool
	CreatedAt string
	UpdatedAt string
}
