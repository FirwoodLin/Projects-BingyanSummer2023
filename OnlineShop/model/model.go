package model

import (
	"database/sql"
	"time"
)

type TimeModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
