package goal

import (
	"testing"
	"time"
)

func TestNewGoal(t *testing.T) {
	g := NewGoal("user-1", "Run 100 miles", "custom", 100, "miles", "monthly", "2026-05-31")

	if g.UserID != "user-1" {
		t.Errorf("expected UserID user-1, got %s", g.UserID)
	}
	if g.Title != "Run 100 miles" {
		t.Errorf("expected Title Run 100 miles, got %s", g.Title)
	}
	if g.Current != 0 {
		t.Errorf("expected Current 0, got %d", g.Current)
	}
	if g.Completed {
		t.Error("expected new goal to be not completed")
	}
	if g.CreatedAt.IsZero() || g.UpdatedAt.IsZero() {
		t.Error("expected timestamps to be set")
	}
}

func TestUpdateProgress(t *testing.T) {
	g := NewGoal("user-1", "Run 100 miles", "custom", 100, "miles", "monthly", "2026-05-31")
	originalUpdatedAt := g.UpdatedAt

	time.Sleep(time.Nanosecond)

	g.UpdateProgress(50)

	if g.Current != 50 {
		t.Errorf("expected Current 50, got %d", g.Current)
	}
	if g.Completed {
		t.Error("expected goal to not be completed yet")
	}
	if !g.UpdatedAt.After(originalUpdatedAt) {
		t.Error("expected updatedAt to change after progress update")
	}
}

func TestUpdateProgressCompletesGoal(t *testing.T) {
	g := NewGoal("user-1", "Run 100 miles", "custom", 100, "miles", "monthly", "2026-05-31")
	g.UpdateProgress(100)

	if !g.Completed {
		t.Error("expected goal to be completed when current reaches target")
	}
}

func TestUpdateProgressExceedsTarget(t *testing.T) {
	g := NewGoal("user-1", "Run 100 miles", "custom", 100, "miles", "monthly", "2026-05-31")
	g.UpdateProgress(150)

	if !g.Completed {
		t.Error("expected goal to be completed when current exceeds target")
	}
	if g.Current != 150 {
		t.Errorf("expected Current 150, got %d", g.Current)
	}
}

func TestGoalUpdate(t *testing.T) {
	g := NewGoal("user-1", "Run 100 miles", "custom", 100, "miles", "monthly", "2026-05-31")
	originalUpdatedAt := g.UpdatedAt

	time.Sleep(time.Nanosecond)

	g.Update("Run 200 miles", "custom", 200, 50, "miles", "monthly", "2026-06-30")

	if g.Title != "Run 200 miles" {
		t.Errorf("expected Title Run 200 miles, got %s", g.Title)
	}
	if g.Target != 200 {
		t.Errorf("expected Target 200, got %d", g.Target)
	}
	if !g.UpdatedAt.After(originalUpdatedAt) {
		t.Error("expected updatedAt to change after update")
	}
}

func TestGoalToggle(t *testing.T) {
	g := NewGoal("user-1", "Run 100 miles", "custom", 100, "miles", "monthly", "2026-05-31")
	g.Toggle()

	if !g.Completed {
		t.Error("expected goal to be completed after toggle")
	}

	g.Toggle()
	if g.Completed {
		t.Error("expected goal to be not completed after second toggle")
	}
}
