// src/core/domain/mentordm/mentor.go

package mentordm

import (
	"time"
)

type Mentor struct {
	ID        string
	Name      string
	Profile   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Plans     []*Plan // メンターに関連付けられたプランのスライス
}

type Plan struct {
	ID       string
	Title    string
	Category string
	Content  string
	Status   int
	// その他のプランに関連するフィールド
}
