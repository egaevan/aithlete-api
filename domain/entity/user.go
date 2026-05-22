package entity

import "time"

type User struct {
	id        string
	email     string
	name      string
	password  string
	avatar    string
	birthday  string
	gender    string
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(email, name, password string) *User {
	now := time.Now()
	return &User{
		email:     email,
		name:      name,
		password:  password,
		createdAt: now,
		updatedAt: now,
	}
}

func RebuildUser(id, email, name, password, avatar, birthday, gender string, createdAt, updatedAt time.Time) *User {
	return &User{
		id:        id,
		email:     email,
		name:      name,
		password:  password,
		avatar:    avatar,
		birthday:  birthday,
		gender:    gender,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (u *User) SetID(id string)     { u.id = id }
func (u *User) GetID() string       { return u.id }
func (u *User) GetEmail() string    { return u.email }
func (u *User) GetName() string     { return u.name }
func (u *User) PasswordHash() string { return u.password }
func (u *User) GetAvatar() string   { return u.avatar }
func (u *User) GetBirthday() string { return u.birthday }
func (u *User) GetGender() string   { return u.gender }
func (u *User) GetCreatedAt() time.Time { return u.createdAt }
func (u *User) GetUpdatedAt() time.Time { return u.updatedAt }

func (u *User) UpdateProfile(name, birthday, gender string) {
	u.name = name
	u.birthday = birthday
	u.gender = gender
	u.updatedAt = time.Now()
}

func (u *User) UpdatePassword(password string) {
	u.password = password
	u.updatedAt = time.Now()
}
