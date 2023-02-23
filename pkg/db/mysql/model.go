package mysql

import "time"

type Model struct {
	Id        uint       `json:"id" gorm:"column:id;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"updated_at"`
}
