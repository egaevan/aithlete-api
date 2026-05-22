package user

import (
	"time"
)

// User represents a user entity in the system
type User struct {
	ID        string
	Email     string
	Name      string
	Password  string // Never expose this in public methods
	Avatar    string
	Birthday  string
	Gender    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser creates a new user instance
func NewUser(email, name, password string) *User {
	now := time.Now()
	return &User{
		Email:     email,
		Name:      name,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// UpdateProfile updates user profile information
func (u *User) UpdateProfile(name, birthday, gender string) {
	u.Name = name
	u.Birthday = birthday
	u.Gender = gender
	u.UpdatedAt = time.Now()
}

// UpdatePassword updates user password
func (u *User) UpdatePassword(password string) {
	u.Password = password
	u.UpdatedAt = time.Now()
}

// GetID returns the user's ID
func (u *User) GetID() string {
	return u.ID
}

// GetEmail returns the user's email
func (u *User) GetEmail() string {
	return u.Email
}

// GetName returns the user's name
func (u *User) GetName() string {
	return u.Name
}

// GetAvatar returns the user's avatar
func (u *User) GetAvatar() string {
	return u.Avatar
}

// GetBirthday returns the user's birthday
func (u *User) GetBirthday() string {
	return u.Birthday
}

// GetGender returns the user's gender
func (u *User) GetGender() string {
	return u.Gender
}

// GetCreatedAt returns the user's creation time
func (u *User) GetCreatedAt() time.Time {
	return u.CreatedAt
}

// GetUpdatedAt returns the user's last update time
func (u *User) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}