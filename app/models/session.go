package models

import (
	"database/sql"
	"time"
)

type Session struct {
	Model
	UUID           string         `db:"uuid" gorm:"unique;not null;index" json:"uuid"`
	UserID         uint           `db:"user_id" gorm:"not null;index" json:"userId,required"`
	TokenId        string         `db:"token_id" gorm:"index" json:"tokenId,omitempty"`
	RefreshTokenId sql.NullString `db:"refresh_token_id" gorm:"index" json:"refreshTokenId,omitempty"`
	IPAddress      string         `db:"ip_addr" gorm:"index" json:"ipAddr,omitempty"`
	UserAgent      string         `db:"user_agent" gorm:"index" json:"userAgent,omitempty"`
	ExpiredAt      time.Time      `db:"expired_at" gorm:"not null" json:"expiredAt"`
}

func (Session) TableName() string {
	return "sessions"
}
