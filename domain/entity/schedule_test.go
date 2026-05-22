package entity

import (
	"testing"
	"time"
)

func TestNewSchedule(t *testing.T) {
	s := NewSchedule("user-1", "2026-05-20", "06:30", "Morning Run", "30 min", "cardio", "")

	if s.UserID != "user-1" {
		t.Errorf("expected UserID user-1, got %s", s.UserID)
	}
	if s.Title != "Morning Run" {
		t.Errorf("expected Title Morning Run, got %s", s.Title)
	}
	if s.Completed {
		t.Error("expected new schedule to be not completed")
	}
	if s.CreatedAt.IsZero() || s.UpdatedAt.IsZero() {
		t.Error("expected timestamps to be set")
	}
}

func TestScheduleUpdate(t *testing.T) {
	s := NewSchedule("user-1", "2026-05-20", "06:30", "Morning Run", "30 min", "cardio", "")
	originalUpdatedAt := s.UpdatedAt

	time.Sleep(time.Nanosecond)

	s.Update("2026-05-21", "07:00", "Yoga", "40 min", "stretching", "Morning session")

	if s.Date != "2026-05-21" {
		t.Errorf("expected Date 2026-05-21, got %s", s.Date)
	}
	if s.Title != "Yoga" {
		t.Errorf("expected Title Yoga, got %s", s.Title)
	}
	if !s.UpdatedAt.After(originalUpdatedAt) {
		t.Error("expected updatedAt to change after update")
	}
}

func TestScheduleToggle(t *testing.T) {
	s := NewSchedule("user-1", "2026-05-20", "06:30", "Morning Run", "30 min", "cardio", "")
	originalUpdatedAt := s.UpdatedAt

	time.Sleep(time.Nanosecond)

	s.Toggle()
	if !s.Completed {
		t.Error("expected schedule to be completed after toggle")
	}
	if !s.UpdatedAt.After(originalUpdatedAt) {
		t.Error("expected updatedAt to change after toggle")
	}

	s.Toggle()
	if s.Completed {
		t.Error("expected schedule to be not completed after second toggle")
	}
}

func TestNewScheduleDefaults(t *testing.T) {
	s := NewSchedule("user-1", "2026-05-20", "06:30", "Morning Run", "30 min", "cardio", "")

	if s.Completed {
		t.Error("expected completed to default to false")
	}
	if s.Notes != "" {
		t.Errorf("expected empty notes, got %s", s.Notes)
	}
}
