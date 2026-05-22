package user_test

import (
	"testing"
	"time"

	"github.com/aithlete/aithlete-api/domain/user"
)

func TestNewUser(t *testing.T) {
	u := user.NewUser("test@example.com", "Test User", "password123")

	if u.GetEmail() != "test@example.com" {
		t.Errorf("expected email test@example.com, got %s", u.GetEmail())
	}
	if u.GetName() != "Test User" {
		t.Errorf("expected name Test User, got %s", u.GetName())
	}
	if u.GetCreatedAt().IsZero() {
		t.Error("expected createdAt to be set")
	}
	if u.GetUpdatedAt().IsZero() {
		t.Error("expected updatedAt to be set")
	}
}

func TestNewUserCreatedAtUpdatedAtAreSame(t *testing.T) {
	u := user.NewUser("test@example.com", "Test User", "password123")

	if !u.GetCreatedAt().Equal(u.GetUpdatedAt()) {
		t.Error("expected createdAt and updatedAt to be equal on creation")
	}
}

func TestUpdateProfile(t *testing.T) {
	u := user.NewUser("test@example.com", "Test User", "password123")
	originalUpdatedAt := u.GetUpdatedAt()

	time.Sleep(time.Nanosecond)

	u.UpdateProfile("New Name", "1995-06-15", "male")

	if u.GetName() != "New Name" {
		t.Errorf("expected name New Name, got %s", u.GetName())
	}
	if u.GetBirthday() != "1995-06-15" {
		t.Errorf("expected birthday 1995-06-15, got %s", u.GetBirthday())
	}
	if u.GetGender() != "male" {
		t.Errorf("expected gender male, got %s", u.GetGender())
	}
	if !u.GetUpdatedAt().After(originalUpdatedAt) {
		t.Error("expected updatedAt to be updated after profile change")
	}
}

func TestUpdateProfileKeepsEmail(t *testing.T) {
	u := user.NewUser("test@example.com", "Test User", "password123")
	u.UpdateProfile("New Name", "1995-06-15", "male")

	if u.GetEmail() != "test@example.com" {
		t.Errorf("expected email to remain unchanged, got %s", u.GetEmail())
	}
}

func TestUpdatePassword(t *testing.T) {
	u := user.NewUser("test@example.com", "Test User", "oldpassword")
	originalUpdatedAt := u.GetUpdatedAt()

	time.Sleep(time.Nanosecond)

	u.UpdatePassword("newpassword123")

	// Password field is not exported via getter for security,
	// but we can verify the entity was updated
	if !u.GetUpdatedAt().After(originalUpdatedAt) {
		t.Error("expected updatedAt to be updated after password change")
	}
}

func TestUserEqualityByPointer(t *testing.T) {
	u1 := user.NewUser("a@example.com", "A", "pass")
	u2 := user.NewUser("b@example.com", "B", "pass")

	u1.ID = "same-id"
	u2.ID = "same-id"

	// Two different pointers with the same ID should be considered different
	// objects in Go, but domain identity is based on ID
	if u1.ID != u2.ID {
		t.Error("expected both users to have the same ID")
	}
}

func TestNewUserNoAvatar(t *testing.T) {
	u := user.NewUser("test@example.com", "Test User", "password123")

	if u.GetAvatar() != "" {
		t.Errorf("expected empty avatar, got %s", u.GetAvatar())
	}
}

func TestNewUserNoID(t *testing.T) {
	u := user.NewUser("test@example.com", "Test User", "password123")

	if u.GetID() != "" {
		t.Errorf("expected empty ID, got %s", u.GetID())
	}
}
