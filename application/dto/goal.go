package dto

type GoalResult struct {
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
	CreatedAt string
	UpdatedAt string
}
