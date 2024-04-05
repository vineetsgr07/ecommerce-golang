// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type PostStatus string

const (
	PostStatusDraft     PostStatus = "draft"
	PostStatusPublished PostStatus = "published"
)

func (e *PostStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PostStatus(s)
	case string:
		*e = PostStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for PostStatus: %T", src)
	}
	return nil
}

type NullPostStatus struct {
	PostStatus PostStatus
	Valid      bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPostStatus) Scan(value interface{}) error {
	if value == nil {
		ns.PostStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PostStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPostStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.PostStatus, nil
}

type UserStatus string

const (
	UserStatusDisabled   UserStatus = "disabled"
	UserStatusUnverified UserStatus = "unverified"
	UserStatusActive     UserStatus = "active"
)

func (e *UserStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserStatus(s)
	case string:
		*e = UserStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for UserStatus: %T", src)
	}
	return nil
}

type NullUserStatus struct {
	UserStatus UserStatus
	Valid      bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserStatus) Scan(value interface{}) error {
	if value == nil {
		ns.UserStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.UserStatus, nil
}

type Post struct {
	ID        int64      `json:"id"`
	AuthorID  int64      `json:"author_id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	Status    PostStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type Reset struct {
	UserID    int64     `json:"user_id"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	ID           int64      `json:"id"`
	Email        string     `json:"email"`
	Pass         string     `json:"pass"`
	Salt         string     `json:"salt"`
	Status       UserStatus `json:"status"`
	Verification string     `json:"verification"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}